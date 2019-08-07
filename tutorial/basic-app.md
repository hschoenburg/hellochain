# Simple App

We will start by assembling and running a very minimal blockchain. First by building a "blank" application and then adding some "Hello World" functionality in the form of a `greeter` module.

You will see below the use of the the package `starter`. This is your "crutch" for this tutorial. It is a heavily configured abstraction for the point of skipping boilerplate and getting something up and running quickly. Later, when you start the nameservice tutorial, you will kick out this "crutch", but for now lets include it.

`app.go` is where you construct your app out of its component modules. `starter` is taking care of most of this for now but we will come back later when its time to add our own application-specific module.

Set up your project with the following in `hellochain/app.go`

<<< @/tutorial/samples/basic-app.go

Great, so lets build and run this app just to make sure its working.
