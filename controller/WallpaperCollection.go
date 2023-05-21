package controller

import (
	"github.com/database"
	"github.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/tools"
	"net/http"
)

func WallpaperCollection(c *gin.Context) {

	var wallpaperCollect []models.WallpaperCollection

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	id, err := tools.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	UserData, err := tools.GetUserDataWithId(id)

	if err := db.Table("wallpaper_collect").Where("user_id = ?", id).Find(&wallpaperCollect).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	var imageUrl []string

	for _, value := range wallpaperCollect {
		imageUrl = append(imageUrl, "http://https://wallpapercollectapi-production.up.railway.app/wallpaper/"+value.ImageId.String())
	}
	c.JSON(http.StatusOK, gin.H{
		"user_name":            UserData.UserName,
		"phone_number":         UserData.PhoneNumber,
		"email":                UserData.Email,
		"wallpaper_collection": imageUrl,
	})

	return
}
