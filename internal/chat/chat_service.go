package chat

import (
	"context"
	"time"

	"github.com/oddx-team/odd-game-server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Services) ListService() ([]*Chat, error) {
	chatCollection := s.Mongo.Collection(CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find()
	findOptions.SetSort(map[string]int{"$natural": -1})
	findOptions.SetLimit(50)
	cur, err := chatCollection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var results []*Chat
	for cur.Next(ctx) {
		var result Chat
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, &result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	utils.ReverseAny(results)
	return results, nil
}

func (s *Services) InsertOneService(newChat Chat) (string, error) {
	chatCollection := s.Mongo.Collection(CollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	inserted, err := chatCollection.InsertOne(ctx, newChat)
	if err != nil {
		return "", err
	}

	return (inserted.InsertedID).(string), nil
}
