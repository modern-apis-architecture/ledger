package repository

import "github.com/modern-apis-architecture/ledger/internal/domain/transaction"

type TransactionRepository interface {
	Register(transaction *transaction.Transaction) (*transaction.Record, error)
}
