package server

type ServerList map[string]*Server

var serverList = make(ServerList)

func FindOne(id string) *Server {
	return serverList[id]
}
func FindAll() ServerList {
	return serverList
}
func Add(server *Server) {
	serverList[server.Id] = server
}
func Remove(server *Server) {
	delete(serverList, server.Id)
}
