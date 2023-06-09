package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func GetDBConnection() (*sql.DB, error) {
	// TODO: Declare a DB server
	db, err := sql.Open("mysql", "user:password@tcp(host:port)/database")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}

func ConnectToDatabase() {
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := os.Getenv("DB_NAME")
	var err error

	// Open a connection to the MySQL database
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	log.Println("Connected to the MySQL database successfully.")
}
