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
	
	// Read user choice
	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
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
	
	// Listen on UDP port 8080
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Printf("Error resolving UDP address: %v\n", err)
		return
	}
	
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Printf("Error listening on UDP: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("UDP server listening on :8080")
	fmt.Println("Send messages from clients. Type 'quit' to stop.")
	
	buffer := make([]byte, 1024)
	
	for {
		// Read from UDP connection
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error reading UDP message: %v\n", err)
			continue
		}
		
		message := string(buffer[:n])
		fmt.Printf("Received from %s: %s\n", clientAddr, message)
		
		// Echo back to client
		response := fmt.Sprintf("Server received: %s", message)
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Printf("Error sending response: %v\n", err)
		}
		
		// Check for quit command
		if strings.TrimSpace(message) == "quit" {
			fmt.Println("Server shutting down...")
			break
		}
	}
}

func startUDPClient() {
	fmt.Println("\n=== Starting UDP Client ===")
	
	// Connect to UDP server
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Printf("Error resolving server address: %v\n", err)
		return
	}
	
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()
	
	fmt.Println("Connected to UDP server at localhost:8080")
	fmt.Println("Type messages to send. Type 'quit' to exit.")
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		// Read user input
		fmt.Print("Client> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			break
		}
		
		message = strings.TrimSpace(message)
		if message == "quit" {
			// Send quit message to server
			conn.Write([]byte(message))
			break
		}
		
		// Send message to server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error sending message: %v\n", err)
			continue
		}
		
		// Read response from server
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("Error reading response: %v\n", err)
			continue
		}
		
		response := string(buffer[:n])
		fmt.Printf("Server: %s\n", response)
	}
	
	fmt.Println("Client disconnected.")
}

func demonstrateBroadcast() {
	fmt.Println("\n=== UDP Broadcast Demo ===")
	
	// Create UDP connection for broadcasting
	conn, err := net.Dial("udp", "255.255.255.255:9999")
	if err != nil {
		fmt.Printf("Error creating broadcast connection: %v\n", err)
		return
	}
	defer conn.Close()
	
	// Send broadcast message
	message := "Hello, everyone! This is a broadcast message."
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error sending broadcast: %v\n", err)
		return
	}
	
	fmt.Println("Broadcast message sent!")
}

func demonstrateMulticast() {
	fmt.Println("\n=== UDP Multicast Demo ===")
	
	// Multicast address (224.0.0.0 to 239.255.255.255)
	multicastAddr := "224.1.1.1:9999"
	
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		fmt.Printf("Error resolving multicast address: %v\n", err)
		return
	}
	
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("Error creating multicast connection: %v\n", err)
		return
	}
	defer conn.Close()
	
	// Send multicast message
	message := "Hello, multicast group!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Printf("Error sending multicast: %v\n", err)
		return
	}
	
	fmt.Println("Multicast message sent!")
}