package server

type ServerList map[string]*Server

var serverList = make(ServerList)

func (sl ServerList) FindOne(id string) *Server {
	return serverList[id]
}
func (sl ServerList) FindAll() ServerList {
	return serverList
}
func (sl ServerList) Add(server *Server) {
	serverList[server.Id] = server
}
func (sl ServerList) Remove(server *Server) {
	delete(serverList, server.Id)
}
