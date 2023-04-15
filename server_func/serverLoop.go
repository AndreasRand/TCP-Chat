package server_func

import (
	"log"
	"net"
	"os"
	"sync"
)

const (
	// The intro message that is prompted when connecting to the server
	introMessage = ("Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    `.       | `' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     `-'       `--'\n" +
		"[ENTER YOUR NAME]:")
	// Error messages
	writeIntroError   = "Error writing intro message"
	writeHistoryError = "Could not write history to "
)

var ServerError = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile) // Server error color format
var server = Server{ServerHistory: ""}                                                           // Create the empty server history
var connectionMap = &sync.Map{}                                                                  // Map for storing the TCP connections

// Function to create a loop that acceps connections to the server
func ServerLoop(listener net.Listener) {
	for {
		connection, err := listener.Accept()
		if err != nil { // Accepting the connection failed
			ServerError.Printf("Could not accept connection from %v\n", connection)
			continue
		}
		go handleConnection(connection) // Handle the connection as a thread process, so connections do not get delayed
	}
}

/*
Function to handle the connection of a new client
Takes as argument the connection
*/
func handleConnection(conn net.Conn) {
	environmentClear(conn)                                      // Clear the command line of the client
	writeToConnection(conn, writeIntroError, introMessage)      // Write the intro message to the connected client
	username, startHistory := readUsername(conn, connectionMap) // Get the username from the client
	if username == "" {                                         // Did not get the username, close the connection
		return
	}
	startHistory = introMessage + startHistory // Add the username prompt history and intro message
	newClient := Client{                       // Create a new Client and add the server history to it
		Connection:    conn,
		ClientHistory: startHistory + "\n\n-Welcome to the chat " + username + "-\n\n" + server.ServerHistory,
	}
	toWrite := username + " has joined the chat...\n"
	writeToAllConnections(connectionMap, toWrite, username)                                                         // Write to all connections that a new user joined
	connectionMap.Store(username, newClient)                                                                        // Store the new user to the map
	environmentClear(conn)                                                                                          // Clear the command line of the client
	writeToConnection(conn, writeHistoryError+username, newClient.ClientHistory+"\n[\033[34m"+username+"\033[0m]:") // Write the chat history to the client
	go handleIncomingRequests(conn, connectionMap, username)                                                        // Handle new requests (messages) from the client as a thread process
}
