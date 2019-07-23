package greeter

import (

	//"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	starter "github.com/cosmos/hellochain/starter"
	"github.com/cosmos/hellochain/x/greeter/client/cli"
	"github.com/spf13/cobra"
)

type AppModuleBasic struct {
	starter.BlankModuleBasic
}

func Name(ab AppModuleBasic) string {
	return ModuleName
}

type AppModule struct {
	starter.BlankModule
	keeper     Keeper
	ModuleName string
}

func (AppModule) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
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

func (ab AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)

}

func (ab AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
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
