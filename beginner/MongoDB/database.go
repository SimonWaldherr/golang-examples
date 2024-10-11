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

type Movies struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Name       string `json:"name" bson:"name"`
	Release    int    `json:"date" bson:"date"`
	Collection string `json:"money" bson:"collection"`
}

var moviesCollection *mongo.Collection

// Separate function for creating MongoDB client
func createMongoClient(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Separate function for connecting to MongoDB
func connectToMongo(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Connect(ctx)
}

// Function to test connection with MongoDB
func pingMongo(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return client.Ping(ctx, readpref.Primary())
}

// Function to initialize MongoDB collection
func initializeMongoCollection(client *mongo.Client, dbName, collectionName string) {
	moviesCollection = client.Database(dbName).Collection(collectionName)
}

// Handler function to get movies from MongoDB
func GetMovies(c *gin.Context) {
	movies, err := fetchMoviesFromDB()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve movies"})
		return
	}
	c.IndentedJSON(http.StatusOK, movies)
}

// Fetch movies from the database
func fetchMoviesFromDB() ([]Movies, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := moviesCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []Movies
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

func main() {
	router := gin.Default()

	client, err := createMongoClient("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	err = connectToMongo(client)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())

	err = pingMongo(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	initializeMongoCollection(client, "name_of_the_database", "name_of_the_collection")

	router.GET("/movies", GetMovies)
	router.Run("localhost:8080")
}
