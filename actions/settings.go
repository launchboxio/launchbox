package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// SettingsIndex default implementation.
func SettingsIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("settings/index.html"))
}
