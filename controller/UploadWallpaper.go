package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/lib/tools"
)

func UploadWallpaper(c *gin.Context) {
	var walpapper models.Wallpaper

	if err := c.ShouldBind(&walpapper); err != nil {
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

	if err := c.SaveUploadedFile(walpapper.Image, "././assets/"+id+"/"+walpapper.Image.Filename); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
