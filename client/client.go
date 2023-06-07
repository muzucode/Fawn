package client

import (
	"log"
	"net"
)

func RunClient(address string) {

	// Connect to server
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	data := [32]byte{}

	// Write data to wire
	conn.Write(data[:])

}
