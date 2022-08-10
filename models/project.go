package models

import (
	"encoding/json"
	"github.com/gosimple/slug"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Blocked ports:
// - 15000: used by the Envoy admin interface exposed over localhost
// - 15001: used by the Envoy outbound listener to accept and proxy outbound traffic sent by applications within the pod
// - 15003: used by the Envoy inbound listener to accept and proxy inbound traffic entering the pod destined to applications within the pod
// - 15010: used by the Envoy inbound Prometheus listener to accept and proxy inbound traffic pertaining to scraping Envoy’s Prometheus metrics
// - 15901: used by Envoy to serve rewritten HTTP liveness probes
// - 15902: used by Envoy to serve rewritten HTTP readiness probes
// - 15903: used by Envoy to serve rewritten HTTP startup probes

// Blocked UID - 1500

// Project is used by pop to map your projects database table to your go code.
type Project struct {
	ID     uuid.UUID `json:"id" db:"id"`
	Name   string    `json:"name" db:"name" form:"name"`
	Status string    `json:"status" db:"status"`
	Slug   string    `json:"slug" db:"slug"`

	ApplicationID uuid.UUID    `json:"application_id" db:"application_id"`
	Application   *Application `json:"application,omitempty" belongs_to:"application"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Project) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Projects is not required by pop and may be deleted
type Projects []Project

// String is not required by pop and may be deleted
func (p Projects) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

func (p *Project) BeforeCreate(tx *pop.Connection) error {
	if p.Slug == "" {
		p.Slug = slug.Make(p.Name)
	}
	return nil
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Project) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
