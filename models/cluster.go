package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Cluster is used by pop to map your clusters database table to your go code.
type Cluster struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Name  string    `json:"name" db:"name" form:"name"`
	Token string    `json:"-" db:"token"`

	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id"`
	OwnerType string    `json:"owner_type" db:"owner_type"`

	LastCheckIn time.Time `json:"last_check_in" db:"last_check_in"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Cluster) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

func (c *Cluster) BeforeCreate(tx *pop.Connection) error {
	// TODO: We should add a suffix for additional variations
	if c.Token == "" {
		var token uuid.UUID
		token, err := uuid.NewV4()

		if err != nil {
			return err
		}

		c.Token = token.String()
	}
	return nil
}

// Clusters is not required by pop and may be deleted
type Clusters []Cluster

// String is not required by pop and may be deleted
func (c Clusters) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Cluster) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Cluster) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Cluster) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
