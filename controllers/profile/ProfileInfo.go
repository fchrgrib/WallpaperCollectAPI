package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/utils/data"
	"net/http"
)

func ProfileInfo(c *gin.Context) {
	userId, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	userData, err := data.GetUserDataWithId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.File(userData.PhotoProfile)
	c.JSON(http.StatusOK, gin.H{
		"user_name":    userData.UserName,
		"phone_number": userData.PhoneNumber,
		"email":        userData.Email,
	})
	return
}
