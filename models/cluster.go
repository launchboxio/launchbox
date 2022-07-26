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
	ID       uuid.UUID `json:"id" db:"id"`
	Name     string    `json:"name" db:"name" form:"name"`
	Provider string    `json:"provider" db:"provider" form:"provider"`
	Region   string    `json:"region" db:"region" form:"region"`
	Version  string    `json:"version" db:"version" form:"version"`
	Token    string    `json:"-" db:"token"`

	Agents       []Agent       `json:"agents,omitempty" has_many:"agents"`
	Applications []Application `json:"applications,omitempty" many_to_many:"cluster_applications"`

	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id"`
	OwnerType string    `json:"owner_type" db:"owner_type"`
	Managed   bool      `json:"managed" db:"-"`

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

func (c *Cluster) AfterFind(tx *pop.Connection) error {
	if c.OwnerId.String() == "00000000-0000-0000-0000-000000000000" {
		c.Managed = true
	} else {
		c.Managed = false
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
