package edge_api

import (
	"fmt"
	"muzucode/fawn/server"
	"muzucode/fawn/servers"

	"github.com/gin-gonic/gin"
)

func FindAllServers() ([]servers.Server, error) {
	rows, err := servers.FindAll()
	if err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}
	defer rows.Close()

	var data []servers.Server

	for rows.Next() {
		var s servers.Server
		err := rows.Scan(
			&s.Id,
			&s.Name,
			&s.AddressIPv4,
			&s.PrivateKeyPath,
			&s.GroupId,
			&s.Description,
			&s.DistributionName,
			&s.DistributionVersion,
			&s.AddressIPv6,
		)
		if err != nil {
			fmt.Printf("%v\n", err)
			return nil, err
		}

		data = append(data, s)
	}

	return data, nil
}

func FindOneServer(serverId int) (servers.Server, error) {
	s, err := servers.FindOne(serverId)
	if err != nil {
		fmt.Printf("%v\n", err)
		return servers.Server{}, err
	}

	return *s, nil
}

func GetFilesFromServer(c *gin.Context) ([]string, error) {
	var s servers.Server

	// Convert param string to int
	serverId := toInt(c.Param("serverId"))

	// Find a server
	s, err := FindOneServer(serverId)
	if err != nil {
		return nil, err
	}

	// Get files in given directory for the given server
	var files []string
	files, err = server.GetFilesInDir(s, "/")
	for i := 0; i < len(files); i++ {
		fmt.Println(files[i])
	}
	// Handle error
	if err != nil {
		return nil, err
	}

	return files, err
}

func InsertOneServer(server *servers.Server) error {
	err := servers.InsertOne(server)
	if err != nil {
		return err
	}
	return nil
}
