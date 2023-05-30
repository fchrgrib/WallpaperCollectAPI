package wallpaperpage

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/libs/utils/data"
	"github.com/models"
	"net/http"
)

func WallpaperCollection(c *gin.Context) {

	var wallpaperCollect []models.WallpaperCollectionDB

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"wallpaper_collection": "",
			"status":               err.Error(),
		})
		return
	}

	id, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"wallpaper_collection": "",
			"status":               err.Error(),
		})
		return
	}

	if err := db.Table("wallpaper_collect").Where("user_id = ?", id).Find(&wallpaperCollect).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"wallpaper_collection": "",
			"status":               err.Error(),
		})
		return
	}
	var imageUrl []string

	for _, value := range wallpaperCollect {
		imageUrl = append(imageUrl, "https://wallpapercollectapi-production.up.railway.app/images/"+value.ImageId)
	}
	c.JSON(http.StatusOK, gin.H{
		"wallpaper_collection": imageUrl,
		"status":               "ok",
	})

	return
}
