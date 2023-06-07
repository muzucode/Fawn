package server

import (
	"fmt"
	"log"
	"net"
)

// Server status
const (
	DoNotDisturb = iota
	Listening
)

type Server struct {
	Id          string
	DisplayName string
	Addr        string
	Port        string
	Status      byte
	Connection  *net.Conn
}

func (s Server) Run() {
	fmt.Printf("Server listening @ %s\n", s.Addr)
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		s.Connection = &conn
		if err != nil {
			log.Fatal(err)
		}
		go s.HandleConnection()
	}
}

func (s Server) Stop() {

}

func (s Server) HandleConnection() {
	conn := *s.Connection

	// DoNotDisturb - Ignore
	if s.Status == DoNotDisturb {
		fmt.Println("Sorry, server isn't listening to your BS!")
		return
	}
	// Listening - Read to buffer
	if s.Status == Listening {
		buf := make([]byte, 128)
		_, err := conn.Read(buf)

		if err != nil {
			log.Fatal(err)
		}

	}

}
