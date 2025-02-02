package keeper

import (
	"context"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (k Keeper) DequeueLockingModuleTx(ctx context.Context) ([]*ethtypes.Transaction, error) {
	const MaxTx = 16

	queue, err := k.EthTxQueue.Get(ctx)
	if err != nil {
		return nil, err
	}

	if len(queue.Rewards) == 0 && len(queue.Unlocks) == 0 {
		return nil, nil
	}

	txNonce, err := k.EthTxNonce.Peek(ctx)
	if err != nil {
		return nil, err
	}

	var txs []*ethtypes.Transaction

	// distributeReward txs
	{
		var n int
		for ; n < len(queue.Rewards) && n < MaxTx; n++ {
			dist := queue.Rewards[n]
			txs = append(txs, dist.EthTx(txNonce))
			txNonce++
		}
		queue.Rewards = queue.Rewards[n:]
	}

	// completeUnlock txs
	{
		var n int
		for ; n < len(queue.Unlocks) && n < MaxTx; n++ {
			unlock := queue.Unlocks[n]
			txs = append(txs, unlock.EthTx(txNonce))
			txNonce++
		}
		queue.Unlocks = queue.Unlocks[n:]
	}

	if err := k.EthTxQueue.Set(ctx, queue); err != nil {
		return nil, err
	}
	if err := k.EthTxNonce.Set(ctx, txNonce); err != nil {
		return nil, err
	}

	return txs, nil
}
