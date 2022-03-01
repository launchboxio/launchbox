package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Metrics struct {
	PrometheusUrl string
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

func (m *Metrics) Register(r *gin.Engine) {
	group := r.Group("/metrics")
	group.GET("", m.Query)

}

func (m *Metrics) Query(c *gin.Context) {
	query := fmt.Sprintf(NamedQueries[c.Query("query")])

	c.JSON(http.StatusOK, gin.H{"query": query})
}
