package profile

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/libs/utils/data"
	"github.com/libs/utils/validation"
	"github.com/models"
	"net/http"
	"time"
)

func UpdateProfileDescription(c *gin.Context) {

	var user models.UserOtherEmailDescDB
	var userUpdate models.UserOtherEmailDescDB

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	userId, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("user").Where("user_id = ?", userId).First(&user).Error; err != nil {
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

	if !validation.ValidateEmail(userUpdate.Email) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "invalid email",
		})
		return
	}

	if !validation.ValidationNumberPhone(userUpdate.PhoneNumber) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "invalid phone number",
		})
		return
	}

	t := time.Now().Local()
	user.UserName = userUpdate.UserName
	user.Email = userUpdate.Email
	user.PhoneNumber = userUpdate.PhoneNumber
	user.UpdatedAt = &t
	if err := db.Save(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
