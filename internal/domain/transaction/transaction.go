package transaction

import "time"

type Transaction struct {
	Id           string
	LocalTime    time.Time
	RegisteredAt time.Time
	CardId       string
	Data         TransactionData
	Type         string
}

type TransactionData struct {
	ExternalId string
	Value      float64
	MerchantId string
	PosId      string
}
