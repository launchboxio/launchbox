package middleware

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"launchbox/models"
	"launchbox/pkg/scopes"
	"net/http"
)

func SetCurrentApplication(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		uid := c.Session().Get("current_user_id").(uuid.UUID)

		tx := c.Value("tx").(*pop.Connection)
		application := models.Application{}

		// To find the Application the parameter application_id is used.
		if err := tx.Scope(scopes.UserScope(uid.String())).Where("id = ?", c.Param("application_id")).Eager().First(&application); err != nil {
			return responder.Wants("html", func(c buffalo.Context) error {
				c.Flash().Add("error", "application.find.error")

				return c.Redirect(302, "/applications")
			}).Wants("json", func(c buffalo.Context) error {
				return c.Error(http.StatusNotFound, errors.New("Application not found"))
			}).Respond(c)
		}

		c.Set("application", application)

		return next(c)
	}
}
