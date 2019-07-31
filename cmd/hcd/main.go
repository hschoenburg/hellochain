package main

import (
	app "github.com/cosmos/hellochain"
	"github.com/cosmos/hellochain/starter"
	"github.com/cosmos/hellochain/x/greeter"

	"github.com/tendermint/tendermint/libs/cli"
)

func main() {

	starter.BuildModuleBasics(greeter.AppModuleBasic{})

	params := starter.ServerCommandParams{

		CmdName:     "hcd",
		CmdDesc:     "hellochain AppDaemon",
		AppCreator:  starter.NewAppCreator(app.NewHelloChainApp),
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
