//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/ledger/internal/adapter"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/repository"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/service"
	"github.com/modern-apis-architecture/ledger/internal/storage/mongo"
)

func buildAppContainer() (*adapter.LedgerServer, error) {
	wire.Build(mongo.ProvideCollection,
		mongo.NewMongoTransactionRepository,
		service.NewTransactionService,
		wire.Bind(new(repository.TransactionRepository), new(*mongo.MongoTransactionRepository)),
		adapter.NewLedgerServer)
	return nil, nil
}
