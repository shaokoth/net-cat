# TCP Chat Application
# Description

* This project is a recreation of the NetCat (nc) system command in a server-client architecture, designed as a simple TCP-based group chat application. The application mimics the original NetCat functionality while incorporating additional features like user names, message timestamps, and multi-client chat.

* This project supports a server mode that listens for incoming connections and a client mode to connect to the server for communication.
Features:

    TCP Connections: Enables communication between a server and multiple clients (1-to-many relationship).
    Client Name Requirement: Each client must provide a unique, non-empty name upon connecting.
    Connection Control: Limits the number of simultaneous clients to a maximum of 10.
    Group Chat: All clients can send messages to the group, and everyone receives the messages.
    Message Formatting:
        Sent messages are displayed with a timestamp and sender's name in the format:
        [YYYY-MM-DD HH:MM:SS][client.name]:[client.message].
        Empty messages are not broadcasted.
    Message History: Newly connected clients receive all previous messages from the chat.
    Join/Leave Notifications:
        When a client joins, others are notified with a message:
        <ClientName> has joined our chat....
        When a client leaves, others are notified with a message:
        <ClientName> has left our chat....
    Persistent Connections: Other clients remain connected when one client disconnects.
    Default Port: If no port is specified, the server defaults to port 8989. Otherwise, it provides a usage message.
    Welcome Banner: Displays a Linux-style ASCII art welcome message upon connection.

## Project Requirements

    Language: Written in Go.
    Concurrency: Uses Go-routines for handling multiple client connections.
    Synchronization: Utilizes channels or mutexes for managing shared resources.
    Error Handling: Handles errors gracefully on both the server and client sides.
    Good Practices: Follows clean coding standards and best practices.
    Testing: Includes test files for unit testing server and client connections.

## Allowed Packages

    io
    log
    os
    fmt
    net
    sync
    time
    bufio
    errors
    strings
    reflect

# Usage
1. Starting the Server

* Run the server on the default port 8989:

$ go run . 
Listening on the port :8989

* Run the server on a custom port (e.g., 2525):

$ go run . 2525
Listening on the port :2525

2. Connecting a Client

Using the nc command, connect a client to the server:

$ nc localhost 2525
Welcome to TCP-Chat!
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
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: <YourName>

Example
Server Output

$ go run . 2525
Listening on the port :2525

Client 1

$ nc localhost 2525
Welcome to TCP-Chat!
[ENTER YOUR NAME]: Alice
[2024-11-19 14:00:01][Alice]: Hello everyone!
[2024-11-19 14:05:15][Alice]: How's it going?
Bob has joined our chat...
[2024-11-19 14:06:32][Bob]: Hi Alice!
[2024-11-19 14:07:00][Bob]: I'm good, thanks!
[2024-11-19 14:07:15][Alice]: Great to hear!
Bob has left our chat...

Client 2

$ nc localhost 2525
Welcome to TCP-Chat!
[ENTER YOUR NAME]: Bob
[2024-11-19 14:00:01][Alice]: Hello everyone!
[2024-11-19 14:05:15][Alice]: How's it going?
[2024-11-19 14:06:32][Bob]: Hi Alice!
[2024-11-19 14:07:00][Bob]: I'm good, thanks!
[2024-11-19 14:07:15][Alice]: Great to hear!

Default Behavior

    If no port is provided:

    [USAGE]: ./TCPChat $port

    Maximum connections exceeded:
        New clients see the message:
        Server is full. Try again later.

Testing

Run unit tests to verify functionality:

$ go test -v ./...

Implementation Details

    Server Logic:
        Accepts and manages multiple client connections.
        Synchronizes shared resources (e.g., message history, active connections).
    Client Logic:
        Reads and sends messages to the server.
        Receives broadcasted messages from the server.
    Concurrency:
        Go-routines handle simultaneous client interactions.
    Synchronization:
        Mutexes and channels prevent race conditions.

## Future Improvements

    Add support for UDP communication.
    Implement private messages between clients.
    Add commands for client interaction (e.g., /who to list active users).
    Enhance UI with terminal-based formatting.

## License

This project is open-source and available under the MIT License.
