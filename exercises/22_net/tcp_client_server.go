// tcp_client_server.go
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
	
	// TODO: Read user choice
	reader := bufio.NewReader(os.Stdin)
	choice, _ := /* read line from reader */
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
	
	// TODO: Listen on TCP port 8080
	listener, err := /* listen on "localhost:8080" */
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	defer listener.Close()
	
	fmt.Println("Server listening on localhost:8080")
	fmt.Println("Waiting for connections...")
	
	for {
		// TODO: Accept incoming connections
		conn, err := /* accept connection */
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}
		
		// TODO: Handle each connection concurrently
		go /* handle connection */
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// TODO: Get client address
	clientAddr := /* get remote address */
	fmt.Printf("New client connected: %s\n", clientAddr)
	
	// TODO: Create scanner to read from connection
	scanner := /* create scanner for connection */
	
	// TODO: Send welcome message
	welcomeMsg := "Welcome to Go TCP Server! Type 'quit' to disconnect.\n"
	/* write welcome message to connection */
	
	for /* scan for messages */ {
		message := /* get scanned text */
		fmt.Printf("Received from %s: %s\n", clientAddr, message)
		
		// TODO: Check for quit command
		if /* check if message is "quit" */ {
			/* send goodbye message */
			fmt.Printf("Client %s disconnected\n", clientAddr)
			break
		}
		
		// TODO: Echo message back with timestamp
		timestamp := time.Now().Format("15:04:05")
		response := fmt.Sprintf("[%s] Echo: %s\n", timestamp, message)
		/* write response to connection */
	}
	
	// TODO: Check for scanner errors
	if err := /* get scanner error */; err != nil {
		fmt.Printf("Connection error: %v\n", err)
	}
}

func startClient() {
	fmt.Println("\n=== Starting TCP Client ===")
	
	// TODO: Connect to server
	conn, err := /* dial "localhost:8080" */
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to server!")
	
	// TODO: Start goroutine to read server messages
	go func() {
		scanner := /* create scanner for connection */
		for /* scan server messages */ {
			message := /* get scanned text */
			fmt.Printf("Server: %s\n", message)
		}
	}()
	
	// TODO: Read user input and send to server
	fmt.Println("Type messages to send (type 'quit' to exit):")
	inputScanner := bufio.NewScanner(os.Stdin)
	
	for /* scan user input */ {
		input := /* get input text */
		
		// TODO: Send message to server
		/* write input + newline to connection */
		
		// TODO: Check for quit command
		if /* input is "quit" */ {
			break
		}
	}
	
	fmt.Println("Client disconnected")
}