package main

import(
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jatink/server-golang/pkg/routes"
)


func main(){
	fmt.Println("This is main function")

	r := mux.NewRouter()

	routes.UserLogin(r)

	fmt.Println("Starting Server")

	log.Fatal(http.ListenAndServe(":4000",r))

	fmt.Println("Server has been started.")
}