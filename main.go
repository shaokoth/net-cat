package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

// TCP Server Function
func startServer(wg *sync.WaitGroup) {
	defer wg.Done()

	// Listen for incoming connections
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Error starting TCP server: %v\n", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on 127.0.0.1:8080")

	for {
		// Accept an incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

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

// TCP Client Function
func startClient(wg *sync.WaitGroup) {
	defer wg.Done()

	// Wait briefly to ensure the server is ready
	time.Sleep(1 * time.Second)

	// Connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")

	// Read user input and send it to the server
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')

		// Exit the client if "exit" is entered
		if message == "exit\n" {
			fmt.Println("Exiting client.")
			break
		}

		// Send the message to the server
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing to server: %v\n", err)
			break
		}

		// Read the response from the server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from server: %v\n", err)
			break
		}
		fmt.Printf("Server response: %s", response)
	}
}

// Main Function
func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	var wg sync.WaitGroup

	// Start the server
	wg.Add(1)
	go startServer(&wg)

	// Start the client
	wg.Add(1)
	go startClient(&wg)

	// Wait for both server and client to finish
	wg.Wait()
}
