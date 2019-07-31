package hellochain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/hellochain/starter"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/hellochain/x/greeter"
	gtypes "github.com/cosmos/hellochain/x/greeter/types"
)

const appName = "hellochain"

var (
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
	greeterKey    *sdk.KVStoreKey
	greeterKeeper greeter.Keeper
}

func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {

	appStarter := starter.NewAppStarter(appName, logger, db, greeter.AppModuleBasic{})

	greeterKey := sdk.NewKVStoreKey(gtypes.StoreKey)

	greeterKeeper := greeter.NewKeeper(greeterKey, appStarter.Cdc)

	var app = &helloChainApp{
		appStarter,
		greeterKey,
		greeterKeeper,
	}

	greeterMod := greeter.NewAppModule(greeterKeeper)

	app.Mm.Modules[greeterMod.Name()] = greeterMod

	app.MountStore(greeterKey, sdk.StoreTypeDB)

	app.InitializeStarter()

	return app
}
