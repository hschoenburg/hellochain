package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	. "github.com/cosmos/hellochain/x/greeter/types"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greeterQueryCmd := &cobra.Command{
		Use:                        ModuleName,
		Short:                      "Querying commands for the greeter module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       utils.ValidateCmd,
	}
	greeterQueryCmd.AddCommand(client.GetCommands(
		GetCmdListGreetings(storeKey, cdc),
	)...)
	return greeterQueryCmd
}

// GetCmdResolveGreetings queries all greetings
func GetCmdListGreetings(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list [addr]",
		Short: "list greetings for address. Usage list [address]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			addr := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/list/%s", queryRoute, addr), nil)
			if err != nil {
				fmt.Printf("%v \n could not find greetings for address - %s \n", addr, err)

				return nil
			}

			out := NewQueryResGreetings()
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
