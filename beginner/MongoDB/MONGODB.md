# MongoDB Integration in Go

This document provides an overview of how to integrate MongoDB with a Go application using the `go.mongodb.org/mongo-driver` package and the Gin web framework.

## Prerequisites

- [Go installed on your machine for mac](https://formulae.brew.sh/formula/go)
- [MongoDB instance running](https://github.com/mongodb/mongo-go-driver)
- [Gin web framework](https://github.com/gin-gonic/gin)

## Setup

1. **Install Dependencies**:
    ```sh
    go get github.com/gin-gonic/gin
    go get go.mongodb.org/mongo-driver/mongo
    go get go.mongodb.org/mongo-driver/mongo/options
    go get go.mongodb.org/mongo-driver/mongo/readpref
    ```

2. **Database Connection**:
    The `connectToMongo` function establishes a connection to the MongoDB instance.
    ```go
    func connectToMongo() (*mongo.Client, error) {
         clientoptions := options.Client().ApplyURI("enter your connection string address")
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
    ```

3. **Define Data Model**:
    Define the `Movies` struct to map the MongoDB documents.
    ```go
    type Movies struct {
         ID         string `json:"id" bson:"_id,omitempty"`
         Name       string `json:"name" bson:"name"`
         Release    int    `json:"date" bson:"Date"`
         Collection string `json:"money" bson:"collection"`
    }
    ```

4. **Retrieve Data**:
    The `GetMovies` function retrieves all movies from the MongoDB collection.
    ```go
    func GetMovies(c *gin.Context) {
         ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
         defer cancel()
         cursor, err := moviesCollection.Find(ctx, bson.M{})
         if err != nil {
              c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "cannot retrieve movies"})
              return
         }
         var movies []Movies
         if err = cursor.All(ctx, &movies); err != nil {
              c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not decode movies"})
              return
         }
         c.IndentedJSON(http.StatusOK, movies)
    }
    ```

5. **Main Function**:
    Set up the Gin router and define the endpoint.
    ```go
    func main() {
         router := gin.Default()
         client, err := connectToMongo()
         if err != nil {
              panic(err)
         }
         moviesCollection = client.Database("Films").Collection("Movies")
         router.GET("/", GetMovies)
         router.Run("give the path") // example can be localhost:8080
    }
    ```

## Running the Application

1. Replace `"enter your connection string address"` with your MongoDB connection string.
2. Replace `"give the path"` with the desired path, e.g., `"localhost:8080"`.
3. Run the application:
    ```sh
    go run database.go
    ```

You should now be able to access the list of movies by navigating to the specified path in your web browser.
