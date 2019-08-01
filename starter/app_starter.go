package starter

import (
	"encoding/json"
	"io"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	pvm "github.com/tendermint/tendermint/privval"
	tmtypes "github.com/tendermint/tendermint/types"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

// nolint
var (
	ModuleBasics    module.BasicManager
	Cdc             *codec.Codec
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hellocli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hellod")
	maccPerms       = map[string][]string{
		auth.FeeCollectorName: nil,
	}
)

//AppStarter is a drop in to make simple hello world blockchains
// TODO put starter in a utils/ dir of the tutorials repo

func init() {
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		supply.AppModuleBasic{},
	)
}

// TODO comment
type AppStarter struct {
	*bam.BaseApp

	// Keys to access the substores
	keyMain    *sdk.KVStoreKey
	keyAccount *sdk.KVStoreKey
	keySupply  *sdk.KVStoreKey
	keyParams  *sdk.KVStoreKey
	tkeyParams *sdk.TransientStoreKey

	// Keepers
	accountKeeper auth.AccountKeeper
	bankKeeper    bank.Keeper
	supplyKeeper  supply.Keeper
	paramsKeeper  params.Keeper
	Cdc           *codec.Codec
	Mm            *module.Manager
}

// TODO comment
var _ abci.Application = AppStarter{}

// TODO comment
func MakeCodec() *codec.Codec {
	cdc := codec.New()
	ModuleBasics.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	Cdc = cdc
	return cdc
}

// TODO comment
func (app *AppStarter) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	config := server.NewDefaultContext().Config
	config.SetRoot(DefaultNodeHome)

	server.UpgradeOldPrivValFile(config)

	_, _, err := genutil.InitializeNodeValidatorFiles(config)
	if err != nil {
		panic(err)
	}

	privValidator := pvm.LoadOrGenFilePV(
		config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())
	valPubKey := tmtypes.TM2PB.PubKey(privValidator.GetPubKey())

	update := abci.ValidatorUpdate{
		PubKey: valPubKey,
		Power:  100}

	var genesisState GenesisState

	err = app.Cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	genesis := app.Mm.InitGenesis(ctx, genesisState)
	genesis.Validators = append(genesis.Validators, update)
	return genesis
}

// TODO comment
func (app *AppStarter) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.Mm.BeginBlock(ctx, req)
}

// TODO comment
func (app *AppStarter) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.Mm.EndBlock(ctx, req)
}

// TODO comment
func (app *AppStarter) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

// TODO comment
func (app *AppStarter) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {

	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	genState := app.Mm.ExportGenesis(ctx)
	appState, err = codec.MarshalJSONIndent(app.Cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	return appState, validators, nil
}

// TODO comment
func BuildModuleBasics(moduleBasics ...module.AppModuleBasic) {

	for _, mb := range moduleBasics {
		ModuleBasics[mb.Name()] = mb
	}
	Cdc = MakeCodec()
}

// TODO comment
func NewAppStarter(appName string, logger log.Logger, db dbm.DB, moduleBasics ...module.AppModuleBasic) *AppStarter {

	BuildModuleBasics(moduleBasics...)

	Cdc = MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(Cdc))

	var app = &AppStarter{
		Cdc:        Cdc,
		BaseApp:    bApp,
		keyMain:    sdk.NewKVStoreKey(bam.MainStoreKey),
		keySupply:  sdk.NewKVStoreKey(supply.StoreKey),
		keyAccount: sdk.NewKVStoreKey(auth.StoreKey),
		keyParams:  sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams: sdk.NewTransientStoreKey(params.TStoreKey),
		Mm:         &module.Manager{},
	}

	app.paramsKeeper = params.NewKeeper(app.Cdc, app.keyParams, app.tkeyParams, params.DefaultCodespace)
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSupspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(
		app.Cdc,
		app.keyAccount,
		authSubspace,
		auth.ProtoBaseAccount,
	)

	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		bankSupspace,
		bank.DefaultCodespace,
	)

	app.supplyKeeper = supply.NewKeeper(
		app.Cdc,
		app.keySupply,
		app.accountKeeper,
		app.bankKeeper,
		supply.DefaultCodespace,
		maccPerms)

	app.Mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
	)
	return app
}

// TODO comment
type GenesisState map[string]json.RawMessage

// TODO comment
func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

// TODO comment
func (app *AppStarter) GetCodec() *codec.Codec {
	return app.Cdc
}

// TODO comment
func (app *AppStarter) InitializeStarter() {

	app.Mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
	)

	app.Mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.supplyKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keySupply,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}
}

// TODO comment
func NewAppCreator(creator func(log.Logger, dbm.DB) abci.Application) server.AppCreator {
	return func(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
		app := creator(logger, db)
		return app
	}
}

// TODO comment
func NewAppExporter(creator func(log.Logger, dbm.DB) abci.Application) server.AppExporter {
	return func(logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tmtypes.GenesisValidator, error) {

		//App := creator(logger, db)

		return nil, nil, nil

		/*
			TODO
			Missing functionality here.
			Not sure how to fix this. LoadHeight and ExportAppState are not on abci.Application interface

				if height != -1 {
					err := App.LoadHeight(height)
					if err != nil {
					}
				}

				json, vals, err := App.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
				return json, vals, err

		*/
	}
}
