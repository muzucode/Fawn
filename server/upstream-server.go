package server

import (
	"os"

	"golang.org/x/crypto/ssh"
)

type UpstreamServer struct {
	Id          string
	DisplayName string
	Address     string
	Port        string
	SSH         SSH
}

type SSH struct {
	Client         *ssh.Client
	Session        *ssh.Session
	Stdin          *os.File
	Stdout         *os.File
	PrivateKeyPath string
}
