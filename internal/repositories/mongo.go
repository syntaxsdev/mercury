package repositories

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoClient(connection string) (*mongo.Client, func()) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	client, err := mongo.Connect(options.Client().ApplyURI(connection))
	if err != nil {
		cancel()
		log.Fatalf("Failed to connect to MongoDB client: %v", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		cancel()
		log.Fatalf("Failed to connect to MongoDB client: %v", err)
	}

	cleanup := func() {
		cancel()
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect to MongoDB client: %v", err)
		} else {
			log.Println("MongoDB disconnected successfully")
		}
	}
	return client, cleanup

}
