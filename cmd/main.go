package main

import (
	"github.com/modern-apis-architecture/ledger/internal/adapter"
	"github.com/modern-apis-architecture/ledger/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	app, _ := buildAppContainer()
	gs := grpc.NewServer()
	api.RegisterLedgerServiceServer(gs, app)
	grpc_health_v1.RegisterHealthServer(gs, adapter.NewHealthServer())
	if err := gs.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
