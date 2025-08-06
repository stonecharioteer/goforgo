// advanced_routing.go - Solution
// Advanced gorilla/mux routing features: regex patterns, query parameters, and route matching

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
	router.HandleFunc("/users/{id:[0-9]+}", getUserByIDHandler)

	// Route with multiple variables and constraints
	router.HandleFunc("/posts/{year:[0-9]{4}}/{month:[0-9]{2}}/{title}", getPostHandler)

	// Route with query parameter requirements
	router.HandleFunc("/search", searchHandler).Queries("q", "")

	// Route with host matching
	router.HandleFunc("/health", healthHandler).Host("api.example.com")

	// Route with multiple HTTP methods
	router.HandleFunc("/data", dataHandler).Methods("GET", "POST")

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
	
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	
	fmt.Fprintf(w, "User ID: %d (converted from string)\n", id)
}

func getPostHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	month := vars["month"]
	title := vars["title"]
	
	fmt.Fprintf(w, "Post: %s (Published: %s-%s)\n", title, year, month)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Fprintf(w, "Searching for: %s\n", query)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"status": "healthy", "service": "api"}`)
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Getting data\n")
	case "POST":
		fmt.Fprintf(w, "Posting data\n")
	}
}