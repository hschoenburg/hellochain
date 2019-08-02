package greeter

import (
	"github.com/cosmos/cosmos-sdk/codec"
	gtypes "github.com/cosmos/hellochain/x/greeter/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the greeter Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// GetGreeting returns the latest greeting for a given address
func (k Keeper) GetGreeting(ctx sdk.Context, addr sdk.AccAddress) gtypes.Greeting {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(addr)) {
		return gtypes.Greeting{}
	}
	bz := store.Get([]byte(addr))
	var greeting gtypes.Greeting
	k.cdc.MustUnmarshalBinaryBare(bz, &greeting)
	return greeting
}

// SetGreeting saves a greeeting for a given address.
func (k Keeper) SetGreeting(ctx sdk.Context, greeting gtypes.Greeting) {
	if greeting.Sender.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(greeting.Recipient.Bytes(), k.cdc.MustMarshalBinaryBare(greeting))
}

// GetGreetingsIterator returns  an iterator over all names in which the keys are the addresses and the values are the greetings.
func (k Keeper) GetGreetingsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
