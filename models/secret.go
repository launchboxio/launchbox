package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Secret is used by pop to map your secrets database table to your go code.
type Secret struct {
	ID uuid.UUID `json:"id" db:"id"`

	Name      string `json:"name" db:"name"`
	Sensitive bool   `json:"sensitive" db:"sensitive"`

	OwnerType string    `json:"owner_type" db:"owner_type"`
	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id"`
	ClusterID uuid.UUID `json:"cluster_id" db:"cluster_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (s Secret) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Secrets is not required by pop and may be deleted
type Secrets []Secret

// String is not required by pop and may be deleted
func (s Secrets) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Secret) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Secret) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Secret) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
