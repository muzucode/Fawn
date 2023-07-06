package servers

import (
	"database/sql"
	"fmt"
	"log"
	db "muzucode/fawn/database"
)

func InsertOne(server *Server) error {
	fmt.Printf("%+v\n", server)
	// Insert the group into the database
	insertQuery := "INSERT INTO Servers (Name, AddressIPv4, PrivateKeyPath, GroupId, Description, DistributionName, DistributionVersion) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, server.Name, server.AddressIPv4, server.PrivateKeyPath, server.GroupId, server.Description, server.DistributionName, server.DistributionVersion)
	if err != nil {
		log.Printf("Failed to add server: %v", err)
		return err
	}

	fmt.Println("Server added successfully.")
	return nil
}

func FindOne(serverId int) (*Server, error) {
	var s Server
	selectQuery := "SELECT * FROM Servers WHERE Id = ?"
	row := db.Db.QueryRow(selectQuery, serverId)

	err := row.Scan(
		&s.Id,
		&s.Name,
		&s.AddressIPv4,
		&s.PrivateKeyPath,
		&s.GroupId,
		&s.Description,
		&s.DistributionName,
		&s.DistributionVersion,
		&s.AddressIPv6,
	)

	return &s, err
}
func FindAll() (*sql.Rows, error) {

	// Retrieve all servers from the database
	fmt.Println("Getting all servers...")

	selectQuery := "SELECT * FROM Servers"
	rows, err := db.Db.Query(selectQuery)

	return rows, err
}
func FindAllByGroupId(groupId int) (*sql.Rows, error) {

	// Retrieve all servers from the database
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
