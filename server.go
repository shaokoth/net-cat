package main

import (
	"log"
	"net"
	"sync"
	
)
const maxConnections=10

var (
	clients =make(map[net.Conn]*Client)
	mutex sync.Mutex
	messageLog []string
	connections int
)
//Handles setting up and 
//Starting the TCP server
//And managing the connections
func startServer(port string){
listener,err:=net.Listen("tcp",":"+port)
if err !=nil{
	log.Fatalf("Error starting server: %v", err)
}
defer listener.Close()

for{
	conn,err:=listener.Accept()
	if err !=nil{
		log.Printf("Error accepting connection: %v",err)
		continue
	}
	mutex.Lock()
	if connections>=maxConnections{
		conn.Write([]byte("Server is full. Try again later. \n"))
		conn.Close()
		mutex.Unlock()
		continue
	}
	connections++
	mutex.Unlock()
	go handleConnection(conn)
}
}