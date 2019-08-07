# Simple Daemon

Ok, so even though our blockchain doesn't do anything yet, lets set up a command to start and run it. Open the file `cmd/hcd/main.go`, this will be your `hellochain daemon` command. For now we will rely on `starter` to bundle things up, but we will come back soon to add our own functionality.

Your `cmd/hcd/main.go` should look like this

<<< @/tutorial/samples/basic-cmd.go

::: tip
The Cosmos SDK uses [cobra])(https://github.com/spf13/cobra) for building and running CLI commands.

Next we will create a Makefile for building and installing the package.
