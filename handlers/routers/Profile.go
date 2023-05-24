package routers

import (
	"github.com/controllers/profile"
	"github.com/gin-gonic/gin"
	"github.com/libs/middleware"
)

func Profile(routers *gin.Engine) {
	privateRouters := routers.Group("/wallpaper")
	privateRouters.Use(middleware.JWT)
	profileRouter := privateRouters.Group("/profile")
	profileRouter.GET("", profile.ProfileInfo)
	profileRouter.PUT("/update_profile", profile.UpdateProfileDescription)
	profileRouter.PUT("/upload_profile_picture", profile.PhotoProfileUpload)
}
