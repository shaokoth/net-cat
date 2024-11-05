package main

import (
	"fmt"
	"net"
	"os"
)

func broadcast(message string) {
	for _, client := range clients {
		client.conn.Write([]byte(message + "\n"))
	}
}

func sendPreviousMessages(conn net.Conn) {
	for _, msg := range messageLog {
		conn.Write([]byte(msg + "\n"))
	}
}

func welcomeMessage() string {
	filepath := "net.text"
     data, err := os.ReadFile(filepath)
     if err != nil {
          fmt.Println("error reading file")
     }
     return string(data)
}
