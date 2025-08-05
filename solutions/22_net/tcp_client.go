// tcp_client.go - SOLUTION
// Learn TCP client implementation

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
	fmt.Println("=== TCP Client ===")
	
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to server localhost:8080")
	
	// Set timeout for operations
	conn.SetDeadline(time.Now().Add(30 * time.Second))
	
	// Send initial greeting
	message := "Hello from Go TCP client!\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		return
	}
	
	// Read response
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read response: %v\n", err)
		return
	}
	
	fmt.Printf("Server response: %s", response)
	
	// Interactive chat loop
	fmt.Println("\nEnter messages (type 'quit' to exit):")
	
	scanner := bufio.NewScanner(os.Stdin)
	reader := bufio.NewReader(conn)
	
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		
		input := scanner.Text()
		if strings.ToLower(input) == "quit" {
			break
		}
		
		// Send message to server
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Printf("Failed to send: %v\n", err)
			break
		}
		
		// Read server response
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read: %v\n", err)
			break
		}
		
		fmt.Printf("Server: %s", response)
	}
	
	// Send goodbye message
	conn.Write([]byte("QUIT\n"))
	fmt.Println("Connection closed")
}