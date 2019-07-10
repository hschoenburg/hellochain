package main

import (
	"github.com/cosmos/cosmos-sdk/client/lcd"
	starter "github.com/cosmos/hellochain/starter"
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/cosmos/hellochain"
	grest "github.com/cosmos/hellochain/x/greeter/client/rest"
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

	// add more Tx and Query commands
	app.ModuleBasics.AddTxCommands(starter.TxCmd(cdc), cdc)
	app.ModuleBasics.AddQueryCommands(starter.QueryCmd(cdc), cdc)

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
