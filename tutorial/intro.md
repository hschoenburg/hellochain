# Getting started

We will be building Hellochain, a "Hello World" blockchain. A lot of basic functionality comes packaged for you, not just in the SDK but in the `starter` package we will be using here as well. `starter` will provide basic accounts, a bank, authentication, Tx verification and more. We will implement a `greeter` module that will add the ability so send arbitrary messages (strings) to a given user (account address)

In this tutorial we will create an app with the following file structure.

```bash
./hellochain
├── go.mod
├── Makefile
├── app.go
├── cmd
│   ├── hccli
│   │   └── main.go
│   └── hcd
│       └── main.go
└── x
    └── greeter
        ├── client
        │   ├── cli
        │   │   ├── query.go
        │   │   └── tx.go
        ├── types
            ├── msgs.go
            └── types.go
        ├── handler.go
        ├── keeper.go
        ├── module.go
        └── querier.go

```

Start by creating a new git repository:

```bash
mkdir -p $GOPATH/src/github.com/{ .Username }/hellochain
cd $GOPATH/src/github.com/{ .Username }/hellochain
git init
```

Then initialize your app as a go module:

```bash
cd ./hellochain
go mod init
```

Ok now you are ready to write some code.
