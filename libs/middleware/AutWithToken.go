package middleware

import (
	"github.com/config"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/models"
	"net/http"
)

func AuthWithToken(c *gin.Context) {
	var (
		User          config.Claims
		justCheckUser models.UserDescDB
	)

	_token, err := c.Cookie("token")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		println("Salah disini tot tapi gatau kenapa ini tokennya \n" + _token + "\n ini errornya " + err.Error())
		return
	}

	token, err := jwt.ParseWithClaims(_token, &User, ValidateAccessJWT)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

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

	if _ = db.Table("user").Where("id = ?", userId).First(&justCheckUser); justCheckUser.Id != "" {
		c.Next()
		return
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}
