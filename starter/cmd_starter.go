package starter

import (
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/auth/genaccounts/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	amino "github.com/tendermint/go-amino"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/lcd"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/types/module"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bankcmd "github.com/cosmos/cosmos-sdk/x/bank/client/cli"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	stakingClient "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"
)

const (
	storeAcc = "acc"
)

func NewCLICommand() *cobra.Command {

	cobra.EnableCommandSorting = false

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	rootCmd := &cobra.Command{
		Use:   "hccli",
		Short: "hellochain Client",
	}

	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	// Construct Root Command
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		client.ConfigCmd(DefaultCLIHome),
		client.LineBreak,
		lcd.ServeCommand(Cdc, RegisterRoutes),
		client.LineBreak,
		keys.Commands(),
		client.LineBreak,
	)
	return rootCmd

}

func RegisterRoutes(rs *lcd.RestServer) {
	rs.CliCtx = rs.CliCtx.WithAccountDecoder(rs.Cdc)
	rpc.RegisterRPCRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterTxRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeAcc)
	bank.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	stakingClient.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
}

func QueryCmd(cdc *amino.Codec) *cobra.Command {
	queryCmd := &cobra.Command{
		Use:     "query",
		Aliases: []string{"q"},
		Short:   "Querying subcommands",
	}

	queryCmd.AddCommand(
		rpc.ValidatorCommand(cdc),
		rpc.BlockCommand(),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
		client.LineBreak,
		authcmd.GetAccountCmd(cdc),
	)

	return queryCmd
}

func TxCmd(cdc *amino.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "tx",
		Short: "Transactions subcommands",
	}

	txCmd.AddCommand(
		bankcmd.SendTxCmd(cdc),
		client.LineBreak,
		authcmd.GetSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		client.LineBreak,
		tx.GetBroadcastCommand(cdc),
		tx.GetEncodeCommand(cdc),
	)

	return txCmd
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}

///////////////////////////////////////////////////////////////////////////////

type ServerCommandParams struct {
	//Cdc          *codec.Codec
	CmdName string
	CmdDesc string
	//ModuleBasics module.BasicManager
	AppCreator  server.AppCreator
	AppExporter server.AppExporter
}

func NewServerCommand(params ServerCommandParams) *cobra.Command {

	cobra.EnableCommandSorting = false

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()

	cdc := MakeCodec()

	rootCmd := &cobra.Command{
		Use:               params.CmdName,
		Short:             params.CmdDesc,
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(
		genutilcli.InitCmd(ctx, cdc, ModuleBasics, DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, DefaultNodeHome),
		genutilcli.GenTxCmd(ctx, cdc, ModuleBasics, staking.AppModuleBasic{}, genaccounts.AppModuleBasic{}, DefaultNodeHome, DefaultCLIHome),
		genutilcli.ValidateGenesisCmd(ctx, cdc, ModuleBasics),
		genaccscli.AddGenesisAccountCmd(ctx, cdc, DefaultNodeHome, DefaultCLIHome),
	)

	server.AddCommands(ctx, cdc, rootCmd, params.AppCreator, params.AppExporter)
	return rootCmd
}
