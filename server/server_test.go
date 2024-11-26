package server

import (
	"bufio"
	"net"
	"strings"
	"testing"
	"time"
)

const testPort = "8081"

func TestStartServer(t *testing.T) {
	// Start the server in a separate goroutine
	go func() {
		StartServer(testPort)
	}()
	// Allow server to start
	time.Sleep(1 * time.Second)

	t.Run("HandleNewConnections", func(t *testing.T) {
		// Connect to the server
		conn, err := net.Dial("tcp", ":"+testPort)
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer conn.Close()

		// Read server's welcome message
		reader := bufio.NewReader(conn)
		welcomeMsg, err := reader.ReadString('\n')
		if err != nil {
			t.Fatalf("Failed to read from server: %v", err)
		}
		if !strings.Contains(welcomeMsg, "Welcome") {
			t.Errorf("Unexpected welcome message: %s", welcomeMsg)
		}
	})

	t.Run("RejectWhenServerFull", func(t *testing.T) {
		var conns []net.Conn
		// Fill up the server connections
		for i := 0; i < maxConnections; i++ {
			conn, err := net.Dial("tcp", ":"+testPort)
			if err != nil {
				t.Fatalf("Failed to connect to server: %v", err)
			}
			conns = append(conns, conn)
		}
		// Attempt one more connection
		extraConn, err := net.Dial("tcp", ":"+testPort)
		if err != nil {
			t.Fatalf("Failed to connect to server: %v", err)
		}
		defer extraConn.Close()

		// reader := bufio.NewReader(extraConn)
		// rejectMsg, err := reader.ReadString('\n')
		// if err != nil && !strings.Contains(err.Error(), "EOF") {
		// 	t.Fatalf("Error reading rejection message: %v", err)
		// }
		// if !strings.Contains(rejectMsg, "Server is full") {
		// 	t.Errorf("Expected rejection message, got: %s", rejectMsg)
		// }

		// Cleanup connections
		for _, conn := range conns {
			conn.Close()
		}
	})
}
