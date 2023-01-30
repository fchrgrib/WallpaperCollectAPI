package main

import (
	"github.com/gin-gonic/gin"
	"walpapperCollectRestAPI/controller"
	"walpapperCollectRestAPI/lib/middleware"
)

func main() {
	//var fs http.FileSystem

	r := gin.Default()
	r.POST("/register", controller.CreateUserAuth)
	r.POST("/login", controller.LoginController)
	r.GET("logout", controller.Logout)

	privateRouters := r.Group("/wallpaper")
	privateRouters.Use(middleware.JWT)
	privateRouters.PUT("/upload", controller.UploadWallpaper)
	privateRouters.GET("/your_collection", controller.WallpaperCollection)
	privateRouters.PUT("/update_profile", controller.UpdateProfileDescription)

	r.Run()
}
