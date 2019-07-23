package greeter

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the hellochain Querier
const (
	ListGreetings = "list"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case ListGreetings:
			return listGreetings(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown greeter query endpoint")
		}
	}
}

func queryDefault(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	return nil, nil
}

func listGreetings(ctx sdk.Context, params []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {

	fmt.Printf("$$$$$$ Path=======%v =======\n", params)
	greetingList := NewQueryResGreetings()

	iterator := keeper.GetGreetingsIterator(ctx)

	addr, err := sdk.AccAddressFromBech32(params[0])
	if err != nil {
		return nil, sdk.ErrInvalidAddress("invalid address queryparameter")
	}

	for ; iterator.Valid(); iterator.Next() {
		key := string(iterator.Key())

		fmt.Printf("$$$$$$ Key =======%v =======\n", key)
		var greeting Greeting

		keeper.cdc.MustUnmarshalBinaryBare(iterator.Value(), &greeting)

		if greeting.Recipient == addr {
			greetingList[key] = append(greetingList[key], greeting)
			fmt.Printf("$$$$$$ Greeting =======%v =======\n", greeting)
		}
	}

	hellos, err2 := codec.MarshalJSONIndent(keeper.cdc, greetingList)

	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return hellos, nil
}
