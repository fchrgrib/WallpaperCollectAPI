package routers

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/lib/middleware"
	"github.com/models"
)

func Images(routers *gin.Engine) {
	var wallpaper []models.WallpaperCollection

	db, err := database.ConnectDB()
	if err != nil {
		return
	}

	db.Table("wallpaper_collect").Find(&wallpaper)

	rImage := routers.Group("/images")
	rImage.Use(middleware.AuthWithToken)

	if len(wallpaper) != 0 {
		for _, values := range wallpaper {
			rImage.Static(values.ImageId.String(), values.Path)
		}
	}
}
