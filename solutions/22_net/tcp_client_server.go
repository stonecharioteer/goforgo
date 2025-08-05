// tcp_client_server.go - SOLUTION
// Learn TCP networking with client and server implementations

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== TCP Networking Demo ===")
	fmt.Println("Choose mode:")
	fmt.Println("1. Start server")
	fmt.Println("2. Start client")
	fmt.Print("Enter choice (1 or 2): ")
	
	// Read user choice
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	
	switch choice {
	case "1":
		startServer()
	case "2":
		startClient()
	default:
		fmt.Println("Invalid choice. Starting server by default.")
		startServer()
	}
}

func startServer() {
	fmt.Println("\n=== Starting TCP Server ===")
	
	// Listen on TCP port 8080
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()
	
	fmt.Println("Server listening on localhost:8080")
	fmt.Println("Waiting for connections...")
	
	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		
		// Handle each connection concurrently
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// Get client address
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("New client connected: %s\n", clientAddr)
	
	// Create scanner to read from connection
	scanner := bufio.NewScanner(conn)
	
	// Send welcome message
	welcomeMsg := "Welcome to Go TCP Server! Type 'quit' to disconnect.\n"
	conn.Write([]byte(welcomeMsg))
	
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("Received from %s: %s\n", clientAddr, message)
		
		// Check for quit command
		if strings.ToLower(strings.TrimSpace(message)) == "quit" {
			conn.Write([]byte("Goodbye!\n"))
			fmt.Printf("Client %s disconnected\n", clientAddr)
			break
		}
		
		// Echo message back with timestamp
		timestamp := time.Now().Format("15:04:05")
		response := fmt.Sprintf("[%s] Echo: %s\n", timestamp, message)
		conn.Write([]byte(response))
	}
	
	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Connection error: %v\n", err)
	}
}

func startClient() {
	fmt.Println("\n=== Starting TCP Client ===")
	
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to server!")
	
	// Start goroutine to read server messages
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			message := scanner.Text()
			fmt.Printf("Server: %s\n", message)
		}
	}()
	
	// Read user input and send to server
	fmt.Println("Type messages to send (type 'quit' to exit):")
	inputScanner := bufio.NewScanner(os.Stdin)
	
	for inputScanner.Scan() {
		input := inputScanner.Text()
		
		// Send message to server
		conn.Write([]byte(input + "\n"))
		
		// Check for quit command
		if strings.ToLower(strings.TrimSpace(input)) == "quit" {
			break
		}
	}
	
	fmt.Println("Client disconnected")
}