package server_func

import "net"

/*
Server and client history are a bit different from each other. Server history is the history
that will be loaded to every client upon connection and that will not contain messages about
users joining or leaving, only the actual message history itself. Client history contains
also messages about the users who left or joined during your session.
*/

// Client structure that contains the connection to the client and the client history
type Client struct {
	Connection    net.Conn
	ClientHistory string
}

// Server structure that contains the server history
type Server struct {
	ServerHistory string
}
