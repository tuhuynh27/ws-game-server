package main

import (
	"context"
	"log"
	"net/http"

	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/internal/chat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg := config.Load()
	mongo := mustConnectMongo(cfg)
	hub := chat.NewHub()
	services := chat.New(mongo, hub)

	log.Println("Started at port 5000!")
	log.Fatal(http.ListenAndServe(":5000", NewRouter(services)))
}

func mustConnectMongo(cfg *config.Config) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Mongo.Host))
	if err != nil {
		log.Panic(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Panic(err)
	}

	return client.Database(cfg.Mongo.DatabaseName)
}
