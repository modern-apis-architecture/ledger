package service

import (
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/repository"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func (ds *TransactionService) Confirmation(transaction *transaction.Transaction) (*transaction.Record, error) {
	return ds.repo.Register(transaction)
}

func (ds *TransactionService) Cancellation(transaction *transaction.Transaction) (*transaction.Record, error) {
	return ds.repo.Register(transaction)
}

func (ds *TransactionService) Reversal(transaction *transaction.Transaction) (*transaction.Record, error) {
	return ds.repo.Register(transaction)
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}
