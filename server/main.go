package main

import (
	"context"
	"log"
	"net"

	pb "github.com/DavidSantia/grpc-agent/protos"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

// Server wraps Go Agent to send transactions and logs
type server struct {
	pb.UnimplementedCreateAppServer
}

// implement CreateApp
func (s *server) CreateApp(ctx context.Context, in *pb.AppConfig) (*pb.AppReply, error) {
	log.Printf("CreateApp(%q)", in.GetName())
	return &pb.AppReply{Message: "Created App " + in.GetName()}, nil
}

func main() {
	log.Printf("Listening for transactions and logs on %s", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCreateAppServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
