package config

import (
	"context"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb+srv://jatink:%40!f3dZm3uyxFNmE@cluster0.iykauql.mongodb.net/?retryWrites=true&w=majority"

	// name of database
	dbName = "tick-tick-clone-db-1"

	// list of collections
)

func GetDB() *mongo.Database{
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	db := client.Database(dbName)

	if err != nil {
		log.Fatal("Failed to connect to database")
		log.Fatal(err)
	}


	return db
}