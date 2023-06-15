package wallpaperpage

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/libs/utils/data"
	"github.com/models"
	"net/http"
	"os"
)

func UploadWallpaper(c *gin.Context, router *gin.RouterGroup) {
	var wallpaper models.Wallpaper

	if err := c.ShouldBind(&wallpaper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "your uploading process is filed",
		})
		return
	}

	id, err := data.GetUserIdFromCookies(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": err.Error(),
		})
		return
	}
	uid := uuid.New().String()
	imageName := wallpaper.Image.Filename
	path := "././assets/" + id + "/wallpaper_collection/" + uid + "_" + imageName
	if err := c.SaveUploadedFile(wallpaper.Image, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := data.AllWallpaperToDB(id, path, uid, imageName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	fileStat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if fileStat.Size() != 0 {
		router.GET(uid, func(context *gin.Context) {
			c.File(path)
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
