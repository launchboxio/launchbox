package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// LogsQuery default implementation.
func LogsQuery(c buffalo.Context) error {
	// Generate the query
	// Get our Loki client
	// Execute the HTTP request to get logs
	// Return logs to the user
	return c.Render(http.StatusOK, r.HTML("logs/query.html"))
}

func LogsStream(c buffalo.Context) error {
	// Query logs
	// Return a stream to the user
	return c.Render(http.StatusOK, r.HTML("logs/query.html"))
}
