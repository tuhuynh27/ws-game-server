package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func NewMongoConnection(host, databaseName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		log.Panic(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}

	return client.Database(databaseName)
}
