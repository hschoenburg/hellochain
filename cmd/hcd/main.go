package main

import (
	app "github.com/cosmos/hellochain"
	"github.com/cosmos/hellochain/starter"

	"github.com/tendermint/tendermint/libs/cli"
)

func main() {

	cdc := app.MakeCodec()

	params := starter.ServerCommandParams{

		Cdc:          cdc,
		CmdName:      "hcd",
		CmdDesc:      "hellochain AppDaemon",
		ModuleBasics: app.ModuleBasics,
		AppCreator:   starter.NewAppCreator(app.NewHelloChainApp),
		// TODO AppExporter is hacky
		AppExporter: starter.NewAppExporter(app.NewHelloChainApp),
	}

	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "HC", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
