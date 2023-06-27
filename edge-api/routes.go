package edge_api

import (
	"log"
	"muzucode/fawn/server"
	"net/http"

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
				"message": err,
			})
		}

		c.JSON(http.StatusOK, data)
	})

	r.GET("/servers/:server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/servers/:server", func(c *gin.Context) {
		var msg server.Message
		c.BindJSON(&msg)
		// make post request to upstream server here
		c.JSON(200, 1)
	})
	r.Run(":9024")
}
