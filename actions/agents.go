package actions

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/x/responder"

	"launchbox/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Agent)
// DB Table: Plural (agents)
// Resource: Plural (Agents)
// Path: Plural (/agents)
// View Template Folder: Plural (/templates/agents/)

// AgentsResource is the resource for the Agent model
type AgentsResource struct {
	buffalo.Resource
}

// List gets all Agents. This function is mapped to the path
// GET /agents
func (v AgentsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	agents := &models.Agents{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Agents from the DB
	if err := q.All(agents); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("agents", agents)
		return c.Render(http.StatusOK, r.HTML("agents/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(agents))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(agents))
	}).Respond(c)
}

// Show gets the data for one Agent. This function is mapped to
// the path GET /agents/{agent_id}
func (v AgentsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Agent
	agent := &models.Agent{}

	// To find the Agent the parameter agent_id is used.
	if err := tx.Find(agent, c.Param("agent_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("agent", agent)

		return c.Render(http.StatusOK, r.HTML("agents/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(agent))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(agent))
	}).Respond(c)
}

// New renders the form for creating a new Agent.
// This function is mapped to the path GET /agents/new
func (v AgentsResource) New(c buffalo.Context) error {
	c.Set("agent", &models.Agent{})

	return c.Render(http.StatusOK, r.HTML("agents/new.plush.html"))
}

// Create adds a Agent to the DB. This function is mapped to the
// path POST /agents
func (v AgentsResource) Create(c buffalo.Context) error {
	// Allocate an empty Agent
	agent := &models.Agent{}

	// Bind agent to the html form elements
	if err := c.Bind(agent); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(agent)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("agent", agent)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("agents/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "agent.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/agents/%v", agent.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(agent))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(agent))
	}).Respond(c)
}

// Edit renders a edit form for a Agent. This function is
// mapped to the path GET /agents/{agent_id}/edit
func (v AgentsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Agent
	agent := &models.Agent{}

	if err := tx.Find(agent, c.Param("agent_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("agent", agent)
	return c.Render(http.StatusOK, r.HTML("agents/edit.plush.html"))
}

// Update changes a Agent in the DB. This function is mapped to
// the path PUT /agents/{agent_id}
func (v AgentsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Agent
	agent := &models.Agent{}

	if err := tx.Find(agent, c.Param("agent_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Agent to the html form elements
	if err := c.Bind(agent); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(agent)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("agent", agent)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("agents/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "agent.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/agents/%v", agent.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(agent))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(agent))
	}).Respond(c)
}

// Destroy deletes a Agent from the DB. This function is mapped
// to the path DELETE /agents/{agent_id}
func (v AgentsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Agent
	agent := &models.Agent{}

	// To find the Agent the parameter agent_id is used.
	if err := tx.Find(agent, c.Param("agent_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(agent); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "agent.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/agents")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(agent))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(agent))
	}).Respond(c)
}
