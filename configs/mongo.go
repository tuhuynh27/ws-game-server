package configs

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

const (
	DatabaseName = "odd-game"
)

func InitMongo() {
	connectionString := os.Getenv("MONGO")
	if connectionString == "" {
		connectionString = "mongodb://localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Panic(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}
	db = client.Database(DatabaseName)
}

func GetMongo() *mongo.Database {
	return db
}
