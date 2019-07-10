package starter

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	"testing"
)

func TestNewServerCmd(t *testing.T) {

	manager := module.NewModuleBasicManager()

	testParams := ServerCommandParams{

		DefaultNodeHome: "someHomeDir",
		DefaultCLIHome:  "someOtherDir",
		Cdc:             codec.New(),
		CmdName:         "hcd",
		CmdDesc:         "hellochain AppDaemon",
		ModuleBasics:    manager,
		AppCreator:      dummyAppCreator,
		AppExporter:     dummyAppExporter,
	}

	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	dbt := db.NewMemDB()

	cdc := codec.New()

	cmd, err := NewServerCommand(testParams)
	require.Equal(t, err, nil)

	require.IsType(t, cobra.Command, cmd)

	//app := NewAppStarter(appName, logger, dbt, cdc)

	//require.Equal(t, len(app.Mm.Modules), 7, "app starter should load 7 modules into its manager on init")

}

func dummyAppCreator() bool {
	return true
}

func dummyAppExporter() bool {
	return true
}
