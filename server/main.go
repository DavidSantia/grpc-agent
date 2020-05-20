package main

import (
	"log"
	"net"

	lg "github.com/DavidSantia/grpc-agent/log"
	pb "github.com/DavidSantia/grpc-agent/protos"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	// initialize
	initData()
	lg.InitLog()

	log.Printf("Listening for transactions and logs on %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGoAgentServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve GoAgent: %v", err)
	}
}
