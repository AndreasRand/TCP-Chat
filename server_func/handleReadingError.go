package server_func

import (
	"io"
	"net"
)

/*
Function to output errors that occur while reading input from the user
The only error that is ignored is the EOF error which occurs when the user uses CTRL+C to stop the program
*/
func errorWhileReading(connection net.Conn, errorMessage string, err error) {
	defer connection.Close()
	if err != io.EOF { // Error is something other than program being stopped while reading input
		ServerError.Println(errorMessage)
	}
}
