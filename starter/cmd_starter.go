package starter

import (
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/auth/genaccounts/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
)

type ServerCommandParams struct {
	DefaultNodeHome string
	DefaultCLIHome  string
	Cdc             *codec.Codec
	CmdName         string
	CmdDesc         string
	ModuleBasics    module.BasicManager
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
		genutilcli.InitCmd(ctx, params.Cdc, params.ModuleBasics, params.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, params.Cdc, genaccounts.AppModuleBasic{}, params.DefaultNodeHome),
		genutilcli.GenTxCmd(ctx, params.Cdc, params.ModuleBasics, staking.AppModuleBasic{}, genaccounts.AppModuleBasic{}, params.DefaultNodeHome, params.DefaultCLIHome),
		genutilcli.ValidateGenesisCmd(ctx, params.Cdc, params.ModuleBasics),
		genaccscli.AddGenesisAccountCmd(ctx, params.Cdc, params.DefaultNodeHome, params.DefaultCLIHome),
	)

	server.AddCommands(ctx, params.Cdc, rootCmd, params.AppCreator, params.AppExporter)
	return rootCmd
}
