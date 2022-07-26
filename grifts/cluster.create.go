package grifts

import (
	"launchbox/models"
	"time"

	"github.com/yelinaung/go-haikunator"

	"github.com/gofrs/uuid"
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("cluster", func() {

	Desc("create", "Generate a record to initialize a new cluster")
	Add("create", func(c *Context) error {
		var name string
		if len(c.Args) == 0 {
			gen := haikunator.New(time.Now().UTC().UnixNano())
			name = gen.Haikunate()
		} else {
			name = c.Args[0]
		}

		ownerId, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
		if err != nil {
			return err
		}

		cluster := &models.Cluster{
			Name:    name,
			OwnerId: ownerId,
		}
		if err = models.DB.Create(cluster); err != nil {
			return err
		}

		agent := &models.Agent{
			Cluster: cluster,
		}

		if err = models.DB.Create(agent); err != nil {
			return err
		}

		return nil
	})

})
