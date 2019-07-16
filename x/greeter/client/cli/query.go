package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	g "github.com/cosmos/hellochain/x/greeter/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greeterQueryCmd := &cobra.Command{
		Use:                        "greetings",
		Short:                      "Querying commands for the greeter module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       utils.ValidateCmd,
	}
	greeterQueryCmd.AddCommand(client.GetCommands(
		GetCmdGreetings(storeKey, cdc),
	)...)
	return greeterQueryCmd
}

// GetCmdResolveName queries information about a name
func GetCmdGreetings(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "greetings [addr]",
		Short: "query greetings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/greetings/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("could not find greetings for address - %s \n", addr)
				return nil
			}

			var out g.QueryResGreetings
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
