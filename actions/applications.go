package actions

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"launchbox/models"
)

type ApplicationsResource struct {
	buffalo.Resource
}

// swagger:route GET /applications ApplicationsResource.List
func (v ApplicationsResource) List(c buffalo.Context) error {
	user := c.Value("current_user").(*models.User)

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	applications := &models.Applications{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Applications from the DB
	if err := q.Where("user_id = ?", user.ID).All(applications); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("applications", applications)
		return c.Render(http.StatusOK, r.HTML("applications/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(applications))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(applications))
	}).Respond(c)
}

// Show gets the data for one Application. This function is mapped to
// the path GET /applications/{application_id}
func (v ApplicationsResource) Show(c buffalo.Context) error {
	application := c.Value("application")

	return responder.Wants("html", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.HTML("applications/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(application))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(application))
	}).Respond(c)
}

// New renders the form for creating a new Application.
// This function is mapped to the path GET /applications/new
func (v ApplicationsResource) New(c buffalo.Context) error {
	c.Set("application", &models.Application{})

	return c.Render(http.StatusOK, r.HTML("applications/new.plush.html"))
}

// Create adds a Application to the DB. This function is mapped to the
// path POST /applications
func (v ApplicationsResource) Create(c buffalo.Context) error {
	user := c.Value("current_user").(*models.User)

	// Allocate an empty Application
	application := &models.Application{
		User: user,
	}

	// Bind application to the html form elements
	if err := c.Bind(application); err != nil {
		return err
	}

	// TODO: Validate the actual existence of application.Name. For some reason,
	// even an empty name will pass validation

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(application)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("application", application)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("applications/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "application.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/applications/%v", application.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(application))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(application))
	}).Respond(c)
}

// Edit renders a edit form for a Application. This function is
// mapped to the path GET /applications/{application_id}/edit
func (v ApplicationsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Application
	application := &models.Application{}

	if err := tx.Find(application, c.Param("application_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("application", application)
	return c.Render(http.StatusOK, r.HTML("applications/edit.plush.html"))
}

// Update changes a Application in the DB. This function is mapped to
// the path PUT /applications/{application_id}
func (v ApplicationsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Application
	application := &models.Application{}

	if err := tx.Find(application, c.Param("application_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Application to the html form elements
	if err := c.Bind(application); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(application)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("application", application)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("applications/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "application.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/applications/%v", application.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(application))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(application))
	}).Respond(c)
}

// Destroy deletes a Application from the DB. This function is mapped
// to the path DELETE /applications/{application_id}
func (v ApplicationsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Application
	application := &models.Application{}

	// To find the Application the parameter application_id is used.
	if err := tx.Find(application, c.Param("application_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(application); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "application.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/applications")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(application))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(application))
	}).Respond(c)
}

func SetCurrentApplication(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		uid := c.Session().Get("current_user_id").(uuid.UUID)

		tx := c.Value("tx").(*pop.Connection)
		application := models.Application{}

		// To find the Application the parameter application_id is used.
		if err := tx.Where("user_id = ?", uid.String()).Where("id = ?", c.Param("application_id")).Eager("Projects").First(&application); err != nil {
			log.Println(err)
			return responder.Wants("html", func(c buffalo.Context) error {
				c.Flash().Add("error", err.Error())

				return c.Redirect(302, "/applications")
			}).Wants("json", func(c buffalo.Context) error {
				return c.Error(http.StatusNotFound, errors.New("Application not found"))
			}).Respond(c)
		}

		c.Set("application", application)

		return next(c)
	}
}
