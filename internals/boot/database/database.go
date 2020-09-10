package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/just-arun/office-today/internals/boot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DataBaseConnection holds database client
var DataBaseConnection *mongo.Client

// Init initialize database
func Init() {
	local()
}

func local() {
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

func network() {
	// Replace the uri string with your MongoDB deployment's connection string.
	uri := fmt.Sprintf(config.DatabaseHost, config.DatabaseName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	DataBaseConnection = client

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
}
