package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/tbxark/g4vercel"
	"walpapperCollectRestAPI/controller"
	"walpapperCollectRestAPI/lib/middleware"
)

func main() {

	r := gin.Default()
	r.POST("/register", controller.CreateUserAuth)
	r.POST("/login", controller.LoginController)
	r.GET("/logout", controller.Logout)

	privateRouters := r.Group("/wallpaper")
	privateRouters.Use(middleware.JWT)
	privateRouters.PUT("/upload", controller.UploadWallpaper)
	privateRouters.GET("", controller.WallpaperCollection)

	profileRouter := privateRouters.Group("/profile")
	profileRouter.GET("", controller.ProfileInfo)
	profileRouter.PUT("/update_profile", controller.UpdateProfileDescription)
	profileRouter.PUT("/upload_profile_picture", controller.PhotoProfileUpload)

	r.Run()
}
