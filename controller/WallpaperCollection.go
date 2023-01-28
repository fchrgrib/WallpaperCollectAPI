package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/lib/tools"
)

func WallpaperCollection(c *gin.Context) {
	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	id, err := tools.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	var paths []models.Paths
	_ = db.Table("wallpaper_collections").Where("user_id = ?", id).Find(&paths)

	c.JSON(http.StatusOK, gin.H{
		"user": paths,
	})
}
