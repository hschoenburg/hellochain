package greeter

import (
	"github.com/cosmos/hellochain/x/greeter/types"
)

// TODO can we avoid usinf this file? Like add the prefix everywhere?

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgSayHello       = types.NewMsgSayHello
	NewQueryResGreetings = types.NewQueryResGreetings
	NewGreeting          = types.NewGreeting
	ModuleCdc            = types.ModuleCdc
	RegisterCodec        = types.RegisterCodec
)

type (
	MsgSayHello       = types.MsgSayHello
	Greeting          = types.Greeting
	QueryResGreetings = types.QueryResGreetings
)
