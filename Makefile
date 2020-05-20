all: server/server client/client protos/go_agent.pb.go

server/server: protos/go_agent.pb.go
	@echo Building server
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $@ github.com/DavidSantia/grpc-agent/server/...

client/client: protos/go_agent.pb.go
	@echo Building client
	go build -o $@ github.com/DavidSantia/grpc-agent/client/...

protos/go_agent.pb.go: protos/go_agent.proto
	@echo Generating go_agent.pb.go
	protoc --proto_path=protos --proto_path=third_party --go_out=plugins=grpc:protos go_agent.proto

images: server/server
	docker build -t grpc-agent server
	docker build -t fluentd-nr fluentd
