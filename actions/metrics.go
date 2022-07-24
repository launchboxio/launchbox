package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// MetricsQuery default implementation.
func MetricsQuery(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("metrics/query.html"))
}

