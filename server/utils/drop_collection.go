package utils

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DropCollectionExecutor struct {
	DB *mongo.Database
}

func NewDropCollectionExecutor(db *mongo.Database) DropCollectionExecutor {
	return DropCollectionExecutor{db}
}

func (executor *DropCollectionExecutor) DropCollection(collectionNames []string) {
	for _, name := range collectionNames {
		coll := executor.DB.Collection(name)

		err := coll.Drop(context.Background())
		if err != nil {
			panic(err)
		}
	}
}
