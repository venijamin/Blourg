package main

import (
	"backend/security"
	"backend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	security.ConnectToDB()

	// Controller mappings
	router := mux.NewRouter()
	router.HandleFunc("/users", service.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/register", service.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", service.LoginUser).Methods("POST")
	router.HandleFunc("/users/delete", service.DeleteUser).Methods("DELETE")

	router.HandleFunc("/comments", service.CreateComment).Methods("POST")
	router.HandleFunc("/posts", service.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", service.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", service.GetPostById).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
