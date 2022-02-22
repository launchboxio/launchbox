package server

import (
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
	var apps []api.Application
	database.Find(&apps)
	c.JSON(http.StatusOK, gin.H{"applications": apps})
}

func (p *Projects) Get(c *gin.Context) {
	id := c.Param("applicationId")
	var app api.Application
	database.First(&app, id)
	c.JSON(http.StatusOK, app)
}

func (p *Projects) Create(c *gin.Context) {
	app := api.Application{}
	err := c.ShouldBind(&app)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	database.Create(&app)
	c.JSON(http.StatusOK, app)
}

func (p *Projects) Update(c *gin.Context) {
	app := api.Application{}
	id := c.Param("applicationId")
	update := struct {
		Name string `json:"name"`
	}{}
	err := c.ShouldBind(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	database.First(&app, id)
	app.Name = update.Name
	database.Save(&app)
	c.JSON(http.StatusOK, gin.H{"data": app})
}

func (p *Projects) Delete(c *gin.Context) {
	database.Where("id = ?", c.Param("applicationId")).Delete(&api.Application{})
	c.JSON(http.StatusNoContent, gin.H{})
}
