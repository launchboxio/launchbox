package api

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID            string `gorm:"primaryKey"`
	Email         string
	EmailVerified time.Time
	Name          string
	AvatarUrl     string `gorm:"column:image"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Deleted       gorm.DeletedAt
}

type Account struct {
	ID                string
	UserID            string
	Type              string
	Provider          string
	ProviderAccountId string
	RefreshToken      string
	AccessToken       string
	ExpiresAt         int
	TokenType         string
	Scope             string
	IdToken           string
	SessionState      string
	OauthTokenSecret  string
	OauthToken        string
}

type Session struct {
	ID           string `gorm:"primaryKey"`
	Expires      time.Time
	SessionToken string `gorm:"column:sessionToken"`
	UserId       string `gorm:"column:userId"`
}

type VerificationToken struct {
	Identifier string
	Token      string
	Expires    time.Time
}
