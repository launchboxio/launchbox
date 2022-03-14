package server

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type Logs struct {
	Config LokiConfig
}

func (l *Logs) Register(r *gin.Engine) {
	r.GET("/logs", l.Read)
}

func (l *Logs) Read(c *gin.Context) {
	if !l.Config.Enabled {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Logs feature not enabled"})
		return
	}
	chanStream := make(chan int, 10)
	go func() {
		defer close(chanStream)
		for i := 0; i < 5; i++ {
			chanStream <- i
			time.Sleep(time.Second * 1)
		}
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}
