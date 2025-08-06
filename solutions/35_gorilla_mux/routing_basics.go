// routing_basics.go - Solution
// Learn HTTP routing with gorilla/mux - the powerful HTTP router for Go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux router
	router := mux.NewRouter()

	// Basic route handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/about", aboutHandler)

	// Method-specific routes
	router.HandleFunc("/users", createUserHandler).Methods("POST")

	// Route with URL variables
	router.HandleFunc("/users/{id}", getUserHandler).Methods("GET")

	// Start the server
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!\n")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Us Page\n")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User created successfully!\n")
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintf(w, "Getting user with ID: %s\n", userID)
}