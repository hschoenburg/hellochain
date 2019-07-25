package cli

import (
	//"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	. "github.com/cosmos/hellochain/x/greeter/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greetingTxCmd := &cobra.Command{
		Use:                        "greeter",
		Short:                      "greeter transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       utils.ValidateCmd,
	}

	greetingTxCmd.AddCommand(client.PostCommands(
		GetCmdSayHello(cdc),
	)...)

	return greetingTxCmd
}

// GetCmdBuyName is the CLI command for sending a BuyName transaction
func GetCmdSayHello(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "say [body] [addr]",
		Short: "send a greeting to another user. Usage: say [body] [address]",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			sender := cliCtx.GetFromAddress()
			body := args[0]
			recipient, err := sdk.AccAddressFromBech32(args[1])

			if err != nil {
				return err
			}

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			msg := NewMsgSayHello(sender, body, recipient)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
