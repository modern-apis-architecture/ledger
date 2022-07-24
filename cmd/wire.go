package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction"
	"github.com/modern-apis-architecture/ledger/internal/storage/mongo"
)

func buildAppContainer() (transaction.TransactionService, error) {
	wire.Build(mongo.ProvideCollection,
		mongo.NewMongoDBRepository,
		transaction.NewDefaultTransactionService,
		wire.Bind(new(transaction.TransactionRepository), new(mongo.MongoDBRepository)),
		wire.Bind(new(transaction.TransactionService), new(transaction.DefaultTransactionService)))
	return nil, nil
}
