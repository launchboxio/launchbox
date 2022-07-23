package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	corsAllowedDomain = "http://localhost:4040"
	authHeader        = "Authorization"
	ctxTokenKey       = "Auth0Token"
	permClaim         = "permissions"
)

func VerifyAuthentication() gin.HandlerFunc {
	// TODO: Build a client from configured Auth providers

	return func(c *gin.Context) {
		token, err := extractToken(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		}
		c.Set("token", token)
	}
}

// TODO: Verify contents of the JWT issuer claims, as well as the that encryption matches public key
func extractToken(req *http.Request) (*jwt.Token, error) {
	authorization := req.Header.Get(authHeader)
	if authorization == "" {
		return nil, errors.New("authorization header missing")
	}
	bearerAndToken := strings.Split(authorization, " ")
	if len(bearerAndToken) < 2 {
		return nil, errors.New("malformed authorization header: " + authorization)
	}
	token, err := jwt.Parse(bearerAndToken[1], func(*jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}
