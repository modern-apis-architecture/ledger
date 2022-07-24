package mongo

import (
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepository struct {
	collection *mongo.Collection
}

func (mr *MongoDBRepository) Register(transaction *transaction.Transaction) (*transaction.Register, error) {
	return nil, nil
}

func NewMongoDBRepository(collection *mongo.Collection) transaction.TransactionRepository {
	return &MongoDBRepository{collection: collection}
}
