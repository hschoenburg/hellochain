package starter

import (
	//"github.com/cosmos/cosmos-sdk/types/module"
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

	app := NewAppStarter(appName, logger, dbt)

	require.Equal(t, len(app.Mm.Modules), 7, "app starter should load 7 modules into its manager on init")

}
