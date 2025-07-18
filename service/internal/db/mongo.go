package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo(uri string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// sanity ping
	if err := client.Database("admin").RunCommand(ctx, map[string]any{"ping": 1}).Err(); err != nil {
		panic(err)
	}

	log.Println("[db] Connected to MongoDB")

	db := client.Database("c9web")

	return db
}
