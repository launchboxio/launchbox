package applications

import (
	haikunator "github.com/atrox/haikunatorgo/v2"
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/api"
	"gorm.io/gorm"
	"net/http"
)

type Applications struct {
	database *gorm.DB
}

func New(database *gorm.DB) *Applications {
	return &Applications{database: database}
}

func (a *Applications) List(c *gin.Context) {
	var apps []api.Application
	a.database.Find(&apps)
	c.JSON(http.StatusOK, gin.H{"applications": &apps})
}

func (a *Applications) Get(c *gin.Context) {
	id := c.Param("applicationId")
	var app api.Application
	if c.Query("deleted") != "" {
		a.database.Unscoped().First(&app, id)
	} else {
		a.database.First(&app, id)
	}
	c.JSON(http.StatusOK, &app)
}

func (a *Applications) Create(c *gin.Context) {
	haiku := haikunator.New()
	app := &api.Application{}
	err := c.ShouldBind(&app)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	app.Namespace = haiku.Haikunate()
	a.database.Create(&app)

	//_, err = createNamespaceTask(app.ID)
	//if err != nil {
	//	fmt.Println(err)
	//}
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
	a.database.First(&app, id)
	app.Name = update.Name
	a.database.Save(&app)
	c.JSON(http.StatusOK, app)
}

func (a *Applications) Delete(c *gin.Context) {
	applicationId := c.Param("applicationId")
	a.database.Where("id = ?", applicationId).Delete(&api.Application{})
	//id, _ := strconv.Atoi(applicationId)
	//_, err := deleteNamespaceTask(uint(id))
	//if err != nil {
	//	fmt.Println(err)
	//}
	c.JSON(http.StatusNoContent, gin.H{})
}
