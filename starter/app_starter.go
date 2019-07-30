package starter

import (
	//"fmt"
	"github.com/cosmos/cosmos-sdk/server"
	. "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	tmtypes "github.com/tendermint/tendermint/types"
	"io"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"

	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	pvm "github.com/tendermint/tendermint/privval"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	//ttbci "github.com/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	tlog "github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/genutil"

	"github.com/cosmos/cosmos-sdk/x/auth/genaccounts"

	"github.com/cosmos/cosmos-sdk/x/bank"
	//distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/params"
	//"github.com/cosmos/cosmos-sdk/x/slashing"
	//"github.com/cosmos/cosmos-sdk/x/staking"
)

var (
	ModuleBasics    module.BasicManager
	Cdc             *codec.Codec
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hellocli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hellod")

//var _ Application = (*BaseApplication)(nil)
)

//AppStarter is a drop in to make simple hello world blockchains
// TODO put starter in a utils/ dir of the tutorials repo

func init() {
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		params.AppModuleBasic{},
		//staking.AppModuleBasic{},
		//distr.AppModuleBasic{},
		//slashing.AppModuleBasic{},
	)

}

type AppStarter struct {
	*bam.BaseApp

	// Keys to access the substores
	keyMain          *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	//keyStaking       *sdk.KVStoreKey
	//tkeyStaking      *sdk.TransientStoreKey
	//keyDistr         *sdk.KVStoreKey
	//tkeyDistr        *sdk.TransientStoreKey
	keyParams  *sdk.KVStoreKey
	tkeyParams *sdk.TransientStoreKey
	//	keySlashing      *sdk.KVStoreKey

	// Keepers
	accountKeeper auth.AccountKeeper
	bankKeeper    bank.Keeper
	//stakingKeeper staking.Keeper
	//slashingKeeper      slashing.Keeper
	//distrKeeper         distr.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	paramsKeeper        params.Keeper
	Cdc                 *codec.Codec
	Mm                  *module.Manager
}

func (app *AppStarter) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	config := server.NewDefaultContext().Config
	config.SetRoot(DefaultNodeHome)

	server.UpgradeOldPrivValFile(config)

	_, _, err := genutil.InitializeNodeValidatorFiles(config)

	privValidator := pvm.LoadOrGenFilePV(
		config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())
	//valPubKey := privValidator.GetPubKey()
	valPubKey := tmtypes.TM2PB.PubKey(privValidator.GetPubKey())

	//fmt.Printf("PUBKEY BABY? %v for NODE: %v \n", valPubKey.String(), nodeID)

	update := abci.ValidatorUpdate{
		PubKey: valPubKey,
		Power:  100}

	// retrieve validator and send lots of voting power?
	// examine how privValidator key is retrieved for the node
	// return  ResponseInitChain
	// then pull out staking, distr, feeCollection, slashing --- maybe bank?
	// look in GenesisInit code
	var genesisState GenesisState

	err = app.Cdc.UnmarshalJSON(req.AppStateBytes, &genesisState)
	if err != nil {
		panic(err)
	}

	genesis := app.Mm.InitGenesis(ctx, genesisState)
	// plop POA validaor into Validators
	genesis.Validators = append(genesis.Validators, update)
	return genesis
}

func (app *AppStarter) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.Mm.BeginBlock(ctx, req)
}
func (app *AppStarter) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.Mm.EndBlock(ctx, req)
}
func (app *AppStarter) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

//_________________________________________________________

func (app *AppStarter) ExportAppStateAndValidators(forZeroHeight bool, jailWhiteList []string,
) (appState json.RawMessage, validators []tmtypes.GenesisValidator, err error) {

	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, abci.Header{Height: app.LastBlockHeight()})

	genState := app.Mm.ExportGenesis(ctx)
	appState, err = codec.MarshalJSONIndent(app.Cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	//validators = staking.WriteValidators(ctx, app.stakingKeeper)

	return appState, validators, nil
}

func NewAppStarter(appName string, logger tlog.Logger, db dbm.DB, cdc *codec.Codec) *AppStarter {

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &AppStarter{
		Cdc:              cdc,
		BaseApp:          bApp,
		keyMain:          sdk.NewKVStoreKey(bam.MainStoreKey),
		keyAccount:       sdk.NewKVStoreKey(auth.StoreKey),
		keyFeeCollection: sdk.NewKVStoreKey(auth.FeeStoreKey),
		//keyStaking:       sdk.NewKVStoreKey(staking.StoreKey),
		//tkeyStaking:      sdk.NewTransientStoreKey(staking.TStoreKey),
		//keyDistr:         sdk.NewKVStoreKey(distr.StoreKey),
		//tkeyDistr:        sdk.NewTransientStoreKey(distr.TStoreKey),
		keyParams:  sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams: sdk.NewTransientStoreKey(params.TStoreKey),
		//keySlashing:      sdk.NewKVStoreKey(slashing.StoreKey),
		Mm: &module.Manager{},
	}

	// The ParamsKeeper handles parameter storage for the application
	app.paramsKeeper = params.NewKeeper(app.Cdc, app.keyParams, app.tkeyParams, params.DefaultCodespace)
	// Set specific supspaces
	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSupspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	//stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	//distrSubspace := app.paramsKeeper.Subspace(distr.DefaultParamspace)
	//slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = auth.NewAccountKeeper(
		app.Cdc,
		app.keyAccount,
		authSubspace,
		auth.ProtoBaseAccount,
	)

	// The BankKeeper allows you perform sdk.Coins interactions
	app.bankKeeper = bank.NewBaseKeeper(
		app.accountKeeper,
		bankSupspace,
		bank.DefaultCodespace,
	)

	// The FeeCollectionKeeper collects transaction fees and renders them to the fee distribution module
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(cdc, app.keyFeeCollection)

	/*
		// The staking keeper
		stakingKeeper := staking.NewKeeper(
			app.Cdc,
			app.keyStaking,
			app.tkeyStaking,
			app.bankKeeper,
			stakingSubspace,
			staking.DefaultCodespace,
		)

		app.distrKeeper = distr.NewKeeper(
			app.Cdc,
			app.keyDistr,
			distrSubspace,
			app.bankKeeper,
			&stakingKeeper,
			app.feeCollectionKeeper,
			distr.DefaultCodespace,
		)

		app.slashingKeeper = slashing.NewKeeper(
			app.Cdc,
			app.keySlashing,
			&stakingKeeper,
			slashingSubspace,
			slashing.DefaultCodespace,
		)

		// register the staking hooks
		// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
		app.stakingKeeper = *stakingKeeper.SetHooks(
			staking.NewMultiStakingHooks(
				app.distrKeeper.Hooks(),
				app.slashingKeeper.Hooks()),
		)
	*/

	app.Mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		//genutil.NewAppModule(app.accountKeeper,
		//app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper, app.feeCollectionKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		//distr.NewAppModule(app.distrKeeper),
	//	slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
	//staking.NewAppModule(app.stakingKeeper, app.feeCollectionKeeper, app.distrKeeper, app.accountKeeper),
	)
	return app
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

func NewDefaultGenesisState() GenesisState {
	return ModuleBasics.DefaultGenesis()
}

func (app *AppStarter) GetCodec() *codec.Codec {
	return app.Cdc
}

func (app *AppStarter) BuildModuleBasics(...moduleBasics []moduleBasic{}) {

		for _, mb := range moduleBasics {
			ModuleBasics[mb.Name()] = mb
		}
	}


func (app *AppStarter) InitializeStarter() {
	//app.Mm.SetOrderBeginBlockers(distr.ModuleName, slashing.ModuleName)
	//app.Mm.SetOrderEndBlockers(staking.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	app.Mm.SetOrderInitGenesis(
		genaccounts.ModuleName,
		//distr.ModuleName,
		//staking.ModuleName,
		auth.ModuleName,
		bank.ModuleName,
		//slashing.ModuleName,
		genutil.ModuleName,
	)

	// register all module routes and module queriers
	app.Mm.RegisterRoutes(app.Router(), app.QueryRouter())

	// The initChainer handles translating the genesis.json file into initial state for the network
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	// The AnteHandler handles signature verification and transaction pre-processing
	app.SetAnteHandler(
		auth.NewAnteHandler(
			app.accountKeeper,
			app.feeCollectionKeeper,
			auth.DefaultSigVerificationGasConsumer,
		),
	)

	app.MountStores(
		app.keyMain,
		app.keyAccount,
		app.keyFeeCollection,
		//app.keyStaking,
		//app.tkeyStaking,
		//app.keyDistr,
		//app.tkeyDistr,
		//app.keySlashing,
		app.keyParams,
		app.tkeyParams,
	)

	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		cmn.Exit(err.Error())
	}
}

func NewAppCreator(creator func(log.Logger, dbm.DB) abci.Application) server.AppCreator {
	return func(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
		app := creator(logger, db)
		return app
	}
}

func NewAppExporter(creator func(log.Logger, dbm.DB) abci.Application) server.AppExporter {
	return func(logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tmtypes.GenesisValidator, error) {

		//App := creator(logger, db)

		return nil, nil, nil

		/*
			TODO not sure how to fix this. LoadHeight and ExportAppState are not on abci.Application
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

var _ abci.Application = AppStarter{}
