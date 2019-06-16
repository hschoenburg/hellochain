package starter

import (
	"encoding/json"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// type check to ensure the interface is properly implemented
var (
	_ sdk.AppModule      = BlankModule{}
	_ sdk.AppModuleBasic = BlankModuleBasic{}
)

func NewBlankModule(name string, keeper interface{}) BlankModule {

	var module = BlankModule{
		BlankModuleBasic: BlankModuleBasic{name},
		keeper:           keeper,
	}
	return module
}

func (bm BlankModuleBasic) Name() string {
	return bm.ModuleName
}

func (BlankModuleBasic) RegisterCodec(cdc *codec.Codec) {
	panic("not implemented")
}

// Validation check of the Genesis
func (bm BlankModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	return nil
}

func (bm BlankModuleBasic) DefaultGenesis() json.RawMessage {
	return nil
}

func (bm BlankModule) Name() string {
	return bm.ModuleName
}

func (bm BlankModule) RegisterInvariants(ir sdk.InvariantRouter) {}

func (bm BlankModule) Route() string {
	return bm.ModuleName
}

func (bm BlankModule) NewQuerierHandler() sdk.Querier {
	panic("not implemented")
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
	panic("not implemented")
}

func (bm BlankModule) QuerierRoute() string {
	return bm.ModuleName
}
