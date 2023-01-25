package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/handler"
)

func CreateUserAuth(c *gin.Context) {
	var user models.User
	db, err := database.ConnectDB()

	if err != nil {
		panic(err)
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		panic(err)
		return
	}
	if err := handler.CreateUser(user); err != nil {
		panic(err)
		return
	}
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	db.Table("users").Create(&user)

}
