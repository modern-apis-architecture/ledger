package service

import (
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction/repository"
)

type TransactionService interface {
	Confirmation(transaction *transaction.Transaction) (*transaction.Record, error)
	Cancellation(transaction *transaction.Transaction) (*transaction.Record, error)
	Reversal(transaction *transaction.Transaction) (*transaction.Record, error)
}

type DefaultTransactionService struct {
	repo repository.TransactionRepository
}

func (ds *DefaultTransactionService) Confirmation(transaction *transaction.Transaction) (*transaction.Record, error) {
	return ds.repo.Register(transaction)
}

func (ds *DefaultTransactionService) Cancellation(transaction *transaction.Transaction) (*transaction.Record, error) {
	return ds.repo.Register(transaction)
}

func (ds *DefaultTransactionService) Reversal(transaction *transaction.Transaction) (*transaction.Record, error) {
	return ds.repo.Register(transaction)
}

func NewDefaultTransactionService(repo repository.TransactionRepository) TransactionService {
	return &DefaultTransactionService{repo: repo}
}
