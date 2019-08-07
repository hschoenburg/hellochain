package hellochain

import (
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/hellochain/starter"
)

const appName = "hellochain"

var (
	// ModuleBasics holds the AppModuleBasic struct of all modules included in the app. We will add to it later
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
}

// NewHelloChainApp returns a fully constructed SDK application
func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {

	appStarter := starter.NewAppStarter(appName, logger, db)

	var app = &helloChainApp{
		appStarter,
	}

	app.InitializeStarter()

	return app
}
