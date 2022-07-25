package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// TracesQuery default implementation.
func TracesQuery(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("traces/query.html"))
}
