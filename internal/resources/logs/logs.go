package logs

import (
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/internal/config"
	"io"
	"net/http"
	"time"
)

type LogsController struct {
	config *config.LokiConfig
}

func New(loki *config.LokiConfig) *LogsController {
	ctrl := &LogsController{
		config: loki,
	}

	return ctrl
}

func (l *LogsController) Read(c *gin.Context) {
	if !l.config.Enabled {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Logs feature not enabled"})
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
