package server_func

import (
	"bufio"
	"net"
	"strings"
	"sync"
)

// Error outputs
const (
	readUsernameError  = "Error reading username input from the connection"
	writeUsernameError = "Error writing new username request to the connection"
)

/*
Function to get the username input, contains a loop that will broken when a correct input is given
There are two checks for the username, if it's not empty and if it's already taken
Takes as argument the current connection that has just connected and the connection map
Returns the username and chat history of the input prompt
*/
func readUsername(connection net.Conn, connectionMap *sync.Map) (string, string) {
	username, history := "", ""
	var err error
	for {
		if username, err = bufio.NewReader(connection).ReadString('\n'); err != nil { // Read input from the user
			errorWhileReading(connection, readUsernameError, err) // Error while reading the input
			return "", ""
		}
		history += username // Add the input to the chat history
		username = strings.Trim(username, "\n")
		if username == "" { // Given input is empty
			newMessage := "Username can't be empty, try again" +
				"\n[ENTER YOUR NAME]:" // Message to output to the connection
			writeToConnection(connection, writeUsernameError, newMessage) // Write the message to the connection
			history += newMessage                                         // Add the message to the chat history
			continue                                                      // Ask for a new input
		}
		if _, ok := connectionMap.Load(username); ok { // Check if the username is already taken (connected currently)
			users := ""
			connectionMap.Range(func(key, value interface{}) bool { // Add all users to the string
				users += ", " + key.(string)
				return true
			})
			newMessage := "Username is already taken, try another one\n" +
				"Active users:" +
				users[1:] +
				"\n[ENTER YOUR NAME]:" // Message to output to the connection that includes currently connected usernames
			writeToConnection(connection, writeUsernameError, newMessage) // Write the message to the connection
			history += newMessage                                         // Add the message to the chat history
			continue                                                      // Ask for a new input
		}
		break // All checks passed, so the loop is broken
	}
	return username, history
}
