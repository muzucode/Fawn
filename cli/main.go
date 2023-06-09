package cli

import (
	"database/sql"
	"fmt"
	"log"
	"muzucode/goroutines/server"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type Group struct {
	ID          string
	Name        string
	Description string
	Servers     server.ServerList
}

var db *sql.DB

func main() {

	// Load .env file vars
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	rootCmd := &cobra.Command{
		Use:   "fawn",
		Short: "Manage environments",
	}

	environmentCmd := &cobra.Command{
		Use:   "env",
		Short: "Manage environments",
	}
	rootCmd.AddCommand(environmentCmd)

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an environment",
		Run:   addGroup,
	}
	environmentCmd.AddCommand(addCmd)

	removeCmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove an environment",
		Run:   removeGroup,
	}

	environmentCmd.AddCommand(removeCmd)

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List environments",
		Run:   listGroups,
	}

	environmentCmd.AddCommand(listCmd)

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an environment",
		Run:   updateGroup,
	}

	environmentCmd.AddCommand(updateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}

func connectDatabase() {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := os.Getenv("DB_NAME")
	var err error

	// Open a connection to the MySQL database
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	log.Println("Connected to the MySQL database successfully.")
}
