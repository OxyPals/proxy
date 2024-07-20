package main

import (
	"fmt"
	"net"
	"proxy/policy/error"
)

func main() {
	/*
		Server is listening for inbound connection
		When connection is accepted we starts a procedure waits for Connect command
		when we
	*/
	fmt.Println("Starting...")
	if listener, err := net.Listen("tcp", ":1080"); err == nil {
		if connection, err := listener.Accept(); err == nil {
			request := make([]byte, 1024)
			if n, _ := connection.Read(request); n > 0 {

			}
		} else {
			error.HandleAndExit(err)
		}
	} else {
		error.HandleAndExit(err)
	}
}
