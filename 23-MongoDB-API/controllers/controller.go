package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ayyush08/mongoapi/db"
	"github.com/ayyush08/mongoapi/models"
)

//MONGODB helpers - file

//insert 1 record

func insertOneMovie(movie models.Netflix) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Prevents hanging requests by setting a timeout of 5 seconds.
	// Ensures the request is canceled if MongoDB takes too long to respond.
	// Improves performance and resource management by avoiding indefinite resource consumption.
	// Handles network latency and slow queries by setting a timeout.

	insertedMovie, err := db.Collection.InsertOne(ctx, movie)

	if err != nil {
		log.Fatal("Error inserting movie: ", err)
	}

	fmt.Println("Inserted movie: ", insertedMovie.InsertedID)
}
