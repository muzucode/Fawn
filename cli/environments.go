package cli

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func addGroup(cmd *cobra.Command, args []string) {
	// Prompt the user for group details
	var group Group
	fmt.Print("Enter Name: ")
	_, err := fmt.Scanln(&group.Name)
	if err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	fmt.Print("Enter Description: ")
	_, err = fmt.Scanln(&group.Description)
	if err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	// Generate a unique ID for the group
	group.ID = "1"

	// Insert the group into the database
	insertQuery := "INSERT INTO groups (Id, Name, Description) VALUES (?, ?, ?)"
	_, err = db.Exec(insertQuery, group.ID, group.Name, group.Description)
	if err != nil {
		log.Printf("Failed to add group: %v", err)
		return
	}

	fmt.Println("Group added successfully.")
}

func removeGroup(cmd *cobra.Command, args []string) {
	// Prompt the user for the group ID
	var groupID string
	fmt.Print("Enter Group ID: ")
	_, err := fmt.Scanln(&groupID)
	if err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	// Delete the group from the database
	deleteQuery := "DELETE FROM groups WHERE Id = ?"
	_, err = db.Exec(deleteQuery, groupID)
	if err != nil {
		log.Printf("Failed to remove group: %v", err)
		return
	}

	fmt.Println("Group removed successfully.")
}

func listGroups(cmd *cobra.Command, args []string) {

	// Connect to DB
	connectDatabase()

	// Retrieve all groups from the database
	fmt.Println("Getting all servers...")
	selectQuery := "SELECT * FROM groups"
	rows, err := db.Query(selectQuery)
	if err != nil {
		log.Printf("Failed to retrieve group information: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("All Groups:")
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.ID, &group.Name, &group.Description)
		if err != nil {
			log.Printf("Failed to retrieve group information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("ID: %s\n", group.ID)
		fmt.Printf("Name: %s\n", group.Name)
		fmt.Printf("Description: %s\n", group.Description)
		fmt.Println("---------------")
	}
}

func updateGroup(cmd *cobra.Command, args []string) {
	// Prompt the user for group details
	var group Group
	fmt.Print("Enter Group ID: ")
	_, err := fmt.Scanln(&group.ID)
	if err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	fmt.Print("Enter New Name: ")
	_, err = fmt.Scanln(&group.Name)
	if err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	fmt.Print("Enter New Description: ")
	_, err = fmt.Scanln(&group.Description)
	if err != nil {
		log.Printf("Invalid input: %v", err)
		return
	}

	// Update the group in the database
	updateQuery := "UPDATE groups SET Name = ?, Description = ? WHERE Id = ?"
	_, err = db.Exec(updateQuery, group.Name, group.Description, group.ID)
	if err != nil {
		log.Printf("Failed to update group: %v", err)
		return
	}

	fmt.Println("Group updated successfully.")
}
