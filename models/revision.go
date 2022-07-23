package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Revision is used by pop to map your revisions database table to your go code.
type Revision struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Status    string    `json:"status"`
	CommitSha string    `json:"commit_sha"`

	ProjectID uuid.UUID `json:"project_id"`
	Project   Project   `json:"project,omitempty" belongs_to:"project"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (r Revision) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Revisions is not required by pop and may be deleted
type Revisions []Revision

// String is not required by pop and may be deleted
func (r Revisions) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (r *Revision) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (r *Revision) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (r *Revision) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
