package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robwittman/launchbox/api"
	"net/http"
)

type Projects struct {
}

func (p *Projects) Register(r *gin.Engine) {
	group := r.Group("/projects")
	group.GET("", p.List)
	group.GET("/:projectId", p.Get)
	group.POST("", p.Create)
	group.PUT("/:projectId", p.Update)
	group.DELETE("/:projectId", p.Delete)
}

func (p *Projects) List(c *gin.Context) {
	var projects []api.Project
	applicationId := c.Query("application_id")
	database.Where("application_id = ?", applicationId).Find(&projects)
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (p *Projects) Get(c *gin.Context) {
	id := c.Param("projectId")
	var project api.Project
	database.First(&project, id)
	c.JSON(http.StatusOK, project)
}

func (p *Projects) Create(c *gin.Context) {
	project := &api.Project{}
	err := c.ShouldBind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	database.Create(&project)
	_, err = createServiceTask(project.ApplicationID, project.ID)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, project)
}

func (p *Projects) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (p *Projects) Delete(c *gin.Context) {
	database.Where("id = ?", c.Param("projectId")).Delete(&api.Project{})
	c.JSON(http.StatusNoContent, gin.H{})
}
