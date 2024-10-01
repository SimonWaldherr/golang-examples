package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Movies represents the structure of a movie document in the MongoDB collection.
// The struct tags `json` and `bson` ensure proper serialization for both JSON and MongoDB BSON formats.
type Movies struct {
	ID         string `json:"id" bson:"_id,omitempty"` // ID field, auto-handled by MongoDB (_id)
	Name       string `json:"name" bson:"name"`        // Name of the movie
	Release    int    `json:"date" bson:"date"`        // Release year of the movie
	Collection string `json:"money" bson:"collection"` // Box office collection
}

// moviesCollection will hold the MongoDB collection object for movies.
// This will be initialized after a successful connection to the database.
var moviesCollection *mongo.Collection

// connectToMongo initializes and connects to the MongoDB instance.
// It returns a MongoDB client object if successful, or an error otherwise.
func connectToMongo() (*mongo.Client, error) {
	// Define options for the MongoDB client, including the connection URI.
	clientOptions := options.Client().ApplyURI("your_connection_string_here")

	// Create a new MongoDB client using the options defined above.
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err // Return the error if client creation fails
	}

	// Create a context with a 10-second timeout for MongoDB operations.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled after the operation completes

	// Connect the client to the MongoDB server.
	err = client.Connect(ctx)
	if err != nil {
		return nil, err // Return the error if the connection fails
	}

	// Ping the MongoDB server to ensure the connection is alive.
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err // Return the error if the ping fails
	}

	fmt.Println("Connected to MongoDB!") // Print a success message on successful connection
	return client, nil                   // Return the MongoDB client
}

// GetMovies is a handler function for the "/movies" endpoint.
// It retrieves all movies from the MongoDB collection and sends them as a JSON response.
func GetMovies(c *gin.Context) {
	// Create a new context with a 10-second timeout to limit the duration of database operations.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled after the operation completes

	// Find all movie documents in the moviesCollection. bson.M{} represents an empty filter (no filtering).
	cursor, err := moviesCollection.Find(ctx, bson.M{})
	if err != nil {
		// If there is an error during the find operation, return a 500 status code and an error message.
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve movies"})
		return
	}
	defer cursor.Close(ctx) // Close the cursor after we're done processing the results

	// Define a slice to hold the list of movies.
	var movies []Movies

	// Decode all the documents from the cursor into the `movies` slice.
	if err = cursor.All(ctx, &movies); err != nil {
		// If decoding fails, return a 500 status code and an error message.
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not decode movies"})
		return
	}

	// Send the movies as a JSON response with a 200 status code.
	c.IndentedJSON(http.StatusOK, movies)
}

// main is the entry point of the application.
func main() {
	// Initialize a new Gin router instance.
	router := gin.Default()

	// Connect to MongoDB using the connectToMongo function.
	client, err := connectToMongo()
	if err != nil {
		// If connection to MongoDB fails, panic and stop the application.
		panic(err)
	}
	defer client.Disconnect(context.TODO()) // Ensure MongoDB client disconnects when the application exits

	//here in this Set the global moviesCollection to point to the "Movies" collection inside the "Films" database.
	moviesCollection = client.Database("name_of_the_database").Collection("name_of_the_collection")
	router.GET("/movies", GetMovies)
	router.Run("use_your_local_host") // use your local host
}
