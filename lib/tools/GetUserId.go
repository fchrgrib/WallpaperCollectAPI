package tools

import (
	"errors"
	"github.com/config"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lib/middleware"
	"github.com/models"
	"net/http"
	"strings"
)

func GetUserIdFromCookies(c *gin.Context) (string, error) {
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
		userId = claims.Id
	}
	return userId, nil
}

func GetUserIdFromUserName(userName string) (string, error) {
	var user models.UserOtherEmailDesc

	db, err := database.ConnectDB()
	if err != nil {
		return "", err
	}

	if err := db.Table("user").Where("user_name = ?", userName).First(&user).Error; err != nil {
		return "", err
	}

	return user.Id.String(), nil
}
