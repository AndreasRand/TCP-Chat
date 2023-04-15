package server_func

import (
	"bufio"
	"net"
	"sync"
	"time"
)

/*
Function to handle messages from the client after it has connected to the server
Takes as argument the connection to the client, the map of all connections and the username of the client
*/
func handleIncomingRequests(connection net.Conn, connectionMap *sync.Map, username string) {
	defer func() { // This will be run at the end of the function
		connection.Close()             // Close connection
		connectionMap.Delete(username) // Delete the connection from the map
		toWrite := username + " has left the chat...\n"
		writeToAllConnections(connectionMap, toWrite, username) // Write the leave message to everyone that's connected
	}()
	for { // Loop to read for input
		userInput := ""
		var err error
		if userInput, err = bufio.NewReader(connection).ReadString('\n'); err != nil { // Read input
			errorWhileReading(connection, "error getting input from user "+username, err) // Error while reading the input
			return
		}
		sentTime := time.Now().Format("02-01-2006 15:04:05")           // Get the time of the message
		toWrite := "[" + sentTime + "][" + username + "]:" + userInput // Format the message
		server.ServerHistory += toWrite                                // Add message to history
		result, _ := connectionMap.Load(username)
		client := result.(Client)
		client.ClientHistory += "[" + sentTime + "][\033[34m" + username + "\033[0m]:" + userInput // Update the client history with custom color username
		connectionMap.Store(username, client)                                                      // Update the stored client value
		writeToAllConnections(connectionMap, toWrite, username)                                    // Write the new message to every client
	}
}
