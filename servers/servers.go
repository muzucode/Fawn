package servers

import (
	"os"

	"golang.org/x/crypto/ssh"
)

type Server struct {
	Id             int
	Name           string
	Description    string
	Address        string
	PrivateKeyPath string
	SSH            *SSH
	GroupId        string
}

type SSH struct {
	Client  *ssh.Client
	Session *ssh.Session
	Stdin   *os.File
	Stdout  *os.File
	Stderr  *os.File
}
