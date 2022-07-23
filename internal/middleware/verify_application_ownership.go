package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/internal/resources/applications"
)

func VerifyApplicationOwnership(ctrl applications.Applications) gin.HandlerFunc {
	return func(c *gin.Context) {
		return
	}
}
