package middleware

import (
	"fmt"
	"github.com/config"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
)

func AuthWithToken(c *gin.Context) {
	var User config.Claims

	_token, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(_token, &User, ValidateAccessJWT)

	var userId string
	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		userId = claims.Id
	}

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("user").Where("id = ?", userId).Error; err == nil {
		c.Next()
		return
	}
	fmt.Println(err)

	c.AbortWithStatus(http.StatusUnauthorized)
}
