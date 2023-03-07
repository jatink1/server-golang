package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jatink/server-golang/pkg/models"
)

func UserContoller(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside UserController() in controllers package")

	//res := models.GetUserById("63e6728f65eeba6ff1884594")
	//models.GetUserById("jatink")
	//models.GetAllUsers()

	res := models.GetUserIdByUsername("jatink")

	resp, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func Default(w http.ResponseWriter, r *http.Request) {
	fmt.Println("It works!")

	resp, _ := json.Marshal("It works!")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
