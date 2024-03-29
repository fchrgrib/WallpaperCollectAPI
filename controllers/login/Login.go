package login

import (
	"github.com/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/handlers/authandlers"
	"github.com/libs/utils/data"
	"github.com/models"
	"net/http"
	"time"
)

func EmailLoginDefaultController(c *gin.Context) {

	var userInput models.UserLogin

	if err := c.ShouldBindJSON(&userInput); err != nil {
		panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	userDB, err := authandlers.Login(userInput)
	if err != nil {
		panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	var value string

	if value, err = data.GetUserIdFromEmail(userDB.Email); err != nil {
		panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	expTime := time.Now().Local().Add(time.Hour * 3)
	claims := &config.Claims{
		Id:    value,
		Email: userDB.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JwtKey)

	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.SetCookie("token", token, 4*3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
