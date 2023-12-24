package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbConnect() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(getDatabaseURL()).SetSocketTimeout(360 * time.Second)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Unable to connect to database.")
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Unable to ping database.")
		return nil, err
	}

	return client.Database(getDatabaseName()), nil
}

func getDatabaseURL() string {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return "mongodb://127.0.0.1:27017"
	}
	return databaseURL
}

func getDatabaseName() string {
	databaseName := os.Getenv("DATABASE_NAME")
	if databaseName == "" {
		return "scatter-protocol"
	}
	return databaseName
}
