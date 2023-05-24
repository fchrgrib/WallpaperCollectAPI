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
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	if err := db.Table("wallpaper_collect").Where("image_id = ?", c.Param("id")).First(&wallpaper).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	os.Remove(wallpaper.Path)

	if err := db.Table("wallpaper_collect").Delete(wallpaper).Error; err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

}
