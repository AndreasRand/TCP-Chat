# TCP-Chat

## Description

TCP-Chat is a Go program that creates a local chat server to send messages between connected clients.
This project was made to learn more about Go, specifically more about goroutines and TCP connections. The program itself is just the server part of it.

## How to run

To run the program you need to have at least Go version 1.19 installed. You can also lower the go version from the mod file if you know what you are doing.

Go can be installed from [here](https://go.dev/doc/install).

The program can be run by using one of the commands
```
go run main.go
go run .
```

You can add an argument to the command to specify the port for the server. The default port is 8989.

Or you can build the executable file with the command
``` 
go build
```

## How to connect to the chat server

To connect to the server you can use netcat (Ncat for windows) or any other similar program.

This is a netcat example of how to connect to the server.

Open up another terminal and insert this command.
```
nc localhost 8989
```

If you want to connect another client repeat the process.

## Dependencies

This project has no dependencies.

## Author

Andreas Randmäe