// routing_basics.go
// Learn HTTP routing with gorilla/mux - the powerful HTTP router for Go
//
// Gorilla Mux is a request router and dispatcher that matches incoming requests
// to their respective handlers. It's known for its powerful URL matching capabilities
// including variables, regular expressions, and methods.
//
// I AM NOT DONE YET!

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux router
	// TODO: Create a new mux router using mux.NewRouter()
	// router := ???

	// Basic route handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/about", aboutHandler)

	// Method-specific routes
	// TODO: Add a POST route for "/users" that calls createUserHandler
	// router.HandleFunc(???, ???).Methods(???)

	// Route with URL variables
	// TODO: Add a GET route for "/users/{id}" that calls getUserHandler
	// The {id} part is a URL parameter that can be extracted in the handler
	// router.HandleFunc(???, ???).Methods(???)

	// Start the server
	fmt.Println("Server starting on :8080")
	// TODO: Use http.ListenAndServe to start the server on port 8080 with the router
	// log.Fatal(http.ListenAndServe(???, ???))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!\n")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us Page\n")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Write a response indicating user creation
	// Example: "User created successfully!"
	fmt.Fprintf(w, ???)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Extract the 'id' parameter from the URL using mux.Vars(r)
	// vars := mux.Vars(r)
	// userID := vars[???]
	
	// TODO: Write a response with the user ID
	// Example format: "Getting user with ID: %s"
	fmt.Fprintf(w, ???, ???)
}