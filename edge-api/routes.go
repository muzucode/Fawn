package edge_api

import (
	"fmt"
	"log"
	"muzucode/fawn/server"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartEdgeAPI() {
	r := gin.Default()

	// GET - ServerTest
	r.GET("/servertest", func(c *gin.Context) {
		_, err := server.ServerTest()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// GET - Servers
	r.GET("/servers", func(c *gin.Context) {

		data, err := FindAllServers()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/servers/:serverId", func(c *gin.Context) {
		// Convert param string to int
		serverIdStr := c.Param("serverId")
		serverId, err := strconv.Atoi(serverIdStr)

		// Handle conversion errors
		if err != nil {
			fmt.Println("Error: Failed to convert string to int")
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}

		// Find a server
		data, err := FindOneServer(serverId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": err,
			})
		}

		// Return data
		c.JSON(http.StatusOK, data)
	})
	r.POST("/servers/:server", func(c *gin.Context) {
		var msg server.Message
		c.BindJSON(&msg)
		// make post request to upstream server here
		c.JSON(200, 1)
	})
	r.Run(":9024")
}
