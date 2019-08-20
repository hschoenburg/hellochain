# Simple App

First lets assemble and run a very minimal blockchain. Most of the code for
this is wrapped into `starter` so the boilerplate is fairly short. We will be
using `starter.AppStarter` which wraps the Cosmos SDK modules `bank`, `auth`,
`params`, `supply`, `genaccounts`, and `genutil` into a minimal app.

`app.go` is where you construct your app out of its component modules.
`starter` is taking care of most of this for now but we will come back later
when its time to add our own application-specific module.

Set up your project with the following in `hellochain/app.go`

<<< @/tutorial/samples/simple-app.go

Next lets build and run this app just to make sure its working and try out some
of the basic functionality.
