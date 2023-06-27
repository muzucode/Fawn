package server

import (
	"muzucode/fawn/servers"
)

type ServerList map[int]*servers.Server

var sl = make(ServerList)

func (sl ServerList) GetInstance() *ServerList {
	return &sl
}
func (sl ServerList) FindOne(serverId int) *servers.Server {
	return sl[serverId]
}
func (sl ServerList) Add(s *servers.Server) {
	sl[s.Id] = s
}
func (sl ServerList) Delete(s *servers.Server) {
	delete(sl, (*s).Id)
}
