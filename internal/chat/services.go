package chat

import "go.mongodb.org/mongo-driver/mongo"

type Services struct {
	Mongo *mongo.Database
	Hub   *Hub
}

func New(mongo *mongo.Database, hub *Hub) *Services {
	return &Services{
		Mongo: mongo,
		Hub:   hub,
	}
}
