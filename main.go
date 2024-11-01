package main

import (
	"fmt"
	"os"
)

const defaultPort = "8989"

func main() {
	port := defaultPort
	if len(os.Args) > 2 || (len(os.Args) == 2 && os.Args[1] == "localhost") {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		port = os.Args[1]
	}
	fmt.Printf("Listening on the port :%s\n", port)
	startServer(port)
}
