package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"tcp-chat/server_func"
)

const (
	HOST = "localhost" // Host IP
	TYPE = "tcp"       // Connection type
)

// Main function gets the PORT from the arguments if there is any and calls the server loop
func main() {
	PORT := "8989"      // The default port
	args := os.Args[1:] // Get the arguments
	if len(args) > 1 {  // Should not be more than 1 argument
		fmt.Println("[USAGE]: go run . $port")
		return
	}
	if len(args) != 0 {
		PORT = args[0] // Set the port from the arguments
	}
	listener, err := net.Listen(TYPE, HOST+":"+PORT) // Listen to connections
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on the port:", PORT)
	defer listener.Close()           // Close listener
	server_func.ServerLoop(listener) // Call the server loop
}
