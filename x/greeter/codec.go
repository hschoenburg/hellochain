package greeter

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSayHello{}, "greeter/SayHello", nil)
}

var ModuleCdc = codec.New()
