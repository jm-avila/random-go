package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmavila/web-server-3/controllers"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	loadEnvVars()
	client := connectToDB()
	defer closeDBConnection(client)

	router := httprouter.New()
	controllers.RegisterUserRoutes(client, router)

	addr := os.Getenv("ADDRESS")
	fmt.Printf("Listening at %s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func loadEnvVars() {
	envVarErr := godotenv.Load()
	if envVarErr != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Loaded Environment Variables!")
}

func connectToDB() *mongo.Client {
	mongoUrl := os.Getenv("MONGO_URL")
	clientOptions := options.Client().ApplyURI(mongoUrl)

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
