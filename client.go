package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

var clients = make(map[net.Conn]*Client)

// Client represents a chat client
type Client struct {
	conn net.Conn
	name string
}

// Connects to the Server using TCP
// Ask for client name and handle empty name case
// Announce new client and add to clients map
// Handles client disconnection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte(welcomeMessage()))
	conn.Write([]byte("[ENTER YOUR NAME]:"))
	nameReader := bufio.NewReader(conn)
	name, err := nameReader.ReadString('\n')
	if err != nil || strings.TrimSpace(name) == "" {
		conn.Write([]byte("Invalid name. Disconnecting...\n"))
		return
	}
	name = strings.TrimSpace(name)

	mutex.Lock()
	client := &Client{conn: conn, name: name}
	clients[conn] = client
	messageLog = append(messageLog, fmt.Sprintf("%s has joined our chat...", name))
	broadcast(fmt.Sprintf("%s has joined our chat...", name))
	sendPreviousMessages(conn)
	mutex.Unlock()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		fullMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)

		mutex.Lock()
		messageLog = append(messageLog, fullMessage)
		broadcast(fullMessage)
		mutex.Unlock()
	}
	mutex.Lock()
	delete(clients, conn)
	messageLog = append(messageLog, fmt.Sprintf("%s has left our chat...", name))
	broadcast(fmt.Sprintf("%s has left our chat...", name))
	connections--
	mutex.Unlock()
}
