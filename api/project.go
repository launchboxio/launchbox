package api

import (
	"fmt"
	"gorm.io/gorm"
	"regexp"
	"strings"
	"time"
)

type Projects struct {
	c *Client
}

type ProjectStatus string

const (
	ProjectStatusCreating ProjectStatus = "creating"
	ProjectStatusUpdating ProjectStatus = "updating"
	ProjectStatusDeleting ProjectStatus = "deleting"
	ProjectStatusActive   ProjectStatus = "spring"
)

type Project struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	Repo          string         `json:"repo"`
	Branch        string         `json:"branch,omitempty"`
	Status        ProjectStatus  `json:"string"`
	ApplicationID uint           `json:"application_id"`
	Application   Application    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedAt     time.Time      `json:"created_at,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at,omitempty"`
	Deleted       gorm.DeletedAt `json:"deleted,omitempty"`
}

func (p *Project) GetFriendlyName() string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.ToLower(reg.ReplaceAllString(p.Name, "-"))
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
