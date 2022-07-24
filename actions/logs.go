package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// LogsQuery default implementation.
func LogsQuery(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("logs/query.html"))
}
