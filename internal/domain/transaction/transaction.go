package transaction

import "time"

type Transaction struct {
	Id           string          `json:"id" bson:"_id"`
	LocalTime    time.Time       `json:"local_time" bson:"local_time"`
	RegisteredAt time.Time       `json:"registered_at" bson:"registered_at"`
	CardId       string          `json:"card_id" bson:"card_id"`
	Data         TransactionData `json:"data" bson:"data"`
	Type         string          `json:"type" bson:"type"`
}

type TransactionData struct {
	ExternalId        string  `json:"external_id" bson:"external_id"`
	Value             float32 `json:"value" bson:"value"`
	MerchantId        string  `json:"merchant_id" bson:"merchant_id"`
	CurrencyCode      string  `json:"currency_code" bson:"currency_code"`
	AuthorizationCode string  `json:"authorization_code" bson:"authorization_code"`
	CountryCode       string  `json:"country_code" bson:"country_code"`
}
