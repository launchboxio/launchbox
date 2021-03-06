package revisions

import (
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/api"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Revisions struct {
	database *gorm.DB
}

func New(database *gorm.DB) *Revisions {
	return &Revisions{
		database: database,
	}
}

func (rev *Revisions) List(c *gin.Context) {
	var revisions []api.Revision
	rev.database.Where("project_Id = ?", c.Param("projectId")).Find(&revisions)
	c.JSON(http.StatusOK, gin.H{"revisions": &revisions})
}

func (rev *Revisions) Get(c *gin.Context) {
	id := c.Param("revisionId")
	var revision api.Revision
	rev.database.First(&revision, id)
	c.JSON(http.StatusOK, revision)
}

func (rev *Revisions) Create(c *gin.Context) {
	revision := api.Revision{}
	err := c.ShouldBind(&revision)
	projectId, _ := strconv.ParseUint(c.Param("projectId"), 10, 0)
	revision.ProjectID = uint(projectId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
	}
	revision.Status = api.RevisionStatusDeploying

	err = rev.database.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&revision).Error; err != nil {
			return err
		}

		if err = tx.Where("id = ?", revision.ProjectID).Update("status", api.ProjectStatusUpdating).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Faied creating revision",
		})
	} else {
		c.JSON(http.StatusOK, revision)
	}
}

func (rev *Revisions) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

func (rev *Revisions) Delete(c *gin.Context) {
	rev.database.Where("id = ?", c.Param("revisionId")).Delete(&api.Revision{})
	c.JSON(http.StatusNoContent, gin.H{})
}
