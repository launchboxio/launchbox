package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/launchboxio/launchbox/internal/resources/revisions"
)

func VerifyRevisionOwnership(ctrl revisions.Revisions) gin.HandlerFunc {
	return func(c *gin.Context) {
		return
	}
}
