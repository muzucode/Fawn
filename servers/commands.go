package servers

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Load all commands
func LoadServerCommands(rootCmd *cobra.Command) {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Manage servers",
	}
	rootCmd.AddCommand(serverCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a server",
		Run:   AddServerCommand,
	}
	serverCmd.AddCommand(addCmd)

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a server",
		Run:   DeleteServerCommand,
	}
	serverCmd.AddCommand(deleteCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all servers",
		Run:   ListServersCommand,
	}
	serverCmd.AddCommand(listCmd)

}

// Commands
func AddServerCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for server details
	var server *Server
	scanner := bufio.NewScanner(os.Stdin)

	// GroupId
	fmt.Print("Enter Id of group to add server to: ")
	if scanner.Scan() {
		server.GroupId = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Name
	fmt.Print("Enter Server Name: ")
	if scanner.Scan() {
		server.Name = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Description
	fmt.Print("Enter Server Description: ")
	if scanner.Scan() {
		server.Description = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}
	// Address
	fmt.Print("Enter Server Description: ")
	if scanner.Scan() {
		server.Description = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}
	// PrivateKeyPath
	fmt.Print("Enter path to private key: ")
	if scanner.Scan() {
		server.Description = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Insert the server into the database

	CreateOne(server)
}
func DeleteServerCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for the server Id
	var serverId int

	fmt.Print("Enter Server Id: ")
	_, err := fmt.Scanf("%d", &serverId)
	if err != nil {
		log.Fatal(err)
	}

	// Delete the server from the database
	DeleteOne(serverId)
}
func ListServersCommand(cmd *cobra.Command, args []string) {
	rows, err := FindAll()
	if err != nil {
		fmt.Println("Failed to retrieve server information:")
		log.Fatal(err)

	}
	defer rows.Close()

	fmt.Println("All Servers:")
	fmt.Println("---------------")
	for rows.Next() {
		var server Server
		err := rows.Scan(
			&server.Id,
			&server.Name,
			&server.AddressIPv4,
			&server.AddressIPv6,
			&server.PrivateKeyPath,
			&server.GroupId,
			&server.Description,
			&server.DistributionName,
			&server.DistributionVersion,
		)
		if err != nil {
			log.Printf("Failed to retrieve server information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %d\n", server.Id)
		fmt.Printf("Name: %s\n", server.Name)
		fmt.Printf("Description: %s\n", server.Description)
		fmt.Printf("IPv4 Address: %s\n", server.AddressIPv4)
		fmt.Printf("IPv6 Address: %s\n", server.AddressIPv6)
		fmt.Printf("Private Key Path: %v\n", server.PrivateKeyPath)
		fmt.Printf("Group Id: %v\n", string(server.GroupId))
		fmt.Println("---------------")
	}
}
