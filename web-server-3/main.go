package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jmavila/web-server-3/controllers"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Starting server")
	client := connectToDB()
	defer closeDBConnection(client)

	router := httprouter.New()
	fmt.Println("1")

	userControllers := controllers.NewUserController(client)
	fmt.Println("2")
	router.GET("/ping", userControllers.GetPing)
	router.GET("/user/:id", userControllers.GetUser)
	router.POST("/user", userControllers.CreateUser)
	router.PUT("/user/:id", userControllers.UpdateUser)
	router.DELETE("/user/:id", userControllers.DeleteUser)
	fmt.Printf("Starting server at port 8000\n")
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("4")
	fmt.Println("Listening at localhost:8000")
}

func connectToDB() *mongo.Client {
	// Define MongoDB connection string.
	clientOptions := options.Client().ApplyURI("mongodb://admin:123123@localhost:27017/")

	// Connect to MongoDB.
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}

func closeDBConnection(client *mongo.Client) {
	client.Disconnect(context.Background())
}
