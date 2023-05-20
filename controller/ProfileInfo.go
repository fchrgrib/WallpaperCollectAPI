package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/tools"
	"net/http"
)

func ProfileInfo(c *gin.Context) {
	userId, err := tools.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	userData, err := tools.GetUserDataWithId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.File(userData.PhotoProfile)
	c.JSON(http.StatusOK, gin.H{
		"user_name":    userData.UserName,
		"description":  userData.Description,
		"phone_number": userData.PhoneNumber,
		"email":        userData.Email,
	})
	return
}
