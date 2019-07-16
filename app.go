package hellochain

import (
	"os"

	"github.com/cosmos/hellochain/starter"
	"github.com/cosmos/hellochain/x/greeter"

	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/codec"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
)

const appName = "hellochain"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hellocli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hellod")
	ModuleBasics    = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
	KeyGreeter    *sdk.KVStoreKey
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

	appStarter := starter.NewAppStarter(appName, logger, db, cdc, greeter.AppModuleBasic{})

	keyGreeter := sdk.NewKVStoreKey(appName)

	greeterKeeper := greeter.NewKeeper(keyGreeter, appStarter.Cdc)

	var app = &helloChainApp{
		appStarter,
		keyGreeter,
		greeterKeeper,
	}

	greeterMod := greeter.NewAppModule(greeterKeeper)
	fmt.Printf("%v", greeterMod)

	greeterMod.RegisterCodec(cdc)

	app.Mm.Modules[greeterMod.Name()] = greeterMod

	app.MountStore(keyGreeter, sdk.StoreTypeDB)

	app.InitializeStarter()

	return app
}
