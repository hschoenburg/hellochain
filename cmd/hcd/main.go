package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/auth/genaccounts/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/hellochain/starter"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	app "github.com/cosmos/hellochain"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	tmtypes "github.com/tendermint/tendermint/types"
)

//TODO how to make this shorter?

func main() {

	cdc := app.MakeCodec()

	params := starter.ServerCommandParams{

		DefaultNodeHome: app.DefaultNodeHome,
		DefaultCLIHome:  app.DefaultCLIHome,
		cdc:             *codec.Codec,
		CmdName:         "hcd",
		CmdDesc:         "hellochain AppDaemon",
		ModuleBasics:    app.ModuleBasics,
		App:             app,
		AppCreator:      app.NewHelloChainApp,
		AppExporter:     exportTMStateAndValidators,
	}

	serverCmd := starter.NewServerCommand(params)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(serverCmd, "HC", params.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		hcApp := app.NewHelloChainApp(logger, db)
		err := hcApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return hcApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	hcApp := app.NewHelloChainApp(logger, db)

	return hcApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
