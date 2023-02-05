package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jatink/server-golang/pkg/controllers"
)

var UserLogin = func (router *mux.Router)  {
	fmt.Println("Inside UserLogin Router")
	
	router.HandleFunc("/api/v1/login", controllers.UserContoller).Methods("GET")
}