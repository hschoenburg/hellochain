package greeter

import (
	//"fmt"

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
	/*
		for _, record := range data.WhoisRecords {
			if record.Owner == nil {
				return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
			}
			if record.Value == "" {
				return fmt.Errorf("Invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
			}
			if record.Price == nil {
				return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
			}
		}
	*/
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
	/*
		var records []Greetings
		iterator := k.GetGreetingsIterator(ctx)
		for ; iterator.Valid(); iterator.Next() {
			name := string(iterator.Key())
			var whois Whois
			whois = k.GetWhois(ctx, name)
			records = append(records, whois)
		}
	*/
	return GenesisState{Greetings: []Greeting{}}
}
