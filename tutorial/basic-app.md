# Basic App

We will start by assembling and running a very minimal blockchain. You will see below the use of the the package `starter`. This is your "crutch" for this tutorial. It is a heavily configured abstraction for the point of skipping boilerplate and getting something up and running quickly. Later, when you start the nameservice tutorial, you will kick out this "crutch", but for now lets include it.

`app.go` is where you construct your app out of its component modules. `starter` is taking care of most of this for now but we will come back later when its time to add our own application-specific module.

Set up your project with the following in `hellochain/app.go`

```go
package hellochain

import (
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/hellochain/starter"
)

const appName = "hellochain"

var (
	// ModuleBasics holds the AppModuleBasic struct of all modules included in the app
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
}

// NewHelloChainApp returns a fully constructed SDK application
func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Application {

	appStarter := starter.NewAppStarter(appName, logger, db)

	var app = &helloChainApp{
		appStarter,
	}

	app.InitializeStarter()

	return app
}
```

