package models

import (
	"context"
	"encoding/json"
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

func GetUserById(userId string) User{

	db := config.GetDB()
	collection := db.Collection(userCollection)

	//fmt.Println(collection)

	// steps to get user by id
	id, _ := primitive.ObjectIDFromHex(userId)
	fmt.Println("UserId : ", id)

	filter := bson.D{{Key: "_id", Value: id}}

	
	err := collection.FindOne(context.TODO(), filter).Decode(&bsonString)

	if err != nil{
		log.Println("Failed to find user with userid: ", userId)
		fmt.Println(err)
	}

	var result User
	// marshal & unmarshall
	finalBytes, _ := bson.Marshal(bsonString)
	bson.Unmarshal(finalBytes, &result)
	fmt.Println("RESULT ", result)
	// fmt.Println("Username: ", result.Username)

	return result
}

func GetUserIdByUsername(username string) string {
	// steps to get users id filtered by username from database

	id := ""

	collection := config.GetDB().Collection(userCollection)
	filter := bson.D{{Key: "Username", Value: username}}
	resp := collection.FindOne(context.TODO(), filter).Decode(&jsonString)

	if resp != nil{
		log.Println("Failed to find user with username : ", username)
	}
	res,_ := json.Marshal(resp)

	fmt.Println("RESPONSE : ", res)
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
