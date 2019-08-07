# CLI

Now that we have implemented `greeter`'s client CLI commands, lets tie them all together into a CLI tool that includes funtionality from all the modules in our chain. We will call this tool `hccli` and its default config dir will be ~/.hccli

Add this to `cmd/hccli/main.go`

<<< @/cmd/hccli/main.go{13}

We call `starter.BuildModuleBasics()` to add  `greeter`. `starter.GetTxCmd` and `starter.GetQueryCmd` collect the Tx and query commands for every module in the ModuleBasicManager (including `greeter`) to assemble a parent command. 


Next we want to build this tool alongside `hcd` so add this line to your Makefile

<<< @/Makefile{4}

Then build your daemon and cli tool.

```bash
$ make install
GO111MODULE=on go install -tags "" ./cmd/hcd
GO111MODULE=on go install -tags "" ./cmd/hccli
```
