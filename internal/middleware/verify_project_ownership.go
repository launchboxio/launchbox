package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/internal/resources/projects"
)

func VerifyProjectOwnership(ctrl projects.Projects) gin.HandlerFunc {
	return func(c *gin.Context) {
		return
	}
}
