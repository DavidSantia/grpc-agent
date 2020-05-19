# Generate protobuf wrapper files


The file `go_agent.pb.go` is generated as follows
```sh
cd ..
protoc --proto_path=protos --proto_path=third_party --go_out=plugins=grpc:protos go_agent.proto
```
