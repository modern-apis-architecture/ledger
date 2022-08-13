package main

import (
	"github.com/modern-apis-architecture/ledger/internal/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	app, _ := buildAppContainer()
	gs := grpc.NewServer()
	api.RegisterLedgerServiceServer(gs, app)
	if err := gs.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
