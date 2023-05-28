package profile

import (
	"github.com/gin-gonic/gin"
	"github.com/libs/utils/data"
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
	c.JSON(http.StatusOK, gin.H{
		"user_name":     userData.UserName,
		"phone_number":  userData.PhoneNumber,
		"email":         userData.Email,
		"photo_profile": userData.PhotoProfile,
	})
	return
}
