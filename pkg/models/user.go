package models

import (
	"context"
	"fmt"
	"log"

	"github.com/jatink/server-golang/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	bsonString bson.M
	jsonString string
)

const (
	userCollection = "users-1"
)

type User struct {
	UserId       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username     string             `json:"username"`
	Userpassword string             `json:"userpassword"`
	Useremail    string             `json:"useremail"`
}

// func init(){
// 	collection := config.GetCollection(userCollection)
// }

func CreateUser(userObj User) bool {
	// steps to create a user and add it in database
	fmt.Println("Creating user... ")

	created := true

	return created
}

func GetUserById(userId string) User {

	db := config.GetDB()
	collection := db.Collection(userCollection)

	// steps to get user by id
	id, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: id}}

	err := collection.FindOne(context.TODO(), filter).Decode(&bsonString)

	if err != nil {
		log.Println("Failed to find user with userid: ", userId)
		fmt.Println(err)
	}

	var result User
	// marshal & unmarshall
	finalBytes, _ := bson.Marshal(bsonString)
	bson.Unmarshal(finalBytes, &result)

	return result
}

func GetUserIdByUsername(username string) string {
	// steps to get users id filtered by username from database

	id := ""

	fmt.Println("Inside GetUserIdByUsername()")

	collection := config.GetDB().Collection(userCollection)
	fmt.Println("75: Inside GetUserIdByUsername()")
	filter := bson.D{{Key: "Username", Value: username}}
	log.Printf("%v 77: Inside GetUserIdByUsername()", filter)
	resp := collection.FindOne(context.TODO(), filter).Decode(&bsonString)
	fmt.Println("79: Inside GetUserIdByUsername()")

	if resp != nil {
		log.Println("Failed to find user with username : ", username)
	}
	res, _ := bson.Marshal(resp)
	bson.Unmarshal(res, &resp)
	fmt.Println("RESPONSE : ", resp)
	return id
}

func GetAllUsers() []primitive.D {

	fmt.Println("GET ALL USERS")
	db := config.GetDB()
	collection := db.Collection(userCollection)

	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatalln(err)
	}

	var users []primitive.D

	for cursor.Next(context.Background()) {
		var user bson.D
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	fmt.Println(users)

	defer cursor.Close(context.Background())

	return users
}

func UpdateUserByUsername(username string) bool {
	// steps to update user info filtered by username

	updateUser := true

	return updateUser
}

func DeleteUserByUsername(username string) bool {
	// steps to delete by user id

	userDeleted := true
	return userDeleted
}
