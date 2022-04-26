package api

import (
	"fmt"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Secrets struct {
	c *Client
}

type SecretType string

const (
	SecretTypeOrganization SecretType = "organization"
	SecretTypeProject      SecretType = "project"
	SecretTypeApplication  SecretType = "application"
)

type Secret struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Path       string    `json:"path"`
	ObjectType string    `json:"object_type"`
	ObjectId   string    `json:"object_id"`
	Name       string    `json:"name"`
	Version    int32     `json:"version"`
	Value      string    `json:"value" gorm:"-"`

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	Deleted   gorm.DeletedAt `json:"deleted,omitempty"`
}

type SecretListResponse struct {
	Secrets []Secret `json:"secrets"`
}

func (c *Client) Secrets() *Secrets {
	return &Secrets{c}
}

func (s *Secrets) List(objectType string, objectId string) (SecretListResponse, error) {
	url := fmt.Sprintf("/%ss/%s/secrets", objectType, objectId)
	secrets := SecretListResponse{}
	err := s.c.get(url, nil, &secrets)
	return secrets, err
}

func (s *Secrets) Find(secret *Secret) error {

}

func (s *Secrets) Create(secret *Secret) error {
	url := fmt.Sprintf("/%ss/%s/secrets", secret.ObjectType, secret.ObjectId)
	err := s.c.post(url, secret)
	return err
}

func (s *Secrets) Update(secret *Secret) error {
	url := fmt.Sprintf("/%ss/%s/secrets/%s", secret.ObjectType, secret.ObjectId, secret.ID)
	err := s.c.post(url, secret)
	return err
}

func (s *Secrets) Delete(secret *Secret) error {
	url := fmt.Sprintf("/%ss/%s/secrets/%s", secret.ObjectType, secret.ObjectId, secret.ID)
	err := s.c.post(url, secret)
	return err
}
