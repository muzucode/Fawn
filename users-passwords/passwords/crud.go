package passwords

import (
	"database/sql"
	"fmt"
	"log"
	db "muzucode/fawn/database"
)

func CreateOne(password *Password) {

	// TODO: Save password to DB, get its ID and save ID to password row

	// Insert the password into the database
	insertQuery := "INSERT INTO Passwords (Salt, HashingAlgorithm, Hash) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, password.Salt, password.HashingAlgorithm, password.Hash)
	if err != nil {
		log.Printf("Failed to add password: %v", err)
		return
	}

	fmt.Println("Password added successfully.")
}

func FindOne(passwordId int) (*Password, error) {
	var password Password
	selectQuery := "SELECT ?,?,?,? FROM Passwords WHERE Id = ?"
	row := db.Db.QueryRow(selectQuery, passwordId)

	err := row.Scan(&password.Id, &password.Salt, &password.HashingAlgorithm, &password.Hash)

	return &password, err
}
func FindAll() (*sql.Rows, error) {

	// Retrieve all passwords from the database
	fmt.Println("Getting all passwords...")

	selectQuery := "SELECT * FROM Passwords"
	rows, err := db.Db.Query(selectQuery)

	return rows, err
}

func DeleteOne(passwordId int) {

	// Delete the password from the database
	deleteQuery := "DELETE FROM Passwords WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, passwordId)
	if err != nil {
		log.Printf("Failed to delete password: %v", err)
		return
	}

	fmt.Println("Password deleted successfully.")
}
