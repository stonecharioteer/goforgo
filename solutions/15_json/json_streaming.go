// json_streaming.go - SOLUTION
// Learn JSON streaming for processing large datasets efficiently

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

func main() {
	fmt.Println("=== JSON Streaming ===")
	
	// Create JSON encoder for streaming output
	var output strings.Builder
	encoder := json.NewEncoder(&output)
	
	// Stream multiple JSON objects
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com", Active: true},
		{ID: 2, Name: "Bob", Email: "bob@example.com", Active: false},
		{ID: 3, Name: "Charlie", Email: "charlie@example.com", Active: true},
	}
	
	fmt.Println("Streaming JSON objects:")
	for _, user := range users {
		// Encode each user individually
		err := encoder.Encode(user)
		if err != nil {
			fmt.Printf("Encoding error: %v\n", err)
		}
	}
	
	fmt.Printf("Streamed output:\n%s\n", output.String())
	
	fmt.Println("\n=== JSON Decoder Streaming ===")
	
	// Create input with multiple JSON objects
	jsonInput := `{"id":1,"name":"David","email":"david@test.com","active":true}
{"id":2,"name":"Eve","email":"eve@test.com","active":false}
{"id":3,"name":"Frank","email":"frank@test.com","active":true}`
	
	// Create decoder for streaming input
	decoder := json.NewDecoder(strings.NewReader(jsonInput))
	
	fmt.Println("Decoding streamed JSON:")
	for decoder.More() {
		var user User
		
		// Decode next JSON object
		err := decoder.Decode(&user)
		if err != nil {
			fmt.Printf("Decoding error: %v\n", err)
			break
		}
		
		fmt.Printf("Decoded user: %+v\n", user)
	}
	
	fmt.Println("\n=== Token-Level Streaming ===")
	
	// Parse JSON at token level
	complexJSON := `{
		"users": [
			{"name": "Alice", "age": 30},
			{"name": "Bob", "age": 25}
		],
		"total": 2
	}`
	
	decoder = json.NewDecoder(strings.NewReader(complexJSON))
	
	fmt.Println("Token-level parsing:")
	for decoder.More() {
		// Get next token
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Token error: %v\n", err)
			break
		}
		
		fmt.Printf("Token: %v (type: %T)\n", token, token)
	}
	
	fmt.Println("\n=== JSON Lines Format ===")
	
	// Handle JSON Lines format (JSONL)
	jsonLines := `{"event": "login", "user": "alice", "timestamp": "2023-01-01T10:00:00Z"}
{"event": "purchase", "user": "bob", "timestamp": "2023-01-01T10:15:00Z"}
{"event": "logout", "user": "alice", "timestamp": "2023-01-01T10:30:00Z"}`
	
	type Event struct {
		Event     string `json:"event"`
		User      string `json:"user"`
		Timestamp string `json:"timestamp"`
	}
	
	// Process each line as separate JSON
	lines := strings.Split(jsonLines, "\n")
	
	fmt.Println("Processing JSON Lines:")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		
		var event Event
		// Unmarshal each line
		err := json.Unmarshal([]byte(line), &event)
		if err != nil {
			fmt.Printf("Line %d error: %v\n", i+1, err)
			continue
		}
		
		fmt.Printf("Event %d: %s by %s at %s\n", i+1, event.Event, event.User, event.Timestamp)
	}
	
	fmt.Println("\n=== Custom JSON Marshaling ===")
	
	// Implement custom marshaling for User
	type CustomUser struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	
	// Implement MarshalJSON method
	func (u CustomUser) MarshalJSON() ([]byte, error) {
		// Create custom JSON format
		return json.Marshal(map[string]interface{}{
			"user_id": u.ID,
			"full_name": strings.ToUpper(u.Name),
			"email_address": u.Email,
			"created_at": "2023-01-01T00:00:00Z",
		})
	}
	
	customUser := CustomUser{ID: 1, Name: "Custom User", Email: "custom@example.com"}
	
	// Marshal with custom method
	customJSON, err := json.Marshal(customUser)
	if err != nil {
		fmt.Printf("Custom marshal error: %v\n", err)
	} else {
		fmt.Printf("Custom JSON: %s\n", customJSON)
	}
}