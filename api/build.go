package api

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Builds struct {
	c *Client
}

type Build struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key"`
	Status     string         `json:"status"`
	RevisionID uint           `json:"revision_id"`
	Revision   Revision       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	Deleted    gorm.DeletedAt `json:"deleted,omitempty"`
}

func (c *Client) Builds() *Builds {
	return &Builds{c}
}

func (b *Builds) Create() {

}

func (b *Builds) Cancel() {

}

func (b *Builds) Get() {

}
