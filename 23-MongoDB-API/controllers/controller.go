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

// get all records from db

func getAllMovies() []bson.M {
	ctx, cancel := getDBContext()
	defer cancel()

	cursor, cursorErr := db.Collection.Find(ctx, bson.M{})

	//A cursor is a mechanism that allows an application to iterate over database results while holding only a subset of them in memory at a given time. Read operations that match multiple documents use a cursor to return those documents in batches rather than all at once.
	if cursorErr != nil {
		log.Fatal("Error finding movies: ", cursorErr)
	}

	var movies []bson.M

	for cursor.Next(ctx){
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal("Error decoding movie: ", err)
		}
		movies = append(movies, movie)

	}

	defer cursor.Close(ctx)

	fmt.Println("Found movies: ", movies)
	return movies
}