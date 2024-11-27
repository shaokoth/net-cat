# TCP-Chat

## About Project
This project consists on recreating the `NetCat` in a `Server-Client` Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.

NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open `TCP` connections, send UDP packages, listen on arbitrary `TCP` and UDP ports and many more.

To see more information about NetCat inspect the manual `man nc`.
## Usage
```
$ go run . $port
Listening on the port :$port
```
### Example
### Connecting to a Client
- `1st` Terminal
```bash
$ nc localhost $yourPort or default port 8989
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
[ENTER YOUR NAME]: Shadrack
[2021-10-27 19:14:10][Shadrack]:Hello? 
[2021-10-27 19:14:20][Shadrack]:!
Cliff has joined our chat...
[2021-10-27 19:14:32][Shadrack]:!
[2021-10-27 19:14:43][Cliff]:Hi
[2021-10-27 19:14:43][Shadrack]:Wow, How Are you?
[2021-10-27 19:15:04][Shadrack]:!
[2021-10-27 19:15:31][Cliff]:I am good, thank you!
[2021-10-27 19:15:31][Shadrack]:!
[2021-10-27 19:15:51][Cliff]:Lets play Lem-in
[2021-10-27 19:15:51][Shadrack]:HA-ha Go!
[2021-10-27 19:16:03][Shadrack]:!
Cliff has left our chat...
[2021-10-27 19:16:10][Shadrack]:^C
$ 
```
- `2nd` Terminal
```bash
$ nc localhost 8989
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
[ENTER YOUR NAME]: Hezborn
[2021-10-27 19:14:20][Shadrack]:Hello?
[2021-10-27 19:14:32][Hezborn]:Hi
[2021-10-27 19:14:43][Hezborn]:!
[2021-10-27 19:15:04][Shadrack]:Wow, How Are you?
[2021-10-27 19:15:04][Hezborn]:I am good, thank you!
[2021-10-27 19:15:31][Hezborn]:Lets play Lem-in
[2021-10-27 19:15:51][Hezborn]:!
[2021-10-27 19:16:03][Shadrack]:HA-ha Go!
[2021-10-27 19:16:03][Hezborn]:^C
$ 
```
### Default settings
Default settings on main.go
```go
MaxConnections = 10 // Max connections count
Port = ":8989"      // Default port if user does not set
```
## Build
```bash
#Build project
$ go build -o TCPChat
#Usage
$ ./TCPChat $port
```
## Future Improvements
Add support for UDP communication.
Implement private messages between clients.
Add commands for client interaction (e.g., /who to list active users).
Enhance UI with terminal-based formatting.
Create more than 1 group chat.

## Authors
- [Shadrack](https://github.com/shaokoth)
- [Cliff](https://github.com/cliffdoyle)
- [Hezborn](https://github.com/Mania124)