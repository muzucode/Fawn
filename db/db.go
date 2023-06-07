package db

import (
	"database/sql"
	"log"
	"muzucode/goroutines/server"

	_ "github.com/go-sql-driver/mysql"
)

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

func FetchUpstreamServers() (map[string]server.Server, error) {
	db, err := GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, address, port FROM upstream_servers")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	servers := make(map[string]server.Server)

	for rows.Next() {
		var server server.Server
		err := rows.Scan(&server.Id, &server.Address, &server.Port)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		servers[server.Id] = server
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return servers, nil
}
