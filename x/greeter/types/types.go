package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Whois is a struct that contains all the metadata of a name
type Greeting struct {
	Sender    sdk.AccAddress `json:"sender"`
	Recipient sdk.AccAddress `json:"receiver"`
	Body      string         `json:"body"`
}

// Returns a new Whois with the minprice as the price
func NewGreeting(sender sdk.AccAddress, receiver sdk.AccAddress, body string) Greeting {
	return Greeting{
		Recipient: receiver,
		Sender:    sender,
		Body:      body,
	}
}

// implement fmt.Stringer
func (g Greeting) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Sender: %s
Recipient: %s
Body: %s`, g.Sender, g.Recipient, g.Body))

}
