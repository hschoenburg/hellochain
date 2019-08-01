package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO comment
const RouterKey = "greeter"

// MsgSetName defines a SetName message
type MsgSayHello struct {
	Body      string
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSayHello(sender sdk.AccAddress, body string, recipient sdk.AccAddress) MsgSayHello {
	return MsgSayHello{
		Body:      body,
		Sender:    sender,
		Recipient: recipient,
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
	if len(msg.Sender) == 0 || len(msg.Body) == 0 || len(msg.Recipient) == 0 {

		return sdk.ErrUnknownRequest("Sender, Recipient and/or Body cannot be empty")
	}
	return nil
}

// TODO comment
func (msg MsgSayHello) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// GetSignBytes encodes the message for signing
func (msg MsgSayHello) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}
