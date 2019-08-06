 # Building and running the application                                                              
                                                                                                     
 ## Building the `nameservice` application                                                           
                                                                                                     
 If you want to build the `nameservice` application in this repo to see the functionalities, **Go 1.1
 2.1+** is required .                                                                                
                                                                                                     
 Add some parameters to environment is necessary if you have never used the `go mod` before.         
                                                                                                     
 ```bash                                                                                             
 mkdir -p $HOME/go/bin                                                                               
 echo "export GOPATH=$HOME/go" >> ~/.bash_profile                                                    
 echo "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile                                                 
 echo "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile                                                
 echo "export GO111MODULE=on" >> ~/.bash_profile                                                     
 source ~/.bash_profile                                                                              
 ```                                                                                                 
                                                                                                     


Now, you can install and run the application.

```bash
# Install the app into your $GOBIN
make install

# Now you should be able to run the following commands:
hcd help
hccli help
```



## Running the live network and using the commands

To initialize configuration and a `genesis.json` file for your application and an account for the transactions, start by running:

> _*NOTE*_: In the below commands addresses are pulled using terminal utilities. You can also just input the raw strings saved from creating keys, shown below. The commands require [`jq`](https://stedolan.github.io/jq/download/) to be installed on your machine.

> _*NOTE*_: If you have run the tutorial before, you can start from scratch with a `hcd unsafe-reset-all` or by deleting both of the home folders `rm -rf ~/.ns*`

> _*NOTE*_: If you have the Cosmos app for ledger and you want to use it, when you create the key with `hccli keys add jack` just add `--ledger` at the end. That's all you need. When you sign, `jack` will be recognized as a Ledger key and will require a device.


```bash
# Initialize configuration files and genesis file
  # moniker is the name of your node
hcd init <moniker> --chain-id namechain


# Copy the `Address` output here and save it for later use
# [optional] add "--ledger" at the end to use a Ledger Nano S
hccli keys add jack

# Copy the `Address` output here and save it for later use
hccli keys add alice

# Add both accounts, with coins to the genesis file
hcd add-genesis-account $(hccli keys show jack -a) 1000nametoken,100000000stake
hcd add-genesis-account $(hccli keys show alice -a) 1000nametoken,100000000stake

# Configure your CLI to eliminate need for chain-id flag
hccli config chain-id namechain
hccli config output json
hccli config indent true
hccli config trust-node true

hcd gentx --name jack <or your key_name>
```

After you have generated a genesis transcation, you will have to input the gentx into the genesis file, so that your nameservice chain is aware of the validators. To do so, run:

`hcd collect-gentxs`

and to make sure your genesis file is correct, run:

`hcd validate-genesis`

You can now start `hcd` by calling `hcd start`. You will see logs begin streaming that represent blocks being produced, this will take a couple of seconds.

You have run your first node successfully.

```bash
# Send a greeting!
hccli query account $(hccli keys show jack -a)
hccli query account $(hccli keys show alice -a)
### Congratulations, you have built a Cosmos SDK application! This tutorial is now complete. If you want to see how to run the same commands using the REST server [click here](run-rest.md).


# Run second node on another machine (Optional)
Open terminal to run commands against that just created to install hcd and hccli
## init use another moniker and same namechain
```bash
hcd init <moniker-2> --chain-id namechain
```

## overwrite ~/.hcd/config/genesis.json with first node's genesis.json

## change persistent_peers
```bash
vim /.hcd/config/config.toml
persistent_peers = "id@firt_node_ip:26656"
run "hccli status" on first node to get id.
```

## start this second node
```bash
hcd start
```

