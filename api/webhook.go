package api

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Webhooks struct {
	c *Client
}

type Webhook struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Token        string         `json:"token,omitempty"`
	TagFilter    string         `json:"tag_filter,omitempty"`
	BranchFilter string         `json:"branch_filter,omitempty"`
	ProjectID    uint           `json:"project_id"`
	Project      Project        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"-"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	UpdatedAt    time.Time      `json:"updated_at,omitempty"`
	Deleted      gorm.DeletedAt `json:"-"`
}

type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

func (c *Client) Webhooks() *Webhooks {
	return &Webhooks{c}
}

func (w *Webhooks) List(projectId uint) (WebhookListResponse, error) {
	webhooks := WebhookListResponse{}
	err := w.c.get(fmt.Sprintf("/projects/%d/webhooks", projectId), nil, &webhooks)
	return webhooks, err
}

func (w *Webhooks) Create(projectId uint, webhook *Webhook) error {
	err := w.c.post(fmt.Sprintf("/projects/%d/webhooks", projectId), webhook)
	return err
}

func (w *Webhooks) Update(projectId uint, webhook *Webhook) error {
	return w.c.put(fmt.Sprintf("/projects/%d/webhooks/%d", projectId, webhook.ID), webhook)
}

func (w *Webhooks) Delete(projectId, webhookId uint) error {
	return w.c.delete(fmt.Sprintf("/projects/%d/webhooks/%d", projectId, webhookId))
}

func (w *Webhooks) Find(projectId uint, webhookId uint) (*Webhook, error) {
	webhook := &Webhook{}
	query := map[string]string{}
	err := w.c.get(fmt.Sprintf("/projects/%d/webhooks/%d", projectId, webhookId), query, webhook)
	return webhook, err
}
