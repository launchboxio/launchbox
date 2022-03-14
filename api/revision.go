package api

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Revisions struct {
	c *Client
}

type RevisionStatus string

const (
	RevisionStatusDeploying  RevisionStatus = "deploying"
	RevisionStatusSucceded   RevisionStatus = "succeeded"
	RevisionStatusRolledBack RevisionStatus = "rolled_back"
	RevisionStatusSuperceded RevisionStatus = "superceded"
)

type Revision struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Status    RevisionStatus `json:"status,omitempty"`
	CommitSha string         `json:"commit_sha,omitempty"`
	ProjectID uint           `json:"project_id"`
	Project   Project        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedAt time.Time      `json:"created_at,omitempty"`
	UpdatedAt time.Time      `json:"updated_at,omitempty"`
	Deleted   gorm.DeletedAt `json:"deleted,omitempty"`
}

type RevisionListResponse struct {
	Revisions []Revision `json:"revisions"`
}

func (c *Client) Revisions() *Revisions {
	return &Revisions{c}
}

func (r *Revisions) List(projectId uint) (RevisionListResponse, error) {
	revisions := RevisionListResponse{}
	err := r.c.get(fmt.Sprintf("/projects/%d/revisions", projectId), nil, &revisions)
	return revisions, err
}

func (r *Revisions) Create(projectId uint, revision *Revision) error {
	err := r.c.post(fmt.Sprintf("/projects/%d/revisions", projectId), revision)
	return err
}

func (r *Revisions) Update() {

}

func (r *Revisions) Delete() {

}

func (r *Revisions) Find(projectId uint, revisionId uint) (*Revision, error) {
	revision := &Revision{}
	err := r.c.get(fmt.Sprintf("/projects/%d/revisions/%d", projectId, revisionId), nil, revision)
	return revision, err
}
