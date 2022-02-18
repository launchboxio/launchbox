package server

import "github.com/gin-gonic/gin"

type ServerOpts struct {
	Port int64
}

func New(opts *ServerOpts) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
