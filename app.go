package hellochain

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/hellochain/starter"
	"github.com/cosmos/hellochain/x/greeter"
	types "github.com/cosmos/hellochain/x/greeter/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

const appName = "hellochain"

var (
	ModuleBasics = starter.ModuleBasics
)

func init() {
	ModuleBasics[types.ModuleName] = greeter.AppModuleBasic{}
}

type helloChainApp struct {
	*starter.AppStarter
	greeterKey    *sdk.KVStoreKey
	greeterKeeper greeter.Keeper
}

func MakeCodec() *codec.Codec {
	cdc := codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

func NewHelloChainApp(logger log.Logger, db dbm.DB) *helloChainApp {

	cdc := MakeCodec()

	appStarter := starter.NewAppStarter(appName, logger, db, cdc)

	greeterKey := sdk.NewKVStoreKey(types.StoreKey)

	greeterKeeper := greeter.NewKeeper(greeterKey, appStarter.Cdc)

	var app = &helloChainApp{
		appStarter,
		greeterKey,
		greeterKeeper,
	}

	greeterMod := greeter.NewAppModule(greeterKeeper)

	app.Mm.Modules[greeterMod.Name()] = greeterMod

	app.InitializeStarter()

	app.MountStore(greeterKey, sdk.StoreTypeDB)

	return app
}
