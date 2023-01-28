package tools

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"walpapperCollectRestAPI/config"
	"walpapperCollectRestAPI/lib/middleware"
)

func GetUserId(c *gin.Context) (string, error) {
	var User config.Claims

	tokesString := c.Request.Header.Get("Cookie")

	if tokesString == "" {
		return "", errors.New("your token is invalid")
	}

	code := strings.Split(tokesString, "=")

	token, err := jwt.ParseWithClaims(code[1], &User, middleware.ValidateAccessJWT)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": err.Error(),
		})
		return "", err
	}
	var userId string
	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		c.Set("id", claims.Id)
		c.Set("user_name", claims.UserName)
		userId = claims.Id.String()
	}
	return userId, nil
}
