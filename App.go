package main

import (
	"blourg/service"
	"blourg/utils/security"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	security.ConnectToDB()              // Establish database connection
	router := mux.NewRouter()           // Create a router
	SetRoutes(router)                   // Set the HTTP routes
	corsRouter := security.CORS(router) // Wrap the router with CORS middleware

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", corsRouter))
}

func SetRoutes(router *mux.Router) {
	// Create a new router
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
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("src/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	})
	//
	//router.HandleFunc("/posts/create", func(w http.ResponseWriter, r *http.Request) {
	//	tmpl, err := template.ParseFiles("src/post-form/post-form.html")
	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	tmpl.Execute(w, nil)
	//}

	router.HandleFunc("/users/{userUUID}}", service.GetUserProfile).Methods("POST")
	router.HandleFunc("/my-profile", service.GetUserProfile).Methods("GET")
	router.HandleFunc("/sign-in", service.Signin).Methods("POST")
	router.HandleFunc("/sign-in", service.ServeSignin).Methods("GET")

	router.HandleFunc("/sign-up", service.Signup).Methods("POST")
	router.HandleFunc("/sign-up", service.ServeSignup).Methods("GET")

	router.HandleFunc("/sign-out", service.Signout).Methods("GET")

	router.HandleFunc("/posts", service.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", service.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", service.GetPostById).Methods("GET")
	router.HandleFunc("/posts/{postId}", service.UpdatePost).Methods("POST")
	router.HandleFunc("/posts/{postId}", service.DeletePostById).Methods("DELETE")

}
