// tcp_client.go
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
	
	// TODO: Connect to TCP server
	conn, err := /* connect to localhost:8080 using tcp */
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to server localhost:8080")
	
	// TODO: Set timeout for operations
	/* set connection deadline to 30 seconds from now */
	
	// TODO: Send initial greeting
	greeting := "Hello from TCP client!"
	/* write greeting to connection */
	
	// TODO: Read server response
	response := make([]byte, 1024)
	n, err := /* read from connection into response buffer */
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}
	
	fmt.Printf("Server response: %s\n", string(response[:n]))
	
	// TODO: Interactive communication loop
	fmt.Println("Enter messages to send to server (type 'quit' to exit):")
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Client> ")
		
		// TODO: Read user input
		message, err := /* read line from stdin reader */
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}
		
		message = strings.TrimSpace(message)
		if message == "quit" {
			break
		}
		
		// TODO: Send message to server
		_, err = /* write message to connection */
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			break
		}
		
		// TODO: Read server response
		/* set read deadline to 5 seconds from now */
		n, err := /* read server response */
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}
		
		fmt.Printf("Server: %s\n", string(response[:n]))
	}
	
	fmt.Println("Client disconnected.")
}

func demonstrateAdvancedClient() {
	fmt.Println("\n=== Advanced TCP Client Features ===")
	
	// TODO: Connect with custom dialer
	dialer := net.Dialer{
		Timeout: 10 * time.Second,
	}
	
	conn, err := /* use dialer to connect to localhost:8080 */
	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer conn.Close()
	
	// TODO: Send structured data
	data := map[string]interface{}{
		"type":      "client_info",
		"timestamp": time.Now().Unix(),
		"message":   "Advanced client connection",
	}
	
	/* marshal data to JSON and send to server */
	
	// TODO: Handle connection errors gracefully
	/* implement retry logic and error handling */
}