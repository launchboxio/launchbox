package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robwittman/launchbox/api"
	"net/http"
	"strconv"
)

type Revisions struct {
}

func (rev *Revisions) Register(r *gin.Engine) {
	group := r.Group("/projects/:projectId/revisions")
	group.GET("", rev.List)
	group.GET("/:revisionId", rev.Get)
	group.POST("", rev.Create)
	group.PUT("/:revisionId", rev.Update)
	group.DELETE("/:revisionId", rev.Delete)
}

func (rev *Revisions) List(c *gin.Context) {
	var revisions []api.Revision
	database.Where("project_Id = ?", c.Param("projectId")).Find(&revisions)
	c.JSON(http.StatusOK, gin.H{"revisions": &revisions})
}

func (rev *Revisions) Get(c *gin.Context) {
	id := c.Param("revisionId")
	var revision api.Revision
	database.First(&revision, id)
	c.JSON(http.StatusOK, revision)
}

func (rev *Revisions) Create(c *gin.Context) {
	revision := api.Revision{}
	err := c.ShouldBind(&revision)
	projectId, _ := strconv.ParseUint(c.Param("projectId"), 10, 0)
	revision.ProjectID = uint(projectId)
	fmt.Println(revision)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	database.Create(&revision)
	c.JSON(http.StatusOK, revision)
}

func (rev *Revisions) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (rev *Revisions) Delete(c *gin.Context) {
	database.Where("id = ?", c.Param("revisionId")).Delete(&api.Revision{})
	c.JSON(http.StatusNoContent, gin.H{})
}
