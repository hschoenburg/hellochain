package hellochain

import (
	//"github.com/cosmos/cosmos-sdk/types/module"
	//"github.com/stretchr/testify/assert"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	//"github.com/cosmos/cosmos-sdk/types/module"
	"os"
	"testing"
)

func TestNewHelloChainApp(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	dbt := db.NewMemDB()

	app := NewHelloChainApp(logger, dbt)

	//require.Equal(t, len(app.Mm.Modules), 7, "app starter should load 7 modules into its manager on init")
	require.Equal(t, app.KeyGreeter.Name(), "hellochain", "greeter key set to hellochain")

}
