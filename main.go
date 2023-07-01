package main

import (
	"flag"
	"log"
	db "muzucode/fawn/database"
	edge_api "muzucode/fawn/edge-api"
	"muzucode/fawn/environments"
	"muzucode/fawn/groups"
	"muzucode/fawn/servers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	// Load .env file vars
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Parse command-line flags
	apiFlag := flag.Bool("server", false, "Start the API server")
	flag.Parse()

	// Connect to DB
	db.ConnectToDatabase()

	// Create root command
	rootCmd := &cobra.Command{
		Use:   "fawn",
		Short: "Manage groups",
	}

	// Load commands
	groups.LoadGroupCommands(rootCmd)
	environments.LoadEnvironmentCommands(rootCmd)
	servers.LoadServerCommands(rootCmd)

	// Check the flag value and run the appropriate code
	if *apiFlag {
		edge_api.StartEdgeAPI()
	} else {
		// Handle root command run
		if err := rootCmd.Execute(); err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
	}
}
