package services

import (
	"context"
	"log"
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
	log.Println(filter)
	log.Println(collection)
	err := coll.FindOne(ctx, filter).Decode(result)
	log.Println(result)
	if err == mongo.ErrNoDocuments {
		if resultPtr, ok := result.(*interface{}); ok {
			*resultPtr = nil
		}
		return nil
	}
	return err

}

func (s *MongoService) All(collection string, filter interface{}, result *[]interface{}) error {
	coll := s.Client.Database(s.DatabaseName).Collection(collection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc interface{}
		if err := cursor.Decode(&doc); err != nil {
			return err
		}
		*result = append(*result, doc)
	}

	if err := cursor.Err(); err != nil {
		return err

	}
	return nil

}
