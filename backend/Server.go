package main

import (
	"backend/security"
	"backend/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// CORS middleware function
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	security.ConnectToDB()

	// Create a new router
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/users", service.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/register", service.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", service.LoginUser).Methods("POST")
	router.HandleFunc("/users/delete", service.DeleteUser).Methods("DELETE")

	router.HandleFunc("/comments", service.CreateComment).Methods("POST")
	router.HandleFunc("/posts", service.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", service.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", service.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postId}", service.DeletePost).Methods("DELETE")
	router.HandleFunc("/posts/{postId}/comments", service.GetAllCommentsForPost).Methods("GET")

	// Wrap the router with CORS middleware
	corsRouter := CORS(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
