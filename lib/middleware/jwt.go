package middleware

import (
	"fmt"
	"github.com/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func ValidateAccessJWT(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return config.JwtKey, nil
}
func JWT(c *gin.Context) {
	var User config.Claims

	tokenString := c.Request.Header.Get("Cookie")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "request does not contain an access token"})
		c.Abort()
		return
	}

	vals := strings.Split(tokenString, "=")

	token, err := jwt.ParseWithClaims(vals[1], &User, ValidateAccessJWT)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		c.Set("id", claims.Id)
		c.Set("email", claims.Email)
	}

	c.Next()
}
