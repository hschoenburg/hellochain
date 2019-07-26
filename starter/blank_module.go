package starter

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

// app module Basics object
type BlankModuleBasic struct {
	ModuleName string
}
type BlankModule struct {
	BlankModuleBasic
	keeper interface{}
}

type BlankModuleGenesisState []string

func NewBlankModule(name string, keeper interface{}) BlankModule {

	return BlankModule{BlankModuleBasic{name}, keeper}
}

func (bm BlankModuleBasic) Name() string {
	return bm.ModuleName
}

func (BlankModuleBasic) RegisterCodec(cdc *codec.Codec) {
	panic("RegisterCodec not implemented")
}

// Validation check of the Genesis
func (bm BlankModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
	//panic("ValidateGenesis not implemented")
}

func (bm BlankModuleBasic) DefaultGenesis() json.RawMessage {

	data := BlankModuleGenesisState{bm.ModuleName}
	cdc := codec.New()
	return cdc.MustMarshalJSON(data)

}

func (bm BlankModule) Name() string {
	return bm.ModuleName
}

func (bm BlankModule) RegisterInvariants(ir sdk.InvariantRouter) {}

func (bm BlankModule) Route() string {
	return bm.ModuleName
}

func (bm BlankModule) NewQuerierHandler() sdk.Querier {
	panic("NewQuerierHandler not implemented")
}

func (bm BlankModuleBasic) GetQueryCmd(*codec.Codec) *cobra.Command {
	panic("GetQueryCmd not implemented")
}

func (bm BlankModuleBasic) GetTxCmd(*codec.Codec) *cobra.Command {
	panic("GetTxCmd not implemented")
}

// Register rest routes
func (BlankModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router, cdc *codec.Codec) {
	//rest.RegisterRoutes(ctx, rtr, cdc, StoreKey)
	panic("RegisterRESTRoutes not implemented")
}

func (bm BlankModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) sdk.Tags {
	return sdk.EmptyTags()
}

func (bm BlankModule) EndBlock(sdk.Context, abci.RequestEndBlock) ([]abci.ValidatorUpdate, sdk.Tags) {
	return []abci.ValidatorUpdate{}, sdk.EmptyTags()
}

func (bm BlankModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (bm BlankModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	return nil
}

func (bm BlankModule) NewHandler() sdk.Handler {
	panic("NewHandler not implemented")
}

func (bm BlankModule) QuerierRoute() string {
	return bm.ModuleName
}

// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = BlankModule{}
	_ module.AppModuleBasic = BlankModuleBasic{}
)
