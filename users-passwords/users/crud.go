package users

import (
	"database/sql"
	"fmt"
	"log"
	db "muzucode/fawn/database"
)

func CreateOne(user *User) {

	// TODO: Save password to DB, get its ID and save ID to user row

	// Insert the user into the database
	insertQuery := "INSERT INTO Users (Name, ThemeId, PasswordId) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, user.Name, user.Theme.Id, user.Password.Id)
	if err != nil {
		log.Printf("Failed to add user: %v", err)
		return
	}

	fmt.Println("User added successfully.")
}

func FindOne(userId int) (*User, error) {
	var user User
	selectQuery := "SELECT ?,?,?,? FROM Users WHERE Id = ?"
	row := db.Db.QueryRow(selectQuery, userId)

	err := row.Scan(&user.Id, &user.Name, &user.Theme.Id, &user.Password.Id)

	return &user, err
}
func FindAll() (*sql.Rows, error) {

	// Retrieve all users from the database
	fmt.Println("Getting all users...")

	selectQuery := "SELECT * FROM Users"
	rows, err := db.Db.Query(selectQuery)

	return rows, err
}

func DeleteOne(userId int) {

	// Delete the user from the database
	deleteQuery := "DELETE FROM Users WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, userId)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return
	}

	fmt.Println("User deleted successfully.")
}
