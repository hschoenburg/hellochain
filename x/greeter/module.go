package greeter

import (

	//"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	starter "github.com/cosmos/hellochain/starter"
	"github.com/spf13/cobra"
)

// Gets the entire Whois metadata struct for a name
const ModuleName = "greeter"

type AppModuleBasic struct {
	starter.BlankModuleBasic
}

type AppModule struct {
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

func (am AppModule) GetQueryCmd(*codec.Codec) *cobra.Command {
	panic("not implemented")
}

func NewAppModule(keeper Keeper) AppModule {
	blank := starter.NewBlankModule(ModuleName, keeper)
	return AppModule{blank, keeper, ModuleName}
}

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)
