package edge_api

import (
	"log"
	"muzucode/fawn/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartEdgeAPI() {
	r := gin.Default()

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
		serverId := toInt(c.Param("serverId"))

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

	r.GET("/servers/:serverId/files", func(c *gin.Context) {
		files, err := GetFilesFromServer(c)
		if err != nil {
			log.Fatal(err)
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"error": err,
			// })
		}

		// Return data
		c.JSON(http.StatusOK, files)
	})

	r.POST("/servers/:server", func(c *gin.Context) {
		var msg server.Message
		c.BindJSON(&msg)
		// make post request to upstream server here
		c.JSON(200, 1)
	})

	// GET - Test
	r.GET("/test", func(c *gin.Context) {
		_, err := server.ServerTest()
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9024")
}
