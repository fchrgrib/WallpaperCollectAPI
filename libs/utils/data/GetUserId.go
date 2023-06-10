package data

import (
	"errors"
	"github.com/config"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/libs/middleware"
	"github.com/models"
	"net/http"
)

func GetUserIdFromCookies(c *gin.Context) (string, error) {
	var User config.Claims

	tokesString, err := c.Cookie("token")

	if err != nil {
		return "", errors.New("your token is invalid")
	}

	token, err := jwt.ParseWithClaims(tokesString, &User, middleware.ValidateAccessJWT)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": err.Error(),
		})
		return "", err
	}
	var userId string
	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		c.Set("id", claims.Id)
		c.Set("email", claims.Email)
		userId = claims.Id
	}
	return userId, nil
}

func GetUserIdFromEmail(email string) (string, error) {
	var user models.UserDescDB

	db, err := database.ConnectDB()
	if err != nil {
		return "", err
	}

	if err := db.Table("user").Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}

	return user.Id, nil
}
