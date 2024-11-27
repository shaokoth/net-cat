package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Client struct {
	conn net.Conn
	name string
}
const maxConnections = 10

var (
	mutex       sync.Mutex
	messageLog  []string
	connections int
	clients = make(map[net.Conn]*Client)
)

/*Handles setting up and
 Starting the TCP server
 And managing the connections*/
func StartServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		mutex.Lock()
		if connections >= maxConnections {
			conn.Write([]byte("Server is full. Try again later. \n"))
			conn.Close()
			mutex.Unlock()
			continue
		}
		mutex.Unlock()

		go handleConnection(conn)
	}
}

// Connects to the Server using TCP 
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// print welcome message and ask for client name and handle empty name case
	conn.Write([]byte(welcomeMessage()))
	conn.Write([]byte("\n[ENTER YOUR NAME]:"))
	nameReader := bufio.NewReader(conn)
	name, err := nameReader.ReadString('\n')
	if err != nil || strings.TrimSpace(name) == "" {
		conn.Write([]byte("Invalid name. Disconnecting...\n"))
		return
	}
	name = strings.TrimSpace(name)

	mutex.Lock()
	// Announce new client and add to clients map
	client := &Client{conn: conn, name: name}
	clients[conn] = client
	connections++
	mutex.Unlock()
	
	broadcast(fmt.Sprintf("%s has joined our chat...", name), conn)

	sendPreviousMessages(conn)

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
		clearScreen := "\033[F\033[K"
		fullMessage := fmt.Sprintf("%v[%s][%s]: %s", clearScreen, timestamp, name, message)
		otherClientsMessage := fmt.Sprintf("[%s][%s]: %s", timestamp, name, message)

		mutex.Lock()
		messageLog = append(messageLog, otherClientsMessage)
		mutex.Unlock()
		// Send the cleared message to the sender only
		conn.Write([]byte(fullMessage + "\n"))

		// Broadcast the non-cleared message to other clients
		for c, client := range clients {
			if c != conn {
				client.conn.Write([]byte(otherClientsMessage + "\n"))
			}
		}
	}
	// Handle client disconnection
	mutex.Lock()
	delete(clients, conn)
	//messageLog = append(messageLog, fmt.Sprintf("%s has left our chat...", name))
	broadcast(fmt.Sprintf("%s has left our chat...", name), conn)
	connections--
	mutex.Unlock()
}

func broadcast(message string, sender net.Conn) {
	for c, client := range clients {
		if c != sender {
			client.conn.Write([]byte(message + "\n"))
		}
	}
}

func sendPreviousMessages(conn net.Conn) {
	for _, msg := range messageLog {
		conn.Write([]byte(msg + "\n"))
	}
}

func welcomeMessage() string {
	filepath := "./net.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("error reading file: %v", err)
		return "Welcome to the Chat server"
	}
	return string(data)
}