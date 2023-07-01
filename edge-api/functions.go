package edge_api

import (
	"fmt"
	"muzucode/fawn/servers"
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
