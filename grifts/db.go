package grifts

import (
	"launchbox/models"

	"github.com/gobuffalo/pop/v6"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		return models.DB.Transaction(func(tx *pop.Connection) error {
			err := tx.TruncateAll()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Set("tx", tx)
			if err = grift.Run("seed:users", c); err != nil {
				return errors.WithStack(err)
			}

			if err = grift.Run("seed:projects", c); err != nil {
				return errors.WithStack(err)
			}

			if err = grift.Run("seed:clusters", c); err != nil {
				return errors.WithStack(err)
			}

			return nil
		})
	})

	grift.Add("seed:users", func(c *grift.Context) error {
		user := &models.User{
			Email:    "admin@launchbox.com",
			Password: "password123",
		}
		if err := models.DB.Create(user); err != nil {
			return err
		}

		c.Set("user_id", user.ID)
		return nil
	})

	grift.Add("seed:applications", func(c *grift.Context) error {
		application := &models.Application{
			Name: "test-app",
		}
		if err := models.DB.Create(application); err != nil {
			return err
		}
		// Create some projects?
		return nil
	})

	grift.Add("seed:clusters", func(c *grift.Context) error {
		cluster := &models.Cluster{
			Name: "default",
		}
		return models.DB.Create(cluster)
	})

})
