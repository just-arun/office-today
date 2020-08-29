package database

import (
	"context"
	"fmt"
	"log"

	"github.com/just-arun/office-today/internals/boot/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DataBaseConnection holds database client
var DataBaseConnection *mongo.Client

// Init initialize database
func Init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(config.DatabaseHost)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	DataBaseConnection = client

	fmt.Println("Connected to MongoDB!")
}
