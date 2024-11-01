package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)
// TCP Client Function
// Wait briefly to ensure the server is ready
// Connect to the server
// Read user input and send it to the server
// Exit the client if "exit" is entered
// Send the message to the server
// Read the response from the server

func startClient(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter message to send (or type 'exit' to quit): ")
		message, _ := reader.ReadString('\n')
		if message == "exit\n" {
			fmt.Println("Exiting client.")
			break
		}
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing to server: %v\n", err)
			break
		}
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from server: %v\n", err)
			break
		}
		fmt.Printf("Server response: %s", response)
	}
}