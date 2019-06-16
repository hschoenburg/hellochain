package hellochain

import (
	"encoding/json"
	"os"

	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/cosmos/hellochain/starter"
	"github.com/cosmos/hellochain/x/greeter"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	tlog "github.com/tendermint/tendermint/libs/log"
)

const appName = "hellochain"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hellocli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hellod")
	ModuleBasics    sdk.ModuleBasicManager
)

type helloChainApp struct {
	*starter.AppStarter
	keyGreeter    *sdk.KVStoreKey
	greeterKeeper greeter.Keeper
}

func NewHelloChainApp(logger tlog.Logger, db dbm.DB) *helloChainApp {

	appStarter := starter.NewAppStarter(appName, logger, db, greeter.AppModuleBasic{})

	keyGreeter := sdk.NewKVStoreKey(appName)

	greeter.RegisterCodec(appStarter.Cdc)

	var app = &helloChainApp{
		appStarter,
		keyGreeter,
		greeter.NewKeeper(keyGreeter, appStarter.Cdc),
	}

	app.mm.Modules[greeter.Name()] = greeter.NewAppModule(app.greeterKeeper)

	app.MountStore(keyGreeter, sdk.StoreTypeDB)

	app.InitializeStarter()

	return app
}
