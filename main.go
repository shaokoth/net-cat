package main

import (
	"fmt"
	"os"
	"sync"

)
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
