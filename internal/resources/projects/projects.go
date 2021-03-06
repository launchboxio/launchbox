package projects

import (
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/api"
	"gorm.io/gorm"
	"net/http"
)

type Projects struct {
	database *gorm.DB
}

func New(database *gorm.DB) *Projects {
	return &Projects{database: database}
}

func (p *Projects) List(c *gin.Context) {
	var projects []api.Project
	applicationId := c.Query("application_id")
	p.database.Where("application_id = ?", applicationId).Find(&projects)
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (p *Projects) Get(c *gin.Context) {
	id := c.Param("projectId")
	var project api.Project
	p.database.First(&project, id)
	c.JSON(http.StatusOK, project)
}

func (p *Projects) Create(c *gin.Context) {
	project := &api.Project{}
	err := c.ShouldBind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	p.database.Create(&project)
	//_, err = createServiceTask(project.ApplicationID, project.ID)
	//if err != nil {
	//	fmt.Println(err)
	//}
	c.JSON(http.StatusOK, project)
}

func (p *Projects) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (p *Projects) Delete(c *gin.Context) {
	p.database.Where("id = ?", c.Param("projectId")).Delete(&api.Project{})
	c.JSON(http.StatusNoContent, gin.H{})
}
