package main

import (
	"backend/security"
	"backend/service"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
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

	router.PathPrefix("/src/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".css") {
			// Set the correct MIME type for CSS files
			w.Header().Set("Content-Type", "text/css")

			// Serve the CSS file
			http.StripPrefix("/src/", http.FileServer(http.Dir("src"))).ServeHTTP(w, r)
		} else {
			http.NotFound(w, r) // Return 404 for non-CSS files
		}
	})

	// Define your routes

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("src/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	router.HandleFunc("/posts/create", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("src/post-form/post-form.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})

	router.HandleFunc("/users", service.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/register", service.RegisterUser).Methods("POST")
	router.HandleFunc("/users/login", service.LoginUser).Methods("POST")
	router.HandleFunc("/users/delete", service.DeleteUser).Methods("DELETE")

	router.HandleFunc("/comments", service.CreateComment).Methods("POST")
	router.HandleFunc("/posts", service.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", service.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", service.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postId}", service.UpdatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", service.DeletePostById).Methods("DELETE")
	router.HandleFunc("/posts/{postId}/comments", service.GetAllCommentsForPost).Methods("GET")

	// Wrap the router with CORS middleware
	corsRouter := CORS(router)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}
