package main

import (
	"bookmarks-api/app"
	"bookmarks-api/controller"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func getPort() string {
	// get port from .env file
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	fmt.Println("Selected port for server: ", port)
	return port
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controller.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controller.Authenticate).Methods("POST")
	router.HandleFunc("/api/contacts/new", controller.CreateContact).Methods("POST")
	router.HandleFunc("/api/me/contacts", controller.GetContactsFor).Methods("GET")

	// Middleware - check JWT-token
	router.Use(app.JwtAuthentication)

	error := http.ListenAndServe(":"+getPort(), router)
	if error != nil {
		fmt.Print(error)
	}
}
