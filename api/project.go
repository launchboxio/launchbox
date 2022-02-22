package api

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Projects struct {
	c *Client
}

type Project struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	Repo          string         `json:"repo"`
	Branch        string         `json:"branch,omitempty"`
	ApplicationID uint           `json:"application_id"`
	Application   Application    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	Deleted       gorm.DeletedAt `json:"deleted,omitempty"`
}

type ProjectListResponse struct {
	Projects []Project `json:"projects"`
}

type ProjectListOptions struct {
	ApplicationId string
}

func (c *Client) Projects() *Projects {
	return &Projects{c}
}

func (p *Projects) List(opts *ProjectListOptions) (ProjectListResponse, error) {
	projects := ProjectListResponse{}
	err := p.c.get("/projects", opts.ToQuery(), &projects)
	return projects, err
}

func (p *Projects) Create(project *Project) error {
	err := p.c.post("/projects", project)
	return err
}

func (p *Projects) Update() {

}

func (p *Projects) Delete(projectId uint) error {
	return p.c.delete(fmt.Sprintf("/projects/%d", projectId))

}

func (p *Projects) Find(projectId uint) (*Project, error) {
	project := &Project{}
	err := p.c.get(fmt.Sprintf("/projects/%d", projectId), nil, project)
	return project, err
}

func (opts *ProjectListOptions) ToQuery() map[string]string {
	return map[string]string{
		"application_id": opts.ApplicationId,
	}
}
