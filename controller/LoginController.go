package controller

import (
	"github.com/config"
	"github.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/handler"
	"net/http"
	"time"
)

func LoginController(c *gin.Context) {

	var userInput models.UserLogin

	if err := c.ShouldBindJSON(&userInput); err != nil {
		panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	userDB, err := handler.Login(userInput)
	if err != nil {
		panic(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	expTime := time.Now().Local().Add(time.Hour * 1)
	claims := &config.Claims{
		Id:       userDB.Id,
		UserName: userDB.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)

	if err != nil {
		panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
