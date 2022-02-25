package api

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Apps struct {
	c *Client
}

type Application struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Namespace string         `gorm:"uniqueIndex" json:"namespace,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	Deleted   gorm.DeletedAt `json:"deleted,omitempty"`
}

func (a *Application) GetNamespace() string {
	return fmt.Sprintf("lb-app-%s", a.Name)
}

type ApplicationListResponse struct {
	Applications []Application `json:"applications"`
}

func (c *Client) Apps() *Apps {
	return &Apps{c}
}

func (a *Apps) List() (ApplicationListResponse, error) {
	apps := ApplicationListResponse{}
	err := a.c.get("/applications", nil, &apps)
	return apps, err
}

func (a *Apps) Create(application *Application) error {
	err := a.c.post("/applications", application)
	return err
}

func (a *Apps) Update(application *Application) error {
	return a.c.put(fmt.Sprintf("/applications/%d", application.ID), application)
}

func (a *Apps) Delete(applicationId uint) error {
	return a.c.delete(fmt.Sprintf("/applications/%d", applicationId))
}

func (a *Apps) Find(applicationId uint) (*Application, error) {
	app := &Application{}
	err := a.c.get(fmt.Sprintf("/applications/%d", applicationId), nil, app)
	return app, err
}
