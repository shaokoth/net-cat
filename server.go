package main

import (
	"fmt"
	"net"
	"sync"
	
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