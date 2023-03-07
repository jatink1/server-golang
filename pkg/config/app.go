package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	bsonString bson.M
	jsonString string
)

const (
	//connectionString = "mongodb+srv://jatink:%40!f3dZm3uyxFNmE@cluster0.iykauql.mongodb.net/?retryWrites=true&w=majority"
	connectionString = "mongodb://localhost:27017/"
	// name of database
	dbName = "tick-tick-clone-db-1"

	// list of collections
)

func GetDB() *mongo.Database {
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	db := client.Database(dbName)

	if err != nil {
		log.Fatal("Failed to connect to database")
		log.Fatal(err)
	}

	databases, _ := client.ListDatabaseNames(context.TODO(), bson.M{})
	log.Printf("%v", databases)
	collections, _ := client.Database(dbName).ListCollectionNames(context.TODO(), bson.M{})
	log.Printf("%v", collections)
	collection := client.Database(dbName).Collection("users-1")
	//filter := bson.M{"t": "1"}
	filter := bson.M{}
	resp, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("42: Inside GetUserIdByUsername()", resp)

	return db
}
