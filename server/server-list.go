package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

type ServerList map[string]*Server

var sl = make(ServerList)

func (sl ServerList) GetInstance() *ServerList {
	return &sl
}
func (sl ServerList) FindOne(id string) *Server {
	return sl[id]
}
func (sl ServerList) Add(s *Server) {
	sl[s.Id] = s
}
func (sl ServerList) Delete(s *Server) {

	delete(sl, (*s).Id)
}
func FetchUpstreamServers() (ServerList, error) {
	// SSH connection configuration
	port := 22
	privateKeyPath := "/Users/sean/.ssh/jungle_rsa"

	// Create servers
	s1 := &Server{
		Id:          "1",
		DisplayName: "Server-1",
		Address:     "127.0.0.1",
	}
	// s2 := &Server{
	// 	Id:          "1",
	// 	DisplayName: "Server-1",
	// 	Address:     "127.0.0.1",
	// 	Port:        "9823",
	// }
	// s3 := &Server{
	// 	Id:          "1",
	// 	DisplayName: "Server-1",
	// 	Address:     "127.0.0.1",
	// 	Port:        "9823",
	// }

	// Read the private key file
	key := readPrivateKey(privateKeyPath)

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Create SSH config
	sshConfig := &ssh.ClientConfig{
		User: "ubuntu	",
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

	// Set Stdout and Stderr
	session.Stdout = s1.SSH.Stdout
	session.Stderr = s1.SSH.Stderr

	// Execute a command on the remote server
	cmd := "ls -al"
	output, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}

	fmt.Println("Command output:")
	fmt.Println(string(output))

	sl.Add(s1)

	return sl, err
}
func LoadServerList() {
	var fetchedServers, err = FetchUpstreamServers()
	if err != nil {
		log.Fatal(err)
	}

	sl = fetchedServers
}

func readPrivateKey(path string) []byte {
	// Read the private key file
	key, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read private key file: %v", err)
	}

	return key
}
