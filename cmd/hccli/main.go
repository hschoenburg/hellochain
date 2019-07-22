package main

import (
	"github.com/cosmos/cosmos-sdk/client/lcd"
	app "github.com/cosmos/hellochain"
	starter "github.com/cosmos/hellochain/starter"
	greeter "github.com/cosmos/hellochain/x/greeter"
	grest "github.com/cosmos/hellochain/x/greeter/client/rest"
	"github.com/tendermint/tendermint/libs/cli"
)

const (
	storeGr = "greeter"
)

func main() {

	cdc := app.MakeCodec()

	params := starter.CLICommandParams{
		RegisterRoutes: registerRoutes,
		Cdc:            cdc,
		CLIHome:        app.DefaultCLIHome,
	}

	rootCmd := starter.NewCLICommand(params)

	txCmd := starter.TxCmd(cdc)
	queryCmd := starter.QueryCmd(cdc)

	rootCmd.AddCommand(txCmd, queryCmd)

	// add more Tx and Query commands
	app.ModuleBasics["greeter"] = greeter.AppModuleBasic{}
	app.ModuleBasics.AddTxCommands(txCmd, cdc)
	app.ModuleBasics.AddQueryCommands(queryCmd, cdc)

	executor := cli.PrepareMainCmd(rootCmd, "HC", app.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func registerRoutes(rs *lcd.RestServer) {
	starter.RegisterRoutes(rs)
	grest.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, storeGr)
}
