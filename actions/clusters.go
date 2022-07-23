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
// Model: Singular (Cluster)
// DB Table: Plural (clusters)
// Resource: Plural (Clusters)
// Path: Plural (/clusters)
// View Template Folder: Plural (/templates/clusters/)

// ClustersResource is the resource for the Cluster model
type ClustersResource struct {
	buffalo.Resource
}

// List gets all Clusters. This function is mapped to the path
// GET /clusters
func (v ClustersResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	clusters := &models.Clusters{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Clusters from the DB
	if err := q.All(clusters); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)

		c.Set("clusters", clusters)
		return c.Render(http.StatusOK, r.HTML("clusters/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(clusters))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(clusters))
	}).Respond(c)
}

// Show gets the data for one Cluster. This function is mapped to
// the path GET /clusters/{cluster_id}
func (v ClustersResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cluster
	cluster := &models.Cluster{}

	// To find the Cluster the parameter cluster_id is used.
	if err := tx.Find(cluster, c.Param("cluster_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("cluster", cluster)

		return c.Render(http.StatusOK, r.HTML("clusters/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(cluster))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(cluster))
	}).Respond(c)
}

// New renders the form for creating a new Cluster.
// This function is mapped to the path GET /clusters/new
func (v ClustersResource) New(c buffalo.Context) error {
	c.Set("cluster", &models.Cluster{})

	return c.Render(http.StatusOK, r.HTML("clusters/new.plush.html"))
}

// Create adds a Cluster to the DB. This function is mapped to the
// path POST /clusters
func (v ClustersResource) Create(c buffalo.Context) error {
	// Allocate an empty Cluster
	cluster := &models.Cluster{}

	// Bind cluster to the html form elements
	if err := c.Bind(cluster); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(cluster)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the new.html template that the user can
			// correct the input.
			c.Set("cluster", cluster)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("clusters/new.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "cluster.created.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/clusters/%v", cluster.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.JSON(cluster))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusCreated, r.XML(cluster))
	}).Respond(c)
}

// Edit renders a edit form for a Cluster. This function is
// mapped to the path GET /clusters/{cluster_id}/edit
func (v ClustersResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cluster
	cluster := &models.Cluster{}

	if err := tx.Find(cluster, c.Param("cluster_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	c.Set("cluster", cluster)
	return c.Render(http.StatusOK, r.HTML("clusters/edit.plush.html"))
}

// Update changes a Cluster in the DB. This function is mapped to
// the path PUT /clusters/{cluster_id}
func (v ClustersResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cluster
	cluster := &models.Cluster{}

	if err := tx.Find(cluster, c.Param("cluster_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Cluster to the html form elements
	if err := c.Bind(cluster); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(cluster)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("cluster", cluster)

			return c.Render(http.StatusUnprocessableEntity, r.HTML("clusters/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "cluster.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/clusters/%v", cluster.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(cluster))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(cluster))
	}).Respond(c)
}

// Destroy deletes a Cluster from the DB. This function is mapped
// to the path DELETE /clusters/{cluster_id}
func (v ClustersResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Cluster
	cluster := &models.Cluster{}

	// To find the Cluster the parameter cluster_id is used.
	if err := tx.Find(cluster, c.Param("cluster_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(cluster); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "cluster.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/clusters")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(cluster))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(cluster))
	}).Respond(c)
}
