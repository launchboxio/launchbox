package server

import (
	"github.com/gin-gonic/gin"
	"github.com/robwittman/launchbox/api"
	"net/http"
)

type Applications struct {
}

func (a *Applications) Register(r *gin.Engine) {
	group := r.Group("/applications")
	group.GET("", a.List)
	group.GET("/:applicationId", a.Get)
	group.POST("", a.Create)
	group.PUT("/:applicationId", a.Update)
	group.DELETE("/:applicationId", a.Delete)
}

func (a *Applications) List(c *gin.Context) {
	var apps []api.Application
	database.Find(&apps)
	c.JSON(http.StatusOK, gin.H{"applications": apps})
}

func (a *Applications) Get(c *gin.Context) {
	id := c.Param("applicationId")
	var app api.Application
	database.First(&app, id)
	c.JSON(http.StatusOK, app)
}

func (a *Applications) Create(c *gin.Context) {
	app := api.Application{}
	err := c.ShouldBind(&app)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	database.Create(&app)
	c.JSON(http.StatusOK, app)
}

func (a *Applications) Update(c *gin.Context) {
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

func (a *Applications) Delete(c *gin.Context) {
	database.Where("id = ?", c.Param("applicationId")).Delete(&api.Application{})
	c.JSON(http.StatusNoContent, gin.H{})
}
