package wallpaperpage

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/libs/utils/data"
	"github.com/models"
	"net/http"
)

func WallpaperCollection(c *gin.Context) {

	var (
		wallpaperCollect []models.WallpaperCollectionDB
		wallpaperStatus  []models.Images
	)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"wallpaper_collection": wallpaperStatus,
			"status":               err.Error(),
		})
		return
	}

	id, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"wallpaper_collection": wallpaperStatus,
			"status":               err.Error(),
		})
		return
	}

	if err := db.Table("wallpaper_collect").Where("user_id = ?", id).Order("created_at").Find(&wallpaperCollect).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"wallpaper_collection": wallpaperStatus,
			"status":               err.Error(),
		})
		return
	}

	for _, value := range wallpaperCollect {
		wallpaperStatus = append(
			wallpaperStatus, models.Images{
				ImageUrl:  "https://wallpapercollectapi-production-c728.up.railway.app/images/" + value.ImageId + "/",
				ImageId:   value.ImageId,
				ImageName: value.ImageName,
			})
	}
	c.JSON(http.StatusOK, gin.H{
		"wallpaper_collection": wallpaperStatus,
		"status":               "ok",
	})

	return
}
