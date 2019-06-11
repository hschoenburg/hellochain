package greeter

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName

// MsgSetName defines a SetName message
type MsgSayHello struct {
	Greeting  string
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSayHello(body string, addr sdk.AccAddress) MsgSayHello {
	return MsgSsyHello{
		Greeting:  body,
		Sender:    sdk.AccAddress,
		Recipient: sdk.AccAdress,
	}
}

// Route should return the name of the module
func (msg MsgSayHello) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSayHello) Type() string { return "say_hello" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSayHello) ValidateBasic() sdk.Error {
	if msg.Recipient.Empty() {
		return sdk.ErrInvalidAddress(msg.Recipient.String())
	}
	if len(msg.Greeting) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSayHello) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
