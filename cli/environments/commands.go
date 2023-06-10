package environments

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Load all commands
func LoadEnvironmentCommands(rootCmd *cobra.Command) {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "Manage environments",
	}
	rootCmd.AddCommand(envCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an environment",
		Run:   AddEnvironmentCommand,
	}
	envCmd.AddCommand(addCmd)

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an environment",
		Run:   DeleteEnvironmentCommand,
	}
	envCmd.AddCommand(deleteCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all environments",
		Run:   ListEnvironmentsCommand,
	}
	envCmd.AddCommand(listCmd)

	defaultCmd := &cobra.Command{
		Use:   "default",
		Short: "Print the default environment",
		Run:   GetDefaultEnvironmentCommand,
	}
	envCmd.AddCommand(defaultCmd)

}

// Commands
func AddEnvironmentCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for environment details
	var environment Environment
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Environment Name: ")
	if scanner.Scan() {
		environment.Name = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	fmt.Print("Enter API Key: ")
	if scanner.Scan() {
		environment.ApiKey = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	fmt.Print("Enter Debug Mode (true/false): ")
	if scanner.Scan() {
		debugMode := scanner.Text()
		if debugMode == "true" || debugMode == "1" {
			environment.DebugMode = true
		} else {
			environment.DebugMode = false
		}
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Insert the environment into the database

	CreateOne(environment)
}

func DeleteEnvironmentCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for the environment Id
	var environmentId int

	fmt.Print("Enter Environment Id: ")
	_, err := fmt.Scanf("%d", &environmentId)
	if err != nil {
		log.Fatal(err)
	}

	// Delete the environment from the database
	DeleteOne(environmentId)
}
func ListEnvironmentsCommand(cmd *cobra.Command, args []string) {
	rows, err := FindAll()
	if err != nil {
		fmt.Println("Failed to retrieve environment information:")
		log.Fatal(err)

	}
	defer rows.Close()

	fmt.Println("All Environments:")
	fmt.Println("---------------")
	for rows.Next() {
		var environment Environment
		err := rows.Scan(&environment.Id, &environment.Name, &environment.ApiKey, &environment.DebugMode, &environment.Position, &environment.CreatedAt, &environment.UpdatedAt)
		if err != nil {
			log.Printf("Failed to retrieve environment information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %d\n", environment.Id)
		fmt.Printf("Environment Name: %s\n", environment.Name)
		fmt.Printf("API Key: %s\n", environment.ApiKey)
		fmt.Printf("Position: %d\n", environment.Position)
		fmt.Printf("Debug Mode: %v\n", environment.DebugMode)
		fmt.Printf("Created At: %v\n", string(environment.CreatedAt))
		fmt.Printf("Updated At: %v\n", string(environment.UpdatedAt))
		fmt.Println("---------------")
	}
}
func GetDefaultEnvironmentCommand(cmd *cobra.Command, args []string) {
	environment, err := FindDefault()
	if err != nil {
		fmt.Println("Failed at GetDefaultEnvironmentCommand")
		log.Fatal(err)
	}

	fmt.Printf("Id: %d\n", environment.Id)
	fmt.Printf("Environment Name: %s\n", environment.Name)
	fmt.Printf("API Key: %s\n", environment.ApiKey)
	fmt.Printf("Position: %d\n", environment.Position)
	fmt.Printf("Debug Mode: %v\n", environment.DebugMode)
	fmt.Printf("Created At: %v\n", string(environment.CreatedAt))
	fmt.Printf("Updated At: %v\n", string(environment.UpdatedAt))
	fmt.Println("---------------")
}
