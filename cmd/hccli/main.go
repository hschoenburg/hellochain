package main

import (
	app "github.com/cosmos/hellochain"
	starter "github.com/cosmos/hellochain/starter"
	"github.com/tendermint/tendermint/libs/cli"
)

func main() {

	cdc := app.MakeCodec()

	params := starter.CLICommandParams{
		Cdc:     cdc,
		CLIHome: starter.DefaultCLIHome,
	}

	rootCmd := starter.NewCLICommand(params)

	txCmd := starter.TxCmd(cdc)
	queryCmd := starter.QueryCmd(cdc)

	// add more Tx and Query commands
	app.ModuleBasics.AddTxCommands(txCmd, cdc)
	app.ModuleBasics.AddQueryCommands(queryCmd, cdc)
	rootCmd.AddCommand(txCmd, queryCmd)

	executor := cli.PrepareMainCmd(rootCmd, "HC", params.CLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
