package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MongoService struct {
	Client       *mongo.Client
	DatabaseName string
}

func (s *MongoService) Insert(collection string, item interface{}) (interface{}, error) {
	coll := s.Client.Database(s.DatabaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := coll.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil

}

func (s *MongoService) First(collection string, filter interface{}, result interface{}) error {
	coll := s.Client.Database(s.DatabaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := coll.FindOne(ctx, filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		if resultPtr, ok := result.(*interface{}); ok {
			*resultPtr = nil
		}
		return nil
	}
	return err

}

func (s *MongoService) All(collection string, filter interface{}, result interface{}) error {
	coll := s.Client.Database(s.DatabaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	// Use the cursor to decode directly into the provided result slice
	if err := cursor.All(ctx, result); err != nil {
		return err
	}

	return nil
}
