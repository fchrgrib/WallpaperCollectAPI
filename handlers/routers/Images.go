package routers

import (
	"github.com/controllers/wallpaperpage"
	"github.com/gin-gonic/gin"
	"github.com/libs/middleware"
)

func Images(routers *gin.Engine) {

	rImage := routers.Group("/images")
	rImage.Use(middleware.AuthWithToken)
	rImage.DELETE("/:id/delete", wallpaperpage.DeleteWallpaperController)

}
