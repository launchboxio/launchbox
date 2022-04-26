package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	vault "github.com/hashicorp/vault/api"
	"github.com/launchboxio/launchbox/api"
	"net/http"
)

type Secrets struct {
	Vault *vault.Client
}

func (s *Secrets) Register(r *gin.Engine) {

	// Manage secrets for an application
	applicationGroup := r.Group("/applications/:applicationId/secrets")
	applicationGroup.GET("", s.ListApplicationSecrets)
	//applicationGroup.GET("/:secretId", s.GetApplicationSecret)
	applicationGroup.POST("", s.CreateApplicationSecret)
	applicationGroup.PUT("/:secretId", s.UpdateApplicationSecret)
	applicationGroup.DELETE("/:secretId", s.DeleteApplicationSecret)

	// Manage secrets for a project
	projectGroup := r.Group("/projects/:projectId/secrets")
	projectGroup.GET("", s.ListProjectSecrets)
	projectGroup.POST("", s.CreateProjectSecret)
	//projectGroup.GET("/:secretId", s.GetProjectSecret)
	projectGroup.PUT("/:secretId", s.UpdateProjectSecret)
	projectGroup.DELETE("/:secretId", s.DeleteProjectSecret)
}

func (s *Secrets) ListApplicationSecrets(c *gin.Context) {
	var secrets []api.Secret
	database.Where("object_type = ? AND object_id = ?", "application", c.Param("applicationId")).Find(&secrets)
	c.JSON(http.StatusOK, gin.H{"secrets": &secrets})
}

func (s *Secrets) GetApplicationSecret(c *gin.Context) {

}

func (s *Secrets) CreateApplicationSecret(c *gin.Context) {
	secret := &api.Secret{}
	err := c.ShouldBind(&secret)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	secretId, err := uuid.NewV4()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	secret.ID = secretId
	secret.Path = getSecretPath(secret)

	err = s.syncSecret(secret)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	database.Create(&secret)

}

func (s *Secrets) UpdateApplicationSecret(c *gin.Context) {

}

func (s *Secrets) DeleteApplicationSecret(c *gin.Context) {

}

func (s *Secrets) ListProjectSecrets(c *gin.Context) {

}

func (s *Secrets) GetProjectSecret(c *gin.Context) {

}

func (s *Secrets) CreateProjectSecret(c *gin.Context) {

}

func (s *Secrets) UpdateProjectSecret(c *gin.Context) {

}

func (s *Secrets) DeleteProjectSecret(c *gin.Context) {

}

func (s *Secrets) syncSecret(secret *api.Secret) error {
	existingSecret, err := s.Vault.Logical().Read(secret.Path)

	fmt.Println(existingSecret)
	fmt.Println(err)
	// Handle not found?
	if err != nil {
		return err
	}

	var inputData map[string]interface{}
	if existingSecret != nil {
		existingData := existingSecret.Data
		existingData[secret.Name] = secret.Value

		inputData = map[string]interface{}{
			"data": existingData,
		}
	} else {
		inputData = map[string]interface{}{
			"data": map[string]interface{}{
				secret.Name: secret.Value,
			},
		}
	}

	_, err = s.Vault.Logical().Write(secret.Path, inputData)

	if err != nil {
		return err
	}

	// TODO: Store the version
	return nil
}

func getSecretPath(s *api.Secret) string {
	return fmt.Sprintf("secret/data/%s/%s", s.ObjectType, s.ObjectId)
}
