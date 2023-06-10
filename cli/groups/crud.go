package groups

import (
	"database/sql"
	"fmt"
	"log"
	db "muzucode/goroutines/database"
)

func CreateOne(group *Group) {

	// Insert the group into the database
	insertQuery := "INSERT INTO groups (Name, Description, EnvironmentId) VALUES (?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, group.Name, group.Description, group.EnvironmentId)
	if err != nil {
		log.Printf("Failed to add group: %v", err)
		return
	}

	fmt.Println("Group added successfully.")
}
func FindAll() (*sql.Rows, error) {

	// Retrieve all groups from the database
	fmt.Println("Getting all groups...")

	selectQuery := "SELECT * FROM Groups"
	rows, err := db.Db.Query(selectQuery)

	return rows, err
}
func FindAllByEnvironmentId(environmentId int) (*sql.Rows, error) {

	// Retrieve all groups from the database
	fmt.Println("Getting all groups...")

	selectQuery := "SELECT * FROM Groups WHERE EnvironmentId = ?"
	rows, err := db.Db.Query(selectQuery, environmentId)

	return rows, err
}
func DeleteOne(id int) {

	// Delete the group from the database
	deleteQuery := "DELETE FROM Groups WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Failed to delete group: %v", err)
		return
	}

	fmt.Println("Group deleted successfully.")
}
