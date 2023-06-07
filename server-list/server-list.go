package serverlist

import (
	"log"
	"muzucode/goroutines/db"
	"muzucode/goroutines/server"
)

type ServerList map[string]server.Server

var serverList = make(ServerList)

func LoadServerList() {
	var fetchedServers, err = db.FetchUpstreamServers()
	if err != nil {
		log.Fatal(err)
	}

	serverList = fetchedServers

}

func GetInstance() *ServerList {
	return &serverList
}

func (sl ServerList) FindOne(id string) server.Server {
	return serverList[id]
}
func (sl ServerList) Add(s *server.Server) {
	serverList[s.Id] = *s
}
func (sl ServerList) Delete(s *server.Server) {

	delete(serverList, (*s).Id)
}
