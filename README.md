# EOSB SDK Backend

[![CircleCI](https://circleci.com/gh/teameosb/eosb-sdk-backend.svg?style=svg)](https://circleci.com/gh/teameosb/eosb-sdk-backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/teameosb/eosb-sdk-backend)](https://goreportcard.com/report/github.com/teameosb/eosb-sdk-backend)

The Eosb SDK is a collection of golang language packages.
You can use it to build a Dapp application backend based on the Eosb contract quickly.
It can help to communicate with Ethereum node, match orders, monitor Ethereum results and so on.
Some general data structures are also provided.

This project cannot be used alone.
You need to add your own application logic.
The following projects are built on top of this SDK.

- [edex](https://github.com/teameosb/edex)




### sdk

The main function of this package is to define the interface to communicate with a blockchain.
We have implemented Ethereum communication codes based on this interface spec.
So as long as the interface is implemented for a blockchain,
eosb SDK backend can be used on top it.This makes it possible to support multi-chain environments easily.

### common

We put some common data structures and interface definitions into this package for sharing with other projects.

### engine


The engine in this package only maintains the orderbook based on the received message
and returns the result of the operation.
It is not responsible for persisting these changes,
nor for pushing messages to users.
Persistent data and push messages are business logic and should be done by the upper application.


### watcher

Blockchain Watcher is responsible for monitoring blockchain changes.
Whenever a new block is generated,
it gets all the transactions in that block.


### websocket

The Websocket package allows you to easily launch a websocket server.
The server is channel based.
Users can join multiple channels and can leave at any time.


There are already a customized channel called `MarketChannel` in this package.
It keep maintaining the newest order book in memory.  
If a new user joins this channel,
it sends a snapshot of current market order book to the user.
After receive a new event from source,
it will update the order book in memory,
then push the change event to all subscribers.

```golang
import (
    github.com/teameosb/eosb-sdk-backend/common
    github.com/teameosb/eosb-sdk-backend/websocket
)

// new a source queue
queue, _ := common.InitQueue(&common.RedisQueueConfig{
    Name:   common.EOSB_WEBSOCKET_MESSAGES_QUEUE_KEY,
    Ctx:    ctx,
    Client: redisClient,
})

// new a websockert server
wsServer := websocket.NewWSServer("localhost:3002", queue)

websocket.RegisterChannelCreator(
    common.MarketChannelPrefix,
    websocket.NewMarketChannelCreator(&websocket.DefaultHttpSnapshotFetcher{
        ApiUrl: os.Getenv("HSK_API_URL"),
    }),
)

// Start the server
// It will block the current process to listen on the `addr` your provided.
wsServer.Start()
```

## License

This project is licensed under the GPL 3.0 - see the [LICENSE](LICENSE) file for details
