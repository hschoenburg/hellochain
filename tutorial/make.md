# Makefile

Now lets add a short makefile so we can build our basic app. Open up your `Makefile` and add the following code. Later we will add the command to build our CLI tools as well.

<<< @/tutorial/samples/MakefileShort

Then install your basic blockchain with `make install`.


```bash
--> Ensure dependencies have not been modified
GO111MODULE=on go mod verify
all modules verified
GO111MODULE=on go install -tags "" ./cmd/hcd
```

:::tip
Remember you need to have Go installed and a proper $GOPATH configured


Once installed, start up your blockchain node. Dont worry it won't be able to find seeds.

```bash
$ hcd start
I[2019-08-06|16:59:15.977] Starting ABCI with Tendermint                module=main
E[2019-08-06|16:59:16.005] Couldn't connect to any seeds                module=p2p
I[2019-08-06|16:59:21.019] Executed block                               module=state height=2 validTxs=0 invalidTxs=0
I[2019-08-06|16:59:21.020] Committed state                              module=state height=2 txs=0 appHash=7377248821C962C10C81007882954D749BC65B1F458EFE40A844F78FBBD9F635
I[2019-08-06|16:59:26.029] Executed block                               module=state height=3 validTxs=0 invalidTxs=0
I[2019-08-06|16:59:31.037] Committed state                              module=state height=4 txs=0 appHash=7377248821C962C10C81007882954D749BC65B1F458EFE40A844F78FBBD9F635
I[2019-08-06|16:59:36.047] Executed block                               module=state height=5 validTxs=0 invalidTxs=0
...and watch the blocks roll by!
```

Good job! Now lets add some functionality to that blockchain!
