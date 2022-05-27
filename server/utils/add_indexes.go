package utils

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddIndexes(coll *mongo.Collection, indexes []mongo.IndexModel) {
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := coll.Indexes().CreateMany(context.Background(), indexes, opts)

	if err != nil {
		log.Fatal("Error creating indexes")
	}
}
