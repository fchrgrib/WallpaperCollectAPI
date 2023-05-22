package profile

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/lib/tools"
	"github.com/models"
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

	userId, err := tools.GetUserIdFromCookies(c)
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

	if err := db.Table("user").Where("user_name = ?", &userUpdate.UserName).First(&models.UserOtherEmailDesc{}); err.Error == nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "user name already exist",
		})
		return
	}

	if !tools.ValidateEmail(userUpdate.Email) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "invalid email",
		})
		return
	}

	if !tools.ValidationNumberPhone(userUpdate.PhoneNumber) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"status": "invalid phone number",
		})
		return
	}
	t := time.Now().Local()
	userUpdate.UpdatedAt = &t
	user = userUpdate
	db.Save(user)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
