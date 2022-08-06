//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/repository"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/service"
	"github.com/modern-apis-architecture/ledger/internal/storage/mongo"
)

func buildAppContainer() (service.TransactionService, error) {
	wire.Build(mongo.ProvideCollection,
		mongo.NewMongoDBRepository,
		service.NewDefaultTransactionService,
		wire.Bind(new(repository.TransactionRepository), new(*mongo.MongoDBRepository)),
		wire.Bind(new(service.TransactionService), new(*service.DefaultTransactionService)))
	return nil, nil
}
