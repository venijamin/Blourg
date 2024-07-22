package main

import (
	"backend/security"
	"backend/service"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to database
	// Load the connection string
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("ConnectionString")
	security.OpenConnection(connectionString)

	router := mux.NewRouter()
	router.HandleFunc("/users", service.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/register", service.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", service.LoginUser).Methods("POST")
	router.HandleFunc("/users/delete", service.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
