package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Users struct {
	database *gorm.DB
}

func New(database *gorm.DB) *Users {
	return &Users{
		database: database,
	}
}

func (u *Users) Create(c *gin.Context) {
	// Accept a JWT from next-auth, and persist

}
