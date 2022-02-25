package server

import "github.com/gin-gonic/gin"

type Logs struct {
}

func (l *Logs) Register(r *gin.Engine) {
	r.GET("/logs", l.Read)
}

func (l *Logs) Read(c *gin.Context) {

}