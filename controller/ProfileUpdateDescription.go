package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/lib/tools"
)

func UpdateProfileDescription(c *gin.Context) {

	var user models.User
	var userUpdate models.UserUpdate

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	userId, err := tools.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("users").Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	db.Model(&user).Updates(userUpdate)
}
