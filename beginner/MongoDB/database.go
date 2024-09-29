package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type Movies struct {
	ID         string `json:"id" bson:"_id, omitempty"`
	Name       string `json:"name" bson:"name"`
	Release    int    `json:"date" bson: "Date"`
	Collection string `json:"money" bson: "collection"`
}

var moviesCollection *mongo.Collection

//connect to mongodb

func connectToMongo() (*mongo.Client, error) {
	clientoptions := options.Client().ApplyURI("enter your connection string addrress")
	client, err := mongo.NewClient(clientoptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB!")
	return client, nil

}
func GetMovies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := moviesCollection.Find(ctx, bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrive movies"})
		return
	}
	var movies []Movies
	if err = cursor.All(ctx, &movies); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not decode movies"})
		return
	}
	c.IndentedJSON(http.StatusOK, movies)

}
func main() {
	router := gin.Default()
	client, err := connectToMongo()
	if err != nil {
		panic(err)
	}
	moviesCollection = client.Database("Films").Collection("Movies")
	router.GET("/", GetMovies)
	router.Run("give the path") //example can be local host
}