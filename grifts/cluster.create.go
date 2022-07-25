package grifts

import (
	"errors"
	"launchbox/models"

	"github.com/gofrs/uuid"
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("cluster", func() {

	Desc("create", "Generate a record to initialize a new cluster")
	Add("create", func(c *Context) error {
		if len(c.Args) == 0 {
			return errors.New("Please provide a name argument")
		}

		name := c.Args[0]

		ownerId, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
		if err != nil {
			return err
		}

		cluster := &models.Cluster{
			Name:    name,
			OwnerId: ownerId,
		}

		return models.DB.Create(cluster)
	})

})
