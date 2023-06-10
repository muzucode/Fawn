package server

import (
	"os"

	"golang.org/x/crypto/ssh"
)

type Server struct {
	Id            string
	Name          string
	Address       string
	SSH           *SSH
	EnvironmentId string
}

type SSH struct {
	Client         *ssh.Client
	Session        *ssh.Session
	Stdin          *os.File
	Stdout         *os.File
	Stderr         *os.File
	PrivateKeyPath string
}
