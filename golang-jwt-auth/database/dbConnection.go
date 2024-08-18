package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	// load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// take the mongo connection url from .env file
	MongoDb := os.Getenv("MONGODB_URL")

	if MongoDb == "" {
		log.Fatal("MONGODB_URL is not set in .env file")
	}
	// Define the MongoDB Server API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Define the MongoDB connection options
	opts := options.Client().ApplyURI(MongoDb).SetServerAPIOptions(serverAPI)

	// Creates a MongoDB client and establishes the connection
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}

	// Determines the timeout of the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verifies the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB successfully! Database : go-auth  is available")
	return client

}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) (*mongo.Collection, error) {
	var Collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return Collection, nil
}
