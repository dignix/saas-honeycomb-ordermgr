package internal

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoDB(ctx context.Context) (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGODB_URI")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return mongoClient, nil
}
