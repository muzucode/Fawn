package groups

import (
	"bufio"
	"fmt"
	"log"
	"muzucode/fawn/environments"
	"os"

	"github.com/spf13/cobra"
)

func LoadGroupCommands(rootCmd *cobra.Command) {
	groupCmd := &cobra.Command{
		Use:   "group",
		Short: "Manage groups",
	}
	rootCmd.AddCommand(groupCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a group",
		Run:   AddOneCommand,
	}
	groupCmd.AddCommand(addCmd)

	removeCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a group",
		Run:   DeleteOneCommand,
	}
	groupCmd.AddCommand(removeCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all groups in environment",
		Run:   ListAllInCurrentEnvironmentCommand,
	}
	groupCmd.AddCommand(listCmd)

	listAllCmd := &cobra.Command{
		Use:   "all",
		Short: "List all groups across all environments",
		Run:   ListAllCommand,
	}
	listCmd.AddCommand(listAllCmd)
}

func AddOneCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for group details
	var group Group
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Name: ")
	if scanner.Scan() {
		group.Name = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	fmt.Print("Enter Description: ")
	if scanner.Scan() {
		group.Description = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	env, err := environments.FindCurrent()
	if err != nil {
		fmt.Println("Failed to get current environment.")
		log.Fatal(err)
	}

	group.EnvironmentId = env.Id

	CreateOne(&group)
}
func DeleteOneCommand(cmd *cobra.Command, args []string) {
	// Prompt the user for the environment Id
	var groupId int

	fmt.Print("Enter Group Id: ")
	_, err := fmt.Scanf("%d", &groupId)
	if err != nil {
		log.Fatal(err)
	}

	DeleteOne(groupId)
}
func ListAllCommand(cmd *cobra.Command, args []string) {

	rows, err := FindAll() // 'Scan' crud command
	if err != nil {
		log.Printf("Failed to retrieve group information: %v", err)
	}

	defer rows.Close()

	fmt.Println("All Groups:")
	fmt.Println("---------------")
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Id, &group.Name, &group.Description, &group.EnvironmentId)
		if err != nil {
			log.Printf("Failed to retrieve group information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %d\n", group.Id)
		fmt.Printf("Name: %s\n", group.Name)
		fmt.Printf("Description: %s\n", group.Description)
		fmt.Println("---------------")
	}
}
func ListAllInCurrentEnvironmentCommand(cmd *cobra.Command, args []string) {

	currentEnvironmentId, err1 := environments.FindCurrentId()
	if err1 != nil {
		log.Printf("Failed to retrieve group information: %v", err1)
	}

	rows, err2 := FindAllByEnvironmentId(currentEnvironmentId)
	if err2 != nil {
		log.Printf("Failed to retrieve group information: %v", err2)
	}

	defer rows.Close()

	fmt.Println("All Groups:")
	fmt.Println("---------------")
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Id, &group.Name, &group.Description, &group.EnvironmentId)
		if err != nil {
			log.Printf("Failed to retrieve group information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %d\n", group.Id)
		fmt.Printf("Name: %s\n", group.Name)
		fmt.Printf("Description: %s\n", group.Description)
		fmt.Println("---------------")
	}
}
