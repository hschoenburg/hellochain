package greeter

import (
	"fmt"

	//TODO we might be able to skip this entirely and rely on blank module to do everything?

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Greetings []Greeting `json:"greetings"`
}

func NewGenesisState(greetings []Greeting) GenesisState {
	return GenesisState{Greetings: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Greetings {
		if record.Recipient == nil {
			return fmt.Errorf("Invalid Greeting: %s. Error: Missing Recipient", record)
		}
		if record.Sender == nil {
			return fmt.Errorf("Invalid Greeting:  %s. Error: Missing Sender", record)
		}
		if record.Body == "" {
			return fmt.Errorf("Invalid Greeting: Value: %s. Error: Missing Body", record)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Greetings: []Greeting{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Greetings {
		keeper.SetGreeting(ctx, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Greeting
	iterator := k.GetGreetingsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		addr := sdk.AccAddress(iterator.Key())
		var greeting Greeting
		greeting = k.GetGreeting(ctx, addr)
		records = append(records, greeting)
	}
	return GenesisState{Greetings: []Greeting{}}
}
