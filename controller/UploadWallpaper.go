package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/handler"
	"walpapperCollectRestAPI/lib/tools"
)

func UploadWallpaper(c *gin.Context) {
	var wallpaper models.Wallpaper

	if err := c.ShouldBind(&wallpaper); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "your uploading process is filed",
		})
		return
	}

	id, err := tools.GetUserId(c)

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

	if err := handler.AllWallpaperToDB(id, path, uid, imageName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
