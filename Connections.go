package main

import (
	"bufio"
	"fmt"
	"net"
)

// Handle Client Connections on Server Side
// Create a buffered reader for reading client messages
// Read message from client
// Send response back to client
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Client connected: %s\n", conn.RemoteAddr().String())
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from client: %v\n", err)
			break
		}
		fmt.Printf("Received from client: %s", message)
		response := "Message received: " + message
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Printf("Error writing to client: %v\n", err)
			break
		}
	}
	fmt.Printf("Client disconnected: %s\n", conn.RemoteAddr().String())
}
