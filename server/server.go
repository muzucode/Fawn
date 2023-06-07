package server

import (
	"fmt"
	"net"
)

// Server status
const (
	Idle = iota
	Listening
)

type Server struct {
	Id          string
	DisplayName string
	Port        int
	Ip          net.IP
	Status      byte
}

func RunServer(address string) {
	fmt.Printf("Server listening @ %s\n", address)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	// Initialize read buf
	outBuf := make([]byte, 256)
	// Read data to buffer
	conn.Read(outBuf)

	defer conn.Close()

	fmt.Printf("%b\n", outBuf)
}
