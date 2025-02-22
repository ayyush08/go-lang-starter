package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const connectionString string = "mongodb://localhost:27017"

const dbName = "netflix"

const collectionName = "watchlist"

// Most important
var Collection *mongo.Collection

// Connect to MongoDB

func init() { // init is a special function that runs very first time entire application runs

	//connect to MongoDB
	client, err := mongo.Connect(options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//context.Background() returns a new empty context. It is never canceled, has no values, and has no deadline. It is typically used by the main function, tests, and command-line tools.
	//context.WithTimeout returns a copy of parent context whose Done channel is closed as soon as the timeout expires.
	//The returned context's Done channel is closed when the returned cancel function is called or when the timeout expires, whichever happens first.
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB: ", err)
	}

	Collection = client.Database(dbName).Collection(collectionName)
	log.Println("Connected to MongoDB")
}
