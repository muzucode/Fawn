package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func getDBConnection() (*sql.DB, error) {
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
