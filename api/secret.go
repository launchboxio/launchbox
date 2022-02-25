package api

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"time"
)

type Secrets struct {
	c *Client
}

type Secret struct {
	ID   uuid.UUID `gorm:"type:uuid;primary_key"`
	Path string    `json:"path"`
	Type string    `json:"type"`

	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	Deleted   gorm.DeletedAt `json:"deleted,omitempty"`
}
