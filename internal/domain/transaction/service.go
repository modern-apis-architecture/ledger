package transaction

type TransactionService interface {
	Confirmation(transaction *Transaction) (*Register, error)
	Cancellation(transaction *Transaction) (*Register, error)
	Reversal(transaction *Transaction) (*Register, error)
}

type DefaultTransactionService struct {
	repo TransactionRepository
}

func (ds *DefaultTransactionService) Confirmation(transaction *Transaction) (*Register, error) {
	return nil, nil
}

func (ds *DefaultTransactionService) Cancellation(transaction *Transaction) (*Register, error) {
	return nil, nil
}

func (ds *DefaultTransactionService) Reversal(transaction *Transaction) (*Register, error) {

	return nil, nil
}

func NewDefaultTransactionService(repo TransactionRepository) TransactionService {
	return &DefaultTransactionService{repo: repo}
}
