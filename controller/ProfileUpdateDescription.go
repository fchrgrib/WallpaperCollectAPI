package controller

import (
	"github.com/database"
	"github.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/tools"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func UpdateProfileDescription(c *gin.Context) {

	var user models.UserOtherEmailDesc
	var userUpdate models.UserOtherEmailDesc

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

	user, err = tools.GetUserDataWithId(userId)
	if err != nil {
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

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(userUpdate.Password), bcrypt.DefaultCost)

	userUpdate.UpdatedAt = time.Now().Local().String()
	userUpdate.Password = string(hashPass)

	db.Table("users").Model(&user).Updates(userUpdate)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
