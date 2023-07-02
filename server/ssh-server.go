package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"muzucode/fawn/servers"
	"time"

	"golang.org/x/crypto/ssh"
)

func ConnectUsingSSH(s servers.Server) (*ssh.Session, error) {
	// SSH connection configuration
	port := 22

	// Read the private key file
	key, err := ioutil.ReadFile(s.PrivateKeyPath)
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
		User: "sean",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.AddressIPv4, port), config)
	if err != nil {
		log.Fatalf("Failed to connect to SSH server: %v", err)
	}

	// Create a new SSH session
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create SSH session: %v", err)
	}

	return session, err

}

func GetFilesInDir(s servers.Server, dirPath string) ([]string, error) {
	session, err := ConnectUsingSSH(s)
	if err != nil {
		log.Fatal(err)
	}
	// defer session.Close()

	// Execute list files command
	cmd := `"ls"`
	output, err := session.CombinedOutput(cmd)

	// Handle errors
	if err != nil {
		fmt.Printf("Failed to execute command: %v\n", err)
		return nil, err
	}

	// Print the command output
	fmt.Println("Command output:")
	fmt.Println(string(output))

	// Split by new line
	var files [][]byte
	var filesStrArr []string
	files = bytes.Split(output, []byte("\n"))

	// Add to array of strings
	for i := 0; i < len(files)-1; i++ {
		// Log each file
		fmt.Printf("Fetched file %d:\n", i)
		fmt.Println(string(files[i]))
		filesStrArr = append(filesStrArr, string(files[i])) // Append to string array
	}

	return filesStrArr, nil
}
