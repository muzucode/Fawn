package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Load all commands
func LoadServerCommands(rootCmd *cobra.Command) {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "Manage environments",
	}
	rootCmd.AddCommand(envCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an server",
		Run:   AddServerCommand,
	}
	envCmd.AddCommand(addCmd)

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an server",
		Run:   DeleteServerCommand,
	}
	envCmd.AddCommand(deleteCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all environments",
		Run:   ListServersCommand,
	}
	envCmd.AddCommand(listCmd)

	defaultCmd := &cobra.Command{
		Use:   "default",
		Short: "Print the default server",
		Run:   GetDefaultServerCommand,
	}
	envCmd.AddCommand(defaultCmd)

}

// Commands
func AddServerCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for server details
	var server Server
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Server Name: ")
	if scanner.Scan() {
		server.Name = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	fmt.Print("Enter API Key: ")
	if scanner.Scan() {
		server.ApiKey = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	fmt.Print("Enter Debug Mode (true/false): ")
	if scanner.Scan() {
		debugMode := scanner.Text()
		if debugMode == "true" || debugMode == "1" {
			server.DebugMode = true
		} else {
			server.DebugMode = false
		}
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Insert the server into the database

	CreateOne(server)
}

func DeleteServerCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for the server Id
	var environmentId int

	fmt.Print("Enter Server Id: ")
	_, err := fmt.Scanf("%d", &environmentId)
	if err != nil {
		log.Fatal(err)
	}

	// Delete the server from the database
	DeleteOne(environmentId)
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
		err := rows.Scan(&server.Id, &server.Name, &server.ApiKey, &server.DebugMode, &server.Position, &server.CreatedAt, &server.UpdatedAt)
		if err != nil {
			log.Printf("Failed to retrieve server information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %d\n", server.Id)
		fmt.Printf("Server Name: %s\n", server.Name)
		fmt.Printf("API Key: %s\n", server.ApiKey)
		fmt.Printf("Position: %d\n", server.Position)
		fmt.Printf("Debug Mode: %v\n", server.DebugMode)
		fmt.Printf("Created At: %v\n", string(server.CreatedAt))
		fmt.Printf("Updated At: %v\n", string(server.UpdatedAt))
		fmt.Println("---------------")
	}
}
func GetDefaultServerCommand(cmd *cobra.Command, args []string) {
	server, err := FindDefault()
	if err != nil {
		fmt.Println("Failed at GetDefaultServerCommand")
		log.Fatal(err)
	}

	fmt.Printf("Id: %d\n", server.Id)
	fmt.Printf("Server Name: %s\n", server.Name)
	fmt.Printf("API Key: %s\n", server.ApiKey)
	fmt.Printf("Position: %d\n", server.Position)
	fmt.Printf("Debug Mode: %v\n", server.DebugMode)
	fmt.Printf("Created At: %v\n", string(server.CreatedAt))
	fmt.Printf("Updated At: %v\n", string(server.UpdatedAt))
	fmt.Println("---------------")
}
