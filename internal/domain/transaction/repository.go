package transaction

type TransactionRepository interface {
	Register(transaction *Transaction) (*Register, error)
}
