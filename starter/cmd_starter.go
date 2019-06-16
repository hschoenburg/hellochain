package starter

import (
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/auth/genaccounts/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	abci "github.com/tendermint/tendermint/abci/types"
)

type ServerCommandParams struct {
	DefaultNodeHome string
	DefaultCLIHome  string
	cdc             *codec.Codec
	CmdName         string
	CmdDesc         string
	ModuleBasics    sdk.ModuleBasicManager
	App             abci.Application
	AppCreator      server.AppCreator
	AppExporter     server.AppExporter
}

func NewServerCommand(params ServerCommandParams) *cobra.Command {

	cobra.EnableCommandSorting = false

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               params.CmdName,
		Short:             params.CmdDesc,
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, params.cdc, params.ModuleBasics, params.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, params.cdc, genaccounts.AppModuleBasic{}, params.DefaultNodeHome),
		genutilcli.GenTxCmd(ctx, params.cdc, params.ModuleBasics, staking.AppModuleBasic{}, genaccounts.AppModuleBasic{}, params.DefaultNodeHome, params.DefaultCLIHome),
		genutilcli.ValidateGenesisCmd(ctx, params.cdc, params.ModuleBasics),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		genaccscli.AddGenesisAccountCmd(ctx, params.cdc, params.DefaultNodeHome, params.DefaultCLIHome),
	)

	server.AddCommands(ctx, params.cdc, rootCmd, params.AppCreator, params.AppExporter)
	return rootCmd
}
