package collections

import (
	"github.com/just-arun/office-today/internals/boot/config"
	"github.com/just-arun/office-today/internals/boot/database"
	"go.mongodb.org/mongo-driver/mongo"
)

// User connection for database
func User() *mongo.Collection {
	return database.DataBaseConnection.
		Database(config.DatabaseName).
		Collection("users")
}

// Post connection for database
func Post() *mongo.Collection {
	return database.DataBaseConnection.
		Database(config.DatabaseName).
		Collection("posts")
}


