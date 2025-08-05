// tcp_server.go - SOLUTION
// Learn TCP server implementation

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== TCP Server ===")
	
	// Listen on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	defer listener.Close()
	
	fmt.Println("Server listening on :8080")
	fmt.Println("Waiting for connections...")
	
	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %v\n", err)
			continue
		}
		
		// Handle connection in goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	
	// Get client address
	clientAddr := conn.RemoteAddr()
	fmt.Printf("New client connected: %s\n", clientAddr)
	
	// Set connection timeout
	conn.SetDeadline(time.Now().Add(5 * time.Minute))
	
	// Send welcome message
	welcome := "Welcome to Go TCP Server!\n"
	conn.Write([]byte(welcome))
	
	// Create reader for client messages
	reader := bufio.NewReader(conn)
	
	for {
		// Read message from client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected: %v\n", clientAddr, err)
			break
		}
		
		message = strings.TrimSpace(message)
		fmt.Printf("Received from %s: %s\n", clientAddr, message)
		
		// Handle special commands
		if strings.ToUpper(message) == "QUIT" {
			fmt.Printf("Client %s requested disconnect\n", clientAddr)
			break
		}
		
		if strings.ToUpper(message) == "TIME" {
			response := fmt.Sprintf("Server time: %s\n", time.Now().Format(time.RFC3339))
			conn.Write([]byte(response))
			continue
		}
		
		if strings.ToUpper(message) == "HELLO" {
			response := "Hello there! How can I help you?\n"
			conn.Write([]byte(response))
			continue
		}
		
		// Echo message back to client
		response := fmt.Sprintf("Echo: %s\n", message)
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Failed to write to client %s: %v\n", clientAddr, err)
			break
		}
		
		// Reset timeout for active connection
		conn.SetDeadline(time.Now().Add(5 * time.Minute))
	}
	
	fmt.Printf("Client %s connection closed\n", clientAddr)
}