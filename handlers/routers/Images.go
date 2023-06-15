package routers

import (
	"github.com/controllers/wallpaperpage"
	"github.com/gin-gonic/gin"
	"github.com/libs/middleware"
)

func Images(routers *gin.Engine) {

	rImage := routers.Group("/images")
	rImage.Use(func(c *gin.Context) {
		middleware.AuthWithToken(c)
	})
	rImage.GET("/:id/download", wallpaperpage.DownloadWallpaper)
	rImage.DELETE("/:id/delete", wallpaperpage.DeleteWallpaperController)

}
