package server_func

import (
	"net"
	"os/exec"
)

/*
Function to clear the command line of the connection
Takes as the argument the connection
This is made to be both bash and cmd friendly
*/
func environmentClear(connection net.Conn) {
	cmd := exec.Command("clear") // CMD command
	cmd.Stdout = connection
	if err := cmd.Run(); err != nil { // If CMD command failed, use bash one
		cmd = exec.Command("cls") // Bash command
		cmd.Stdout = connection
		cmd.Run()
	}
}
