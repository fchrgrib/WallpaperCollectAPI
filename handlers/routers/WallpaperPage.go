package routers

import (
	"github.com/controllers/wallpaperpage"
	"github.com/gin-gonic/gin"
	"github.com/lib/middleware"
)

func WallpaperPage(routers *gin.Engine) {
	privateRouters := routers.Group("/wallpaper")
	privateRouters.Use(middleware.JWT)
	privateRouters.POST("/upload", wallpaperpage.UploadWallpaper)
	privateRouters.GET("", wallpaperpage.WallpaperCollection)
}
