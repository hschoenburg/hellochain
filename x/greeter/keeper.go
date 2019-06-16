package greeter

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StoreKey to be used when creating the KVStore
const StoreKey = "grtr"

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nameservice Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func (k Keeper) GetGreeting(ctx sdk.Context, addr sdk.AccAddress) Greeting {
	store := ctx.KVStore(k.storeKey)
	if !store.Has([]byte(addr)) {
		return Greeting{}
	}
	bz := store.Get([]byte(addr))
	var greeting Greeting
	k.cdc.MustUnmarshalBinaryBare(bz, &greeting)
	return greeting
}

func (k Keeper) SetGreeting(ctx sdk.Context, greeting Greeting) {
	if greeting.Sender.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(greeting.Recipient, k.cdc.MustMarshalBinaryBare(greeting))
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetGreetingsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
