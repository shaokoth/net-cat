package main

import (
	"net"
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
	return `Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    .       |' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     -'       --'
`
}
