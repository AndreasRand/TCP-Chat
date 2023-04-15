package server_func

import (
	"net"
	"sync"
)

const messageSendError = "Failed to send message to "

/*
Function to write a message to the given connection
Takes as argument the connection to write to, the error message that will be output in case of an error and the message to write to the connection
*/
func writeToConnection(connection net.Conn, errorMessage string, toWrite string) {
	if _, err := connection.Write([]byte(toWrite)); err != nil {
		ServerError.Println(errorMessage)
	}
}

/*
Function to write a message to every single connected client
Takes as argument the map of the connections, message to write and the username of the message author
*/
func writeToAllConnections(connMap *sync.Map, toWrite string, sentFrom string) {
	connMap.Range(func(key, value interface{}) bool { // Iterate through every client
		client, _ := value.(Client)
		username := key.(string)
		environmentClear(client.Connection) // Clear the command line of the client
		if username != sentFrom {           // If the client does not match the author
			client.ClientHistory += toWrite // Add message to client history
			connMap.Store(key, client)      // Store the updated value
		}
		writeToConnection(client.Connection, messageSendError+username, client.ClientHistory+"\n[\033[34m"+username+"\033[0m]:") // Write the new history to the client
		return true
	})
}
