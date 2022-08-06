//go:build wireinject
// +build wireinject

package main

import (
	"github.com/modern-apis-architecture/ledger/internal/adapter"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/repository"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/service"
	"github.com/modern-apis-architecture/ledger/internal/storage/mongo"
)

func buildAppContainer() (*adapter.LedgerServer, error) {
	wire.Build(mongo.ProvideCollection,
		mongo.NewMongoDBRepository,
		service.NewDefaultTransactionService,
		wire.Bind(new(repository.TransactionRepository), new(*mongo.MongoDBRepository)),
		wire.Bind(new(service.TransactionService), new(*service.DefaultTransactionService))),
		adapter.NewLedgerServer
	return nil, nil
}
