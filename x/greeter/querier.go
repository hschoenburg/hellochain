package greeter

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the hellochain Querier
const (
	QueryGreetings = "greeting"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryGreetings:
			return queryGreetings(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown nameservice query endpoint")
		}
	}
}

func queryGreetings(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {

	var greetingList QueryResGreetings

	iterator := keeper.GetGreetingsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		recipient := string(iterator.Key())
		var greeting Greeting
		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &greeting)
		greetingList[recipient] = append(greetingList[recipient], greeting)
	}

	hellos, err2 := codec.MarshalJSONIndent(keeper.cdc, greetingList)

	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return hellos, nil
}
