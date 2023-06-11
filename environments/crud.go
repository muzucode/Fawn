package environments

import (
	"database/sql"
	"fmt"
	"log"
	db "muzucode/fawn/database"
)

// Environments
func CreateOne(environment Environment) {

	// Insert the environment into the database
	insertQuery := "INSERT INTO Environments (Name, ApiKey, DebugMode) VALUES (?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, environment.Name, environment.ApiKey, environment.DebugMode)
	if err != nil {
		log.Printf("Failed to add environment: %v", err)
		return
	}

	fmt.Println("Environment added successfully.")
}
func DeleteOne(id int) {

	// Delete the environment from the database
	deleteQuery := "DELETE FROM Environments WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, id)
	if err != nil {
		log.Printf("Failed to remove environment: %v", err)
		return
	}

	fmt.Println("Environment removed successfully.")
}
func FindAll() (*sql.Rows, error) {
	// Retrieve all environments from the database
	fmt.Println("Getting all environments...")
	selectQuery := "SELECT * FROM Environments"
	rows, err := db.Db.Query(selectQuery)

	return rows, err
}
func FindOne(environmentId int) (*Environment, error) {
	var environment Environment
	selectQuery := "SELECT * FROM Environments WHERE Id = ?"
	row := db.Db.QueryRow(selectQuery, environmentId)

	err := row.Scan(&environment.Id, &environment.Name, &environment.ApiKey, &environment.DebugMode, &environment.Position, &environment.CreatedAt, &environment.UpdatedAt)

	return &environment, err
}

// CurrentEnvironment
func UpdateCurrent(environmentId int) error {

	selectQuery := "UPDATE CurrentEnvironment SET EnvironmentId = ?"
	_, err := db.Db.Exec(selectQuery, environmentId)

	return err
}
func FindCurrent() (*Environment, error) {

	var selectQuery string
	var row *sql.Row
	var environment Environment

	currentEnvironmentId, err := FindCurrentId()

	// Get environment with the relevant Id
	selectQuery = "SELECT * FROM Environments WHERE Id = ?"
	row = db.Db.QueryRow(selectQuery, currentEnvironmentId)

	err = row.Scan(&environment.Id, &environment.Name, &environment.ApiKey, &environment.DebugMode, &environment.Position, &environment.CreatedAt, &environment.UpdatedAt)

	return &environment, err
}
func FindCurrentId() (int, error) {
	var selectQuery string
	var row *sql.Row
	var currentEnvironmentId int

	// Get current environment Id
	selectQuery = "SELECT EnvironmentId FROM CurrentEnvironment"
	row = db.Db.QueryRow(selectQuery)
	err := row.Scan(&currentEnvironmentId)
	if err != nil {
		log.Fatal("Failed to scan the current environment Id")
	}

	return currentEnvironmentId, err
}

// DefaultEnvironment
func UpdateDefault(environmentId int) error {

	// TODO:
	selectQuery := "UPDATE CurrentEnvironment SET EnvironmentId = ?"
	_, err := db.Db.Exec(selectQuery, environmentId)

	return err
}
func FindDefault() (*Environment, error) {
	var environment Environment
	defaultEnvironmentId, err := FindDefaultId()
	if err != nil {
		fmt.Println("Failed to FindDefaultId() inside FindDefault")
		log.Fatal(err)
	}

	selectQuery := "SELECT * FROM Environments WHERE Id = ?"
	row := db.Db.QueryRow(selectQuery, defaultEnvironmentId)

	err = row.Scan(&environment.Id, &environment.Name, &environment.ApiKey, &environment.DebugMode, &environment.Position, &environment.CreatedAt, &environment.UpdatedAt)
	return &environment, err
}
func FindDefaultId() (int, error) {
	var environmentId int
	selectQuery := "SELECT EnvironmentId FROM DefaultEnvironment"
	row := db.Db.QueryRow(selectQuery)
	err := row.Scan(&environmentId)

	return environmentId, err
}
