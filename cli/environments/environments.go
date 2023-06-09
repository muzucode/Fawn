package environments

import (
	"bufio"
	"fmt"
	"log"
	db "muzucode/goroutines/database"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Environment struct {
	Id        string
	AppName   string
	ApiKey    string
	DebugMode bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AddEnvironment(cmd *cobra.Command, args []string) {
	// Prompt the user for environment details
	var environment Environment
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter App Name: ")
	if scanner.Scan() {
		environment.AppName = scanner.Text()
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
	insertQuery := "INSERT INTO Environments (AppName, ApiKey, DebugMode) VALUES (?, ?, ?)"
	_, err := db.Db.Exec(insertQuery, environment.AppName, environment.ApiKey, environment.DebugMode)
	if err != nil {
		log.Printf("Failed to add environment: %v", err)
		return
	}

	fmt.Println("Environment added successfully.")
}

func RemoveEnvironment(cmd *cobra.Command, args []string) {
	// Prompt the user for the environment Id
	var environmentId string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter Environment Id: ")
	if scanner.Scan() {
		environmentId = scanner.Text()
	} else {
		log.Println("Failed to read input:", scanner.Err())
		return
	}

	// Delete the environment from the database
	deleteQuery := "DELETE FROM Environments WHERE Id = ?"
	_, err := db.Db.Exec(deleteQuery, environmentId)
	if err != nil {
		log.Printf("Failed to remove environment: %v", err)
		return
	}

	fmt.Println("Environment removed successfully.")
}

func ListEnvironments(cmd *cobra.Command, args []string) {
	// Retrieve all environments from the database
	fmt.Println("Getting all environments...")
	selectQuery := "SELECT * FROM Environments"
	rows, err := db.Db.Query(selectQuery)
	if err != nil {
		log.Printf("Failed to retrieve environment information: %v", err)
		return
	}
	defer rows.Close()

	fmt.Println("All Environments:")
	for rows.Next() {
		var environment Environment
		err := rows.Scan(&environment.Id, &environment.AppName, &environment.ApiKey, &environment.DebugMode, &environment.CreatedAt, &environment.UpdatedAt)
		if err != nil {
			log.Printf("Failed to retrieve environment information while scanning rows: %v", err)
			continue
		}

		fmt.Printf("Id: %s\n", environment.Id)
		fmt.Printf("App Name: %s\n", environment.AppName)
		fmt.Printf("API Key: %s\n", environment.ApiKey)
		fmt.Printf("Debug Mode: %v\n", environment.DebugMode)
		fmt.Printf("Created At: %v\n", environment.CreatedAt)
		fmt.Printf("Updated At: %v\n", environment.UpdatedAt)
		fmt.Println("---------------")
	}
}