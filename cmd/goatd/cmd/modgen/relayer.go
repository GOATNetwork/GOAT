package modgen

import (
	"bytes"
	"fmt"
	"slices"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/server"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/goatnetwork/goat/x/relayer/types"
	"github.com/spf13/cobra"
)

func Relayer() *cobra.Command {
	const (
		FlagParamElectingPeriod = "param.electing_period"
		FlagPubkey              = "key.tx"
		FlagVoteKey             = "key.vote"
	)

	cmd := &cobra.Command{
		Use:   "relayer",
		Short: "update relayer module genesis",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)

			config := serverCtx.Config.SetRoot(clientCtx.HomeDir)
			genesisFile := config.GenesisFile()

			return UpdateModuleGenesis(genesisFile, types.ModuleName, new(types.GenesisState), clientCtx.Codec, func(genesis *types.GenesisState) error {
				period, err := cmd.Flags().GetDuration(FlagParamElectingPeriod)
				if err != nil {
					return err
				}
				genesis.Params.ElectingPeriod = period
				return nil
			})
		},
	}

	appendVoter := &cobra.Command{
		Use:   "append address",
		Short: "append new voter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)

			config := serverCtx.Config.SetRoot(clientCtx.HomeDir)
			genesisFile := config.GenesisFile()

			addrcodec := clientCtx.TxConfig.SigningContext().AddressCodec()

			addr := args[0]
			addrByte, err := addrcodec.StringToBytes(addr)
			if err != nil {
				return fmt.Errorf("invalid address: %s", addr)
			}

			voteKeyStr, err := cmd.Flags().GetString(FlagVoteKey)
			if err != nil {
				return err
			}

			voteKey, err := DecodeHexOrBase64String(voteKeyStr)
			if err != nil {
				return err
			}

			txKeyStr, err := cmd.Flags().GetString(FlagPubkey)
			if err != nil {
				return err
			}

			txRawKey, err := DecodeHexOrBase64String(txKeyStr)
			if err != nil {
				return err
			}

			if err := IsValidSecp256Pubkey(txRawKey); err != nil {
				return err
			}

			txKey := &secp256k1.PubKey{Key: txRawKey}
			if txKeyAddr := txKey.Address().Bytes(); !bytes.Equal(txKeyAddr, addrByte) {
				addr, _ := addrcodec.BytesToString(txKeyAddr)
				return fmt.Errorf("address and public key not matched, expected address %s", addr)
			}

			serverCtx.Logger.Info("update genesis", "module", types.ModuleName)
			if err := UpdateModuleGenesis(genesisFile, types.ModuleName, new(types.GenesisState), clientCtx.Codec, func(genesis *types.GenesisState) error {
				if _, ok := genesis.Voters[addr]; ok {
					serverCtx.Logger.Info("relayer already added", "addr", addr)
					return nil
				}

				if genesis.Relayer == nil {
					genesis.Relayer = &types.Relayer{
						LastElected:      time.Now().UTC(),
						ProposerAccepted: true,
					}
				}

				genesis.Relayer.Voters = append(genesis.Relayer.Voters, addr)
				slices.Sort(genesis.Relayer.Voters)

				genesis.Relayer.Proposer = genesis.Relayer.Voters[0]
				genesis.Relayer.Voters = genesis.Relayer.Voters[1:]

				genesis.Voters[addr] = &types.Voter{VoteKey: voteKey, Status: types.VOTER_STATUS_ACTIVATED}
				return genesis.Validate()
			}); err != nil {
				return err
			}

			serverCtx.Logger.Info("update genesis", "module", authtypes.ModuleName)
			// Add the relayer account to auth module to allow it sending tx
			return UpdateModuleGenesis(genesisFile, authtypes.ModuleName, new(authtypes.GenesisState), clientCtx.Codec, func(genesis *authtypes.GenesisState) error {
				baseAccount, err := authtypes.NewBaseAccountWithPubKey(txKey)
				if err != nil {
					return err
				}

				if err := genesis.UnpackInterfaces(clientCtx.Codec); err != nil {
					return err
				}

				for _, v := range genesis.GetAccounts() {
					var acc authtypes.GenesisAccount
					if err := clientCtx.Codec.UnpackAny(v, &acc); err != nil {
						return err
					}

					if bytes.Equal(acc.GetAddress().Bytes(), addrByte) {
						return nil
					}
				}

				if err := baseAccount.SetAccountNumber(uint64(len(genesis.GetAccounts()))); err != nil {
					return err
				}

				baseAccountAny, err := codectypes.NewAnyWithValue(baseAccount)
				if err != nil {
					return err
				}
				genesis.Accounts = append(genesis.Accounts, baseAccountAny)
				return nil
			})
		},
	}

	cmd.Flags().Duration(FlagParamElectingPeriod, time.Minute*10, "")
	appendVoter.Flags().String(FlagPubkey, "", "the voter tx public key(compressed secp256k1)")
	appendVoter.Flags().String(FlagVoteKey, "", "the voter vote public key(compressed bls12381 G2)")
	cmd.AddCommand(appendVoter)
	return cmd
}
