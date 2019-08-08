# Simple Daemon

Now lets set up command to start our chain as well as to interact with it. For this we will build both a "daemon" command and a CLI command.

### hcd

Lets start with our "daemon". Open the file `cmd/hcd/main.go`, this will be your `hellochain daemon` command. For now we will rely on `starter` to bundle things up, but we will come back later to add our own `greeter` functionality when its ready.

Your `./cmd/hcd/main.go` should look like this. We will use it to start and run our node.

<<< @/tutorial/samples/simple-cmd.go

::: tip
The Cosmos SDK uses [cobra])(https://github.com/spf13/cobra) for building and running CLI commands.
:::

### hccli

Next lets start our CLI tool. Open `./cmd/hccli/main.go` and add the following code. Again we will come back to add our greeter feature later.

<<< @/tutorial/samples/simple-cli.go

Next we will create a Makefile for building and installing the package.
