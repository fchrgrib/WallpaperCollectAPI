package routers

import (
	"github.com/controllers/wallpaperpage"
	"github.com/gin-gonic/gin"
	"github.com/libs/middleware"
)

func WallpaperPage(routers *gin.Engine) {
	privateRouters := routers.Group("/wallpaper")
	privateRouters.Use(middleware.JWT)
	privateRouters.GET("", wallpaperpage.WallpaperCollection)
	privateRouters.POST("/upload", func(c *gin.Context) {
		wallpaperpage.UploadWallpaper(c, routers)
	})
}
