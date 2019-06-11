package main

import (
	"fmt"
	"github.com/cosmos/hello-world/x/starter"
	dbm "github.com/tendermint/tendermint/libs/db"
	log "github.com/tendermint/tendermint/libs/log"
)

func main() {

	MyApp := NewHelloWorldApp(log.logger, db)

}
