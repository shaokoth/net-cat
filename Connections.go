package main

import (
	"bufio"
	"fmt"
	"net"
)

// Handle Client Connections on Server Side
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Client connected: %s\n", conn.RemoteAddr().String())

	// Create a buffered reader for reading client messages
	reader := bufio.NewReader(conn)
	for {
		// Read message from client
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from client: %v\n", err)
			break
		}
		fmt.Printf("Received from client: %s", message)

		// Send response back to client
		response := "Message received: " + message
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Error writing to client: %v\n", err)
			break
		}
	}
	fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr().String())
}
