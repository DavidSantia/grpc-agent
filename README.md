# grpc-agent

gRPC wrapper for New Relic Go Agent

Provides following functions for the 10 programming languages supported by gRPC.
This includes C++ and C#, so you can use the same functions for native Windows support.
```
  rpc CreateApp (Config) returns (Index) {};
  rpc NewTxn (NameIndex) returns (Index) {};
  rpc NewSeg (NameIndex) returns (Index) {};
  rpc EndSeg (Index) returns (Index) {};
  rpc EndTxn (Index) returns (Index) {};
  rpc LogTxn (LogIndex) returns (Index) {};
```

## Build server and client
```sh
make
```

## Build docker images for grpc-agent and log Docker forwarding
First, update the file `fluentd/fluentd-nr.conf` by adding your Insights Insert API key

Then build the images
```sh
make images
```

## Run fluentd and grpc-agent server
```sh
./run.sh
```

## Run the test client
The test client is built in Go, but you can use `protos/go_agent.proto` with all the languages gRPC supports.

First set env var `NEW_RELIC_LICENSE_KEY` then run the program:
```sh
export NEW_RELIC_LICENSE_KEY=YOUR_LICENSE_KEY
./client/client
```

