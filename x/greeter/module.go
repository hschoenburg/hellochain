package greeter

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	starter "github.com/cosmos/hellochain/starter"
	gtypes "github.com/cosmos/hellochain/x/greeter/types"

	"github.com/cosmos/hellochain/x/greeter/client/cli"
)

var (
	// TODO Comment
	ModuleCdc = codec.New()
)

// TODO Comment
type AppModuleBasic struct {
	starter.BlankModuleBasic
}

// TODO Comment
type AppModule struct {
	starter.BlankModule
	keeper     Keeper
	ModuleName string
}

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// TODO Comment
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(gtypes.MsgGreet{}, "greeter/SayHello", nil)
}

// TODO Comment
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// TODO Comment
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// TODO Comment
func (am AppModule) QuerierRoute() string {
	return am.ModuleName
}

// TODO Comment
func (ab AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(gtypes.StoreKey, cdc)

}

// TODO Comment
func (ab AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(gtypes.StoreKey, cdc)
}

// TODO Comment
func NewAppModule(keeper Keeper) AppModule {
	blank := starter.NewBlankModule(gtypes.ModuleName, keeper)
	return AppModule{blank, keeper, gtypes.ModuleName}
}
