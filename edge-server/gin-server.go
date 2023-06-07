package edgeserver

import (
	"muzucode/goroutines/protocol"
	serverlist "muzucode/goroutines/server-list"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	r.GET("/servers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/servers/:server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/servers/:server", func(c *gin.Context) {
		var msg protocol.Message
		c.BindJSON(&msg)
		// make post request to upstream server here
		c.JSON(200, 1)
	})
	r.POST("/servers/:server", func(c *gin.Context) {
		var msg protocol.Message
		c.BindJSON(&msg)
		// make post request to upstream server here
		c.JSON(200, 1)
	})
	r.DELETE("/servers/:serverId", func(c *gin.Context) {
		var msg protocol.Message
		delete(*serverlist.GetInstance(), c.Param("serverId"))
		c.BindJSON(&msg)
		// make post request to upstream server here
		c.JSON(200, 1)
	})
	r.Run(":9823") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
