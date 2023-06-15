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
	profileRouter.GET("", profile.Info)
	profileRouter.PUT("/update_profile_desc", profile.UpdateProfileDescription)

	rProfile := routers.Group("photo_profile")

	profileRouter.PUT("/update_profile_picture", func(c *gin.Context) {
		profile.UpdatePhotoProfile(c, rProfile)
	})
	profileRouter.POST("/upload_profile_picture", func(c *gin.Context) {
		profile.PhotoProfileUpload(c, rProfile)
	})
}
