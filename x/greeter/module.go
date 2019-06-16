package greeter

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	starter "github.com/cosmos/hellochain/starter"
)

// Gets the entire Whois metadata struct for a name
const ModuleName = "greeter"

type AppModuleBasic struct {
	starter.BlankModuleBasic
}

type AppModule struct {
	AppModuleBasic
	starter.BlankModule
	keeper     Keeper
	ModuleName string
}

func (AppModule) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) QuerierRoute() string {
	return am.ModuleName
}

func (am AppModule) DefaultGenesis() json.RawMessage {
	return nil
}

func (am AppModule) ValidateGenesis(json.RawMessage) error {
	return nil
}

func NewAppModule(keeper Keeper) AppModule {
	blank := starter.NewBlankModule(ModuleName, keeper)
	return AppModule{AppModuleBasic{starter.BlankModuleBasic{ModuleName}}, blank, keeper, ModuleName}
}

// type check to ensure the interface is properly implemented
var (
	_ sdk.AppModule      = AppModule{}
	_ sdk.AppModuleBasic = AppModuleBasic{}
)
