package mongo

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/ledger/internal/domain/transaction"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTransactionRepository struct {
	collection *mongo.Collection
}

func (mr *MongoTransactionRepository) Register(t *transaction.Transaction) (*transaction.Record, error) {
	ctx := context.Background()
	opts := options.InsertOne()
	doc, err := mr.collection.InsertOne(ctx, t, opts)
	if err != nil {
		return nil, fmt.Errorf("could not save document to mongo: %w", err)
	}
	id := fmt.Sprintf("%v", doc.InsertedID)
	reg := &transaction.Record{Id: id}
	return reg, nil

}

func NewMongoTransactionRepository(collection *mongo.Collection) *MongoTransactionRepository {
	return &MongoTransactionRepository{collection: collection}
}
