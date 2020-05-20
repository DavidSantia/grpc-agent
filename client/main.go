// Package main implements a client to test go_agent server
package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/DavidSantia/grpc-agent/protos"
	"google.golang.org/grpc"
)

func main() {
	var address, key, host string

	// Get environment variable settings
	address = os.Getenv("NEW_RELIC_GRPC_ADDRESS")
	if len(address) == 0 {
		address = "localhost:9000"
	}
	key = os.Getenv("NEW_RELIC_LICENSE_KEY")
	if len(key) == 0 {
		log.Fatalf("please set env var NEW_RELIC_LICENSE_KEY")
	}
	host, _ = os.Hostname()

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	name := "hello"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := pb.NewGoAgentClient(conn)
	a, err := c.CreateApp(ctx, &pb.Config{Name: name, License: key, Host: host})
	if err != nil {
		log.Fatalf("could not create app: %v", err)
	}
	log.Printf("App idx: %d", a.GetIdx())

	time.Sleep(time.Second)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	t, err := c.NewTxn(ctx, &pb.NameIndex{Name: "api", Idx: a.GetIdx()})
	if err != nil {
		log.Fatalf("could not create txn: %v", err)
	}
	log.Printf("Txn idx: %d", t.GetIdx())

	time.Sleep(time.Second)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	s, err := c.NewSeg(ctx, &pb.NameIndex{Name: "data", Idx: t.GetIdx()})
	if err != nil {
		log.Fatalf("could not create seg: %v", err)
	}
	log.Printf("Seg idx: %d", s.GetIdx())

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	l, err := c.LogTxn(ctx, &pb.LogIndex{Idx: t.GetIdx(), Level: 3, Message: "It is a very hot day today"})
	if err != nil {
		log.Fatalf("could not log message: %v", err)
	}
	log.Printf("Log idx: %d", l.GetIdx())

	time.Sleep(time.Second)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	e, err := c.EndTxn(ctx, &pb.Index{Idx: t.GetIdx()})
	if err != nil {
		log.Fatalf("could not end txn: %v", err)
	}
	log.Printf("Ended Txn idx: %d", e.GetIdx())
}
