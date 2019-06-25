package starter

import (
	"github.com/cosmos/cosmos-sdk/codec"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	"testing"
)

var appName = "testapp"
var modName = "testModule"

func TestNewAppStarter(t *testing.T) {

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	dbt := db.NewMemDB()

	cdc := codec.New()

	app := NewAppStarter(appName, logger, dbt, cdc)

	require.Equal(t, len(app.Mm.Modules), 7, "app starter should load 7 modules into its manager on init")

}
