package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"muzucode/fawn/servers"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func ServerTest() (ServerList, error) {
	// SSH connection configuration
	port := 22
	privateKeyPath := "/Users/sean/.ssh/jungle_rsa"

	// Create servers
	s1 := &servers.Server{
		Id:      1,
		Name:    "Jungle",
		Address: os.Getenv("HOST_IP"),
		SSH: &servers.SSH{
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		},
	}

	// Read the private key file
	key := readPrivateKey(privateKeyPath)

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Create SSH config
	sshConfig := &ssh.ClientConfig{
		User: os.Getenv("HOST_USER"),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// Connect to the SSH server
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s1.Address, port), sshConfig)
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

	// Execute a command on the remote server
	cmd := "ls -la"
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Println("Command output:")
	fmt.Println(string(output))

	sl.Add(s1)

	return sl, err
}

func readPrivateKey(path string) []byte {
	// Read the private key file
	key, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	return key
}

func ListFiles() {

}
