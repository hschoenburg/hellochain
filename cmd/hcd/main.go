package main

import (
	"encoding/json"
	"io"

	app "github.com/cosmos/hellochain"
	"github.com/cosmos/hellochain/starter"

	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"
)

func main() {

	cdc := app.MakeCodec()

	params := starter.ServerCommandParams{

		Cdc:          cdc,
		CmdName:      "hcd",
		CmdDesc:      "hellochain AppDaemon",
		ModuleBasics: app.ModuleBasics,
		AppCreator:   newApp,
		AppExporter:  exportAppStateAndValidators,
	}

	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "HC", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

// TODO pull these functions into starter? by just passing app to starter?
// rename to newApp
// add interface to starter - expectedApp interface
// with NewApp
// starter.GetNewApp(app) returns func with new proper signature
// starter.

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewHelloChainApp(logger, db)
}

// same as above GetExporter(app) returns proper function
func exportAppStateAndValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	hcApp := app.NewHelloChainApp(logger, db)

	if height != -1 {
		err := hcApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
	}

	return hcApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
