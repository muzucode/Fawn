package servers

import (
	"os"

	"golang.org/x/crypto/ssh"
)

type Server struct {
	Id                  int
	Name                string
	Address             string
	Description         string
	DistributionName    string
	DistributionVersion string
	PrivateKeyPath      string
	SSH                 *SSH // not stored in DB
	GroupId             string
}

type SSH struct {
	Client  *ssh.Client
	Session *ssh.Session
	Stdin   *os.File
	Stdout  *os.File
	Stderr  *os.File
}
