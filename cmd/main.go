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

	gs := grpc.NewServer()
	api.RegisterLedgerServiceServer(gs)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
