// advanced_routing.go
// Advanced gorilla/mux routing features: regex patterns, query parameters, and route matching
//
// Gorilla Mux provides powerful routing capabilities beyond basic path matching,
// including regular expressions, query parameter constraints, host matching,
// and custom matchers for sophisticated URL routing.
//
// I AM NOT DONE YET!

package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Basic routes
	router.HandleFunc("/", homeHandler)

	// Route with regex constraint
	// TODO: Create a route for "/users/{id:[0-9]+}" (only numeric IDs) that calls getUserByIDHandler
	// This should match "/users/123" but not "/users/abc"
	// router.HandleFunc(???, ???)

	// Route with multiple variables and constraints
	// TODO: Create a route for "/posts/{year:[0-9]{4}}/{month:[0-9]{2}}/{title}" that calls getPostHandler
	// This matches posts with year (4 digits), month (2 digits), and any title
	// router.HandleFunc(???, ???)

	// Route with query parameter requirements
	// TODO: Create a route for "/search" with required query parameter "q" that calls searchHandler
	// Use .Queries() to specify required query parameters
	// router.HandleFunc(???, ???).Queries(???, ???)

	// Route with host matching
	// TODO: Create a route that only matches when host is "api.example.com" for "/health" that calls healthHandler
	// Use .Host() to specify the host constraint
	// router.HandleFunc(???, ???).Host(???)

	// Route with multiple HTTP methods
	// TODO: Create a route for "/data" that accepts both GET and POST methods and calls dataHandler
	// router.HandleFunc(???, ???).Methods(???, ???)

	fmt.Println("Server starting on :8080")
	fmt.Println("Try these URLs:")
	fmt.Println("- http://localhost:8080/users/123 (numeric ID)")
	fmt.Println("- http://localhost:8080/posts/2024/03/go-tutorial")
	fmt.Println("- http://localhost:8080/search?q=golang")
	fmt.Println("- http://api.example.com:8080/health (requires host header)")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Advanced Gorilla Mux Routing Examples\n")
	fmt.Fprintf(w, "Try the different route patterns listed in the console.\n")
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	
	// TODO: Convert the userID string to an integer using strconv.Atoi()
	// Handle the error by writing a 400 Bad Request response
	// id, err := strconv.Atoi(???)
	// if err != nil {
	//     http.Error(w, ???, http.StatusBadRequest)
	//     return
	// }
	
	// TODO: Write a response with the converted integer ID
	// Format: "User ID: %d (converted from string)"
	fmt.Fprintf(w, ???, ???)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	// TODO: Extract year, month, and title from the URL variables
	// year := vars[???]
	// month := vars[???]  
	// title := vars[???]
	
	// TODO: Write a response showing all extracted variables
	// Format: "Post: %s (Published: %s-%s)"
	fmt.Fprintf(w, ???, ???, ???, ???)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Get the required query parameter "q" from the request
	// query := r.URL.Query().Get(???)
	
	// TODO: Write a response with the search query
	// Format: "Searching for: %s"
	fmt.Fprintf(w, ???, ???)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write a JSON health response
	// w.Header().Set("Content-Type", "application/json")
	// fmt.Fprintf(w, ???)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Handle different HTTP methods differently
	// Use r.Method to check the HTTP method
	// For GET: return "Getting data"
	// For POST: return "Posting data" 
	// switch r.Method {
	// case ???:
	//     fmt.Fprintf(w, ???)
	// case ???:
	//     fmt.Fprintf(w, ???)
	// }
}