package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func ConnectUsingSSH() {
	// SSH connection configuration
	host := "your_host"
	port := 22
	user := "your_username"
	privateKeyPath := "path_to_private_key"

	// Read the private key file
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	// Parse the private key
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// SSH client configuration
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// Connect to the SSH server
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		log.Fatalf("Failed to connect to SSH server: %v", err)
	}
	defer conn.Close()

	// Create a new SSH session
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Set up IO streams for session
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// Execute a command on the remote server
	cmd := "ls -al"
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Println("Command output:")
	fmt.Println(string(output))
}
