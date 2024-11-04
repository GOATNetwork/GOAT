package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/goatnetwork/goat/pkg/ethrpc"
	"github.com/goatnetwork/goat/x/goat/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		addressCodec address.Codec
		storeService store.KVStoreService
		logger       log.Logger

		Schema     collections.Schema
		Params     collections.Item[types.Params]
		BeaconRoot collections.Item[[]byte] // the cometbft blockhash
		Block      collections.Item[types.ExecutionPayload]

		engineClient  ethrpc.EngineClient
		bitcoinKeeper types.BitcoinKeeper
		lockingKeeper types.LockingKeeper
		relayerKeeper types.RelayerKeeper
		accountKeeper types.AccountKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService store.KVStoreService,
	logger log.Logger,

	bitcoinKeeper types.BitcoinKeeper,
	lockingKeeper types.LockingKeeper,
	relayerKeeper types.RelayerKeeper,
	accountKeeper types.AccountKeeper,
	engineClient ethrpc.EngineClient,
) Keeper {
	sb := collections.NewSchemaBuilder(storeService)

	k := Keeper{
		cdc:          cdc,
		addressCodec: addressCodec,
		storeService: storeService,
		logger:       logger,

		bitcoinKeeper: bitcoinKeeper,
		lockingKeeper: lockingKeeper,
		relayerKeeper: relayerKeeper,
		accountKeeper: accountKeeper,
		Params:        collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		Block:         collections.NewItem(sb, types.BlockKey, "block", codec.CollValue[types.ExecutionPayload](cdc)),
		BeaconRoot:    collections.NewItem(sb, types.ConsHashKey, "consensus_hash", collections.BytesValue),
		engineClient:  engineClient,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema

	return k
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Finalized notifies goat-geth to update fork choice state
// if there are any errors, the FinalizeBlock phase will be failed
// we don't use timeout here, validators are responsible for a reliable node
func (k Keeper) Finalized(ctx context.Context) error { // EndBlock phase only!
	block, err := k.Block.Get(ctx)
	if err != nil {
		return err
	}

	k.Logger().Info("Notify NewPayload", "number", block.BlockNumber)
	response, err := k.engineClient.NewPayloadV4(ctx, types.PayloadToExecutableData(&block),
		[]common.Hash{}, common.BytesToHash(block.BeaconRoot), block.Requests)
	if err != nil {
		return err
	}

	if response.Status == engine.INVALID {
		return errors.New("invalid from NewPayloadV4 api")
	}

	// set current block hash to head state and set previous block hash to safe and finalized state
	k.Logger().Info("Notify ForkChoiceUpdated",
		"head", hexutil.Encode(block.BlockHash), "finalized", hexutil.Encode(block.ParentHash))
	parentHash := common.BytesToHash(block.ParentHash)
	forkRes, err := k.engineClient.ForkchoiceUpdatedV3(ctx, &engine.ForkchoiceStateV1{
		HeadBlockHash: common.BytesToHash(block.BlockHash),
		SafeBlockHash: parentHash, FinalizedBlockHash: parentHash,
	}, nil)
	if err != nil {
		return err
	}

	if forkRes.PayloadStatus.Status == engine.INVALID {
		return errors.New("invalid from ForkchoiceUpdatedV3 api")
	}

	return nil
}
