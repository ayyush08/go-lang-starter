package controllers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ayyush08/mongoapi/db"
	"github.com/ayyush08/mongoapi/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

//MONGODB helpers - file

func getDBContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

// insert 1 record
// name of fun starts with lowercase so it is private to this package, if we want to use it outside package then we need to export it by starting with uppercase
func insertOneMovie(movie models.Netflix) {
	ctx, cancel := getDBContext()
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

func updateOneMovie(movieId string) {
	ctx, cancel := getDBContext()
	defer cancel()

	id, idErr := bson.ObjectIDFromHex(movieId) //convert string to bson object id accepted by mongodb

	if idErr != nil {
		log.Fatal("Error converting id: ", idErr)
	}

	filter := bson.M{"_id": id} // we can als use bson.D instead of bson.M, the difference is bson.D is ordered and bson.M is unordered but more readable and shorter
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := db.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		log.Fatal("Error updating movie: ", err)
	}

	fmt.Println("Updated movie: ", res)

}


func deleteOneMovie(movieId string) {
	ctx, cancel := getDBContext()
	defer cancel()

	id, idErr := bson.ObjectIDFromHex(movieId)

	if idErr != nil {
		log.Fatal("Error converting id: ", idErr)
	}

	filter := bson.M{"_id": id}

	res,err := db.Collection.DeleteOne(ctx,filter)

	if err != nil {
		log.Fatal("Error deleting movie: ", err)
	}

	fmt.Println("Deleted movie: ", res)
}

func deleteAllMovies() {
	ctx, cancel := getDBContext()
	defer cancel()

	res,err := db.Collection.DeleteMany(ctx,bson.M{}, nil)

	if err != nil {
		log.Fatal("Error deleting movies: ", err)
	}

	fmt.Println("Deleted movies: ", res)
}