package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/launchboxio/launchbox/api"
	"net/http"
	"strconv"
)

type Webhooks struct {
}

type WebhookParams struct {
	TagFilter    string `json:"tag_filter"`
	BranchFilter string `json:"branch_filter"`
}

func (w *Webhooks) Register(r *gin.Engine) {
	group := r.Group("/projects/:projectId/webhooks")
	group.GET("", w.List)
	group.GET("/:webhookId", w.Get)
	group.POST("", w.Create)
	group.POST("/:webhookToken", w.Receive)
	group.PUT("/:webhookId", w.Update)
	group.DELETE("/:webhookId", w.Delete)
}

func (w *Webhooks) List(c *gin.Context) {
	var webhooks []api.Webhook
	database.Where("project_id = ?", c.Param("projectId")).Find(&webhooks)
	c.JSON(http.StatusOK, gin.H{"webhooks": &webhooks})
}

func (w *Webhooks) Get(c *gin.Context) {
	id := c.Param("webhookId")
	var webhook api.Webhook
	database.First(&webhook, id)
	c.JSON(http.StatusOK, &webhook)
}

func (w *Webhooks) Create(c *gin.Context) {
	var postParams WebhookParams
	err := c.ShouldBind(&postParams)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
		return
	}

	projectId, _ := strconv.ParseUint(c.Param("projectId"), 10, 0)

	token := uuid.New()
	webhook := &api.Webhook{}
	webhook.BranchFilter = postParams.BranchFilter
	webhook.TagFilter = postParams.TagFilter
	webhook.Token = token.String()
	webhook.ProjectID = uint(projectId)

	if err = database.Create(webhook).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, webhook)
}

func (w *Webhooks) Receive(c *gin.Context) {
	projectId, _ := strconv.ParseUint(c.Param("projectId"), 10, 0)
	tokenId := c.Param("webhookToken")
	fmt.Println("Received request webhook %s for project %d", tokenId, uint(projectId))
}

func (w *Webhooks) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updating webhook"})
}

func (w *Webhooks) Delete(c *gin.Context) {
	webhookId := c.Param("webhookId")
	database.Where("id = ?", webhookId).Delete(&api.Webhook{})
	c.JSON(http.StatusNoContent, gin.H{})
}
