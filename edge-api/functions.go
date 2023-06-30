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
			&s.Address,
			&s.PrivateKeyPath,
			&s.GroupId,
			&s.Description,
			&s.DistributionName,
			&s.DistributionVersion,
		)
		if err != nil {
			fmt.Printf("%v\n", err)
			return nil, err
		}

		data = append(data, s)
	}

	return data, nil
}
