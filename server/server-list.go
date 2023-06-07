package server

import (
	"log"
	"muzucode/goroutines/db"
)

type ServerList map[string]Server

var sl = make(ServerList)

func (sl ServerList) GetInstance() *ServerList {
	return &sl
}
func (sl ServerList) FindOne(id string) Server {
	return sl[id]
}
func (sl ServerList) Add(s *Server) {
	sl[s.Id] = *s
}
func (sl ServerList) Delete(s *Server) {

	delete(sl, (*s).Id)
}
func FetchUpstreamServers() (ServerList, error) {
	db, err := db.GetDBConnection()
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

	servers := make(ServerList)

	for rows.Next() {
		var s Server
		err := rows.Scan(&s.Id, &s.Address, &s.Port)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		servers[s.Id] = s
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return servers, nil
}
func LoadServerList() {
	var fetchedServers, err = FetchUpstreamServers()
	if err != nil {
		log.Fatal(err)
	}

	sl = fetchedServers
}
