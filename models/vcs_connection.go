package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// VcsConnection is used by pop to map your vcs_connections database table to your go code.
type VcsConnection struct {
	ID uuid.UUID `json:"id" db:"id"`

	Provider       string `json:"provider" db:"provider"`
	Hostname       string `json:"hostname" db:"hostname"`
	Name           string `json:"name" db:"name"`
	Email          string `json:"email" db:"email"`
	Nickname       string `json:"nickname" db:"nickname"`
	ProviderUserId string `json:"provider_user_id" db:"provider_user_id"`
	AccessToken    string `json:"access_token" db:"access_token"`
	ExpiresAt      string `json:"expires_at" db:"expires_at"`
	RefreshToken   string `json:"refresh_token" db:"refresh_token"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (v VcsConnection) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// VcsConnections is not required by pop and may be deleted
type VcsConnections []VcsConnection

// String is not required by pop and may be deleted
func (v VcsConnections) String() string {
	jv, _ := json.Marshal(v)
	return string(jv)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (v *VcsConnection) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (v *VcsConnection) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (v *VcsConnection) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
