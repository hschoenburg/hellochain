# Getting started

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
mkdir -p $GOPATH/src/github.com/{ .Username }/nameservice
cd $GOPATH/src/github.com/{ .Username }/nameservice
git init
```

Then initialize your go.mod:
```bash
cd ./hellochain
go mod init
```

[Next start your app.go ->](/tutorial/basic_app.md)




