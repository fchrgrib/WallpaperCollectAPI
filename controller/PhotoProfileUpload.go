package controller

import (
	"github.com/database"
	"github.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/tools"
	"net/http"
)

func PhotoProfileUpload(c *gin.Context) {
	var ppUpload models.PhotoProfile
	var user models.User

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&ppUpload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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

	uid := uuid.New().String()
	imageName := ppUpload.Image.Filename
	path := "././assets/" + userId + "/profile/" + uid + "_" + imageName

	if err := c.SaveUploadedFile(ppUpload.Image, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	userData, err := tools.GetUserDataWithId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	userData.PhotoProfile = path

	if err := db.Model(&user).Updates(userData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
}