package hellochain

// TODO organize import statements into blocks.
// stdlib at top, then anything from 3rd party, then tendermint, then cosmos

import (
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/hellochain/starter"
	"github.com/cosmos/hellochain/x/greeter"

	gtypes "github.com/cosmos/hellochain/x/greeter/types"
)

const appName = "hellochain"

var (
	ModuleBasics = starter.ModuleBasics
)

// hacky?

//BuildModuleBasics(greeter.appModuleBasic{})
// TODO put starter in separate repo?

func init() {
	ModuleBasics[gtypes.ModuleName] = greeter.AppModuleBasic{}
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
