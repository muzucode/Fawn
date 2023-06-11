package servers

import (
	"database/sql"
	"fmt"
	"log"
	db "muzucode/goroutines/database"
)

func CreateOne(server *Server) {

	// Insert the group into the database
	insertQuery := "INSERT INTO Servers (Name, Description, Address, PrivateKeyPath, GroupId) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, server.Name, server.Description, server.Address, server.PrivateKeyPath, server.GroupId)
	if err != nil {
		log.Printf("Failed to add group: %v", err)
		return
	}

	fmt.Println("Group added successfully.")
}

func FindOne(serverId int) (*Server, error) {
	var server Server
	selectQuery := "SELECT * FROM Servers WHERE Id = ?"
	row := db.Db.QueryRow(selectQuery, serverId)

	err := row.Scan(&server.Id, &server.Name, &server.Description, &server.Address, &server.PrivateKeyPath, &server.GroupId)

	return &server, err
}
func FindAll() (*sql.Rows, error) {

	// Retrieve all groups from the database
	fmt.Println("Getting all groups...")

	selectQuery := "SELECT * FROM Groups"
	rows, err := db.Db.Query(selectQuery)

	return rows, err
}
func FindAllByGroupId(groupId int) (*sql.Rows, error) {

	// Retrieve all groups from the database
	fmt.Println("Getting all servers in group...")

	selectQuery := "SELECT * FROM Servers WHERE GroupId = ?"
	rows, err := db.Db.Query(selectQuery, groupId)

	return rows, err
}

func DeleteOne(serverId int) {

	// Delete the group from the database
	deleteQuery := "DELETE FROM Servers WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, serverId)
	if err != nil {
		log.Printf("Failed to delete server: %v", err)
		return
	}

	fmt.Println("Group deleted successfully.")
}
