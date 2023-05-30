package wallpaperpage

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"net/http"
	"os"
)

func DeleteWallpaperController(c *gin.Context) {
	var wallpaper models.WallpaperCollectionDB

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}

	if err := db.Table("wallpaper_collect").Where("image_id = ?", c.Param("id")).First(&wallpaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}

	if err := os.Remove(wallpaper.Path); err != nil {
		return
	}

	if err := db.Table("wallpaper_collect").Delete(wallpaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}

}
