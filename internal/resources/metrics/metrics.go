package metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/internal/config"
	"net/http"
)

type Metrics struct {
	config config.PrometheusConfig
}

var NamedQueries = map[string]string{
	"APP_CPU_USAGE":                    "",
	"APP_MEMORY_USAGE":                 "",
	"APP_HTTP_REQUESTS_PER_MINUTE":     "",
	"APP_HTTP_LATENCY_PER_MINUTE":      "",
	"APP_HTTP_ERRORS_PER_MINUTE":       "",
	"PROJECT_CPU_USAGE":                "",
	"PROJECT_MEM_USAGE":                "",
	"PROJECT_HTTP_REQUESTS_PER_MINUTE": "",
	"PROJECT_HTTP_LATENCY_PER_MINUTE":  "",
	"PROJECT_HTTP_ERRORS_PER_MINUTE":   "",
}

func New(conf config.PrometheusConfig) *Metrics {
	return &Metrics{config: conf}
}

func verifyMetricsEnabled(conf *config.PrometheusConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !conf.Enabled {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Metrics feature not enabled"})
			return
		}
	}
}

func (m *Metrics) Query(c *gin.Context) {
	query := fmt.Sprintf(NamedQueries[c.Query("query")])

	c.JSON(http.StatusOK, gin.H{"query": query})
}
