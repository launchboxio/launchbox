package api

import (
	"gorm.io/gorm"
	"time"
)

type Projects struct {
	c *Client
}

type Project struct {
	gorm.Model
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	Repo          string         `json:"repo"`
	Branch        string         `json:"branch,omitempty"`
	ApplicationID uint           `json:"application_id"`
	Application   Application    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	Deleted       gorm.DeletedAt `json:"deleted,omitempty"`
}

func (c *Client) Projects() *Projects {
	return &Projects{c}
}

func (a *Projects) List() {

}

func (a *Projects) Create() {

}

func (a *Projects) Update() {

}

func (a *Projects) Delete() {

}

func (a *Projects) Find() {

}
