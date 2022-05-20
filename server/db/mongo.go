package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func New(uri, dbName string) *mongo.Database {
	if uri == "" || dbName == "" {
		log.Fatal("You must set your 'MONGODB_URI' and 'DB_NAME' environmental variable")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to MongoDB using %s database", dbName)

	return client.Database(dbName)
}

func Disconnect(client *mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connection to MongoDB closed.")
}
