package greeter

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSayHello:
			return handleMsgSayHello(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSayHello(ctx sdk.Context, keeper Keeper, msg MsgSayHello) sdk.Result {
	if msg.Recipient == nil {
		return sdk.ErrUnauthorized("Missing Recipient").Result() // If not, throw an error
	}

	greeting := NewGreeting(msg.Sender, msg.Recipient, msg.Body)

	keeper.SetGreeting(ctx, greeting)

	return sdk.Result{}
}
