package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct {
	UserId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username"`
	Userpassword string `json:"userpassword"`
	Usermemail string `json:"useremail"`
}

func CreateUser(userObj User) bool{
	// steps to create a user and add it in database
	fmt.Println("Creating user... ")

	created := true

	return created
}

func GetUserIdByUsername(username string) string{
	// steps to get users id filtered by username from database

	id := ""

	return id
}

func GetUserById(userId string) User{
	// steps to get user by id

	id, _ := primitive.ObjectIDFromHex(userId)
	fmt.Println("Userid : ", id)
	
	userdata := User{
		UserId: primitive.NewObjectID(),
		Username: "username",
		Userpassword: "password",
		Usermemail: "user@email.com",
	}

	return userdata
}

func UpdateUserByUsername(username string) bool{
	// steps to update user info filtered by username

	updateUser := true

	return updateUser
}

func DeleteUserByUsername(username string) bool{
	// steps to delete by user id

	userDeleted :=  true
	return userDeleted
}