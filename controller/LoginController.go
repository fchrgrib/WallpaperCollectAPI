package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
	"walpapperCollectRestAPI/config"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/handler"
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

	expTime := time.Now().Add(time.Hour * 1)
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

	//http.SetCookie(gin.Context{}.Writer, &http.Cookie{
	//	Name:     "token",
	//	Path:     "/",
	//	Value:    token,
	//	HttpOnly: true,
	//})
	c.SetCookie("token", token, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
