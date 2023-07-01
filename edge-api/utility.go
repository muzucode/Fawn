package edge_api

import (
	"fmt"
	"strconv"
)

func toInt(value string) int {
	// Convert param string to int
	intValue, err := strconv.Atoi(value)

	// Handle conversion errors
	if err != nil {
		fmt.Println("Error: Failed to convert string to int")
	}

	return intValue
}
