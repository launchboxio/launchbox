package models

import (
	"encoding/json"
	"time"

	"github.com/yelinaung/go-haikunator"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Application is used by pop to map your applications database table to your go code.
type Application struct {
	ID uuid.UUID `json:"id" db:"id"`

	Name      string `json:"name" db:"name" form:"name"`
	Namespace string `json:"namespace" db:"namespace"`

	Projects []Project `json:"projects,omitempty" has_many:"projects"`
	Clusters []Cluster `json:"clusters" many_to_many:"cluster_applications" db:"-"`

	UserID    uuid.UUID `json:"-" db:"user_id"`
	User      *User     `json:"tree,omitempty" belongs_to:"user"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (a Application) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Applications is not required by pop and may be deleted
type Applications []Application

// String is not required by pop and may be deleted
func (a Applications) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

func (a *Application) BeforeCreate(tx *pop.Connection) error {
	// TODO: We should add a suffix for additional variations
	if a.Namespace == "" {
		gen := haikunator.New(time.Now().UTC().UnixNano())
		a.Namespace = gen.Haikunate()
	}
	return nil
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Application) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Application) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Application) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
