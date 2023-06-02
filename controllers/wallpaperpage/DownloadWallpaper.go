package wallpaperpage

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/models"
	"net/http"
)

func DownloadWallpaper(c *gin.Context) {

	var wallpaperUser models.WallpaperCollectionDB

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("wallpaper_collect").Where("image_id = ?", c.Param("id")).First(&wallpaperUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.File(wallpaperUser.Path)
	return
}
