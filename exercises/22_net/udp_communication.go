// udp_communication.go
// Learn UDP networking for connectionless communication

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
	fmt.Println("=== UDP Networking Demo ===")
	fmt.Println("Choose mode:")
	fmt.Println("1. Start UDP server")
	fmt.Println("2. Start UDP client")
	fmt.Print("Enter choice (1 or 2): ")
	
	// TODO: Read user choice
	reader := bufio.NewReader(os.Stdin)
	choice, _ := /* read line from reader */
	choice = strings.TrimSpace(choice)
	
	switch choice {
	case "1":
		startUDPServer()
	case "2":
		startUDPClient()
	default:
		fmt.Println("Invalid choice. Starting server by default.")
		startUDPServer()
	}
}

func startUDPServer() {
	fmt.Println("\n=== Starting UDP Server ===")
	
	// TODO: Resolve UDP address
	addr, err := /* resolve UDP address "localhost:8081" */
	if err != nil {
		fmt.Printf("Error resolving address: %v\n", err)
		return
	}
	
	// TODO: Listen on UDP port
	conn, err := /* listen on UDP address */
	if err != nil {
		fmt.Printf("Error starting UDP server: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("UDP Server listening on localhost:8081")
	fmt.Println("Waiting for UDP packets...")
	
	// TODO: Buffer for receiving messages
	buffer := make([]byte, 1024)
	
	for {
		// TODO: Read UDP packet
		n, clientAddr, err := /* read from UDP connection */
		if err != nil {
			fmt.Printf("Error reading UDP packet: %v\n", err)
			continue
		}
		
		// TODO: Process received message
		message := string(buffer[:n])
		fmt.Printf("Received from %s: %s\n", clientAddr, message)
		
		// TODO: Send response back to client
		timestamp := time.Now().Format("15:04:05")
		response := fmt.Sprintf("[%s] UDP Echo: %s", timestamp, message)
		
		/* write response to client address */
		if err != nil {
			fmt.Printf("Error sending response: %v\n", err)
		}
	}
}

func startUDPClient() {
	fmt.Println("\n=== Starting UDP Client ===")
	
	// TODO: Resolve server address
	serverAddr, err := /* resolve UDP address "localhost:8081" */
	if err != nil {
		fmt.Printf("Error resolving server address: %v\n", err)
		return
	}
	
	// TODO: Create UDP connection
	conn, err := /* dial UDP to server */
	if err != nil {
		fmt.Printf("Error creating UDP connection: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to UDP server!")
	fmt.Println("Type messages to send (type 'quit' to exit):")
	
	// TODO: Create input scanner
	scanner := bufio.NewScanner(os.Stdin)
	
	for /* scan user input */ {
		input := /* get input text */
		
		// TODO: Check for quit command
		if /* input is "quit" */ {
			break
		}
		
		// TODO: Send message to server
		_, err := /* write input to connection */
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			continue
		}
		
		// TODO: Read response from server
		buffer := make([]byte, 1024)
		/* set read deadline */
		
		n, err := /* read from connection */
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}
		
		response := string(buffer[:n])
		fmt.Printf("Server: %s\n", response)
	}
	
	fmt.Println("UDP Client disconnected")
}

func demonstrateUDPBroadcast() {
	fmt.Println("\n=== UDP Broadcast Demo ===")
	
	// TODO: Create UDP connection for broadcasting
	conn, err := /* dial UDP to broadcast address */
	if err != nil {
		fmt.Printf("Error creating broadcast connection: %v\n", err)
		return
	}
	defer conn.Close()
	
	// TODO: Enable broadcast
	/* set broadcast socket option */
	
	// TODO: Send broadcast message
	message := "Hello from UDP broadcast!"
	/* write message to connection */
	
	fmt.Println("Broadcast message sent!")
}