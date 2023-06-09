package groups

import (
	"bufio"
	"fmt"
	"log"
	db "muzucode/goroutines/database"
	"os"

	"github.com/spf13/cobra"
)

type Group struct {
	Id            string
	Name          string
	Description   string
	EnvironmentId string
}

func AddGroup(cmd *cobra.Command, args []string) {
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

	group.EnvironmentId = "3"

	// Insert the group into the database
	insertQuery := "INSERT INTO groups (Name, Description, EnvironmentId) VALUES (?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, group.Name, group.Description, group.EnvironmentId)
	if err != nil {
		log.Printf("Failed to add group: %v", err)
		return
	}

	fmt.Println("Group added successfully.")
}

func RemoveGroup(cmd *cobra.Command, args []string) {
	// Prompt the user for the group Id
	var groupId string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Group Id: ")
	if scanner.Scan() {
		groupId = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Delete the group from the database
	deleteQuery := "DELETE FROM groups WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, groupId)
	if err != nil {
		log.Printf("Failed to remove group: %v", err)
		return
	}

	fmt.Println("Group removed successfully.")
}

func ListGroups(cmd *cobra.Command, args []string) {
	// Retrieve all groups from the database
	fmt.Println("Getting all servers...")
	selectQuery := "SELECT * FROM groups"
	rows, err := db.Db.Query(selectQuery)
	if err != nil {
		log.Printf("Failed to retrieve group information: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("All Groups:")
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Id, &group.Name, &group.Description)
		if err != nil {
			log.Printf("Failed to retrieve group information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %s\n", group.Id)
		fmt.Printf("Name: %s\n", group.Name)
		fmt.Printf("Description: %s\n", group.Description)
		fmt.Println("---------------")
	}
}
