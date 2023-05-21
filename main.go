package main

import (
	"github.com/controller"
	"github.com/database"
	"github.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/lib/middleware"
)

func main() {
	var wallpaper []models.WallpaperCollection

	db, err := database.ConnectDB()
	if err != nil {
		return
	}

	db.Table("wallpaper_collect").Find(&wallpaper)

	r := gin.Default()
	r.POST("/register", controller.CreateUserAuth)
	r.POST("/login", controller.LoginController)
	r.GET("/logout", controller.Logout)

	privateRouters := r.Group("/wallpaper")

	for _, values := range wallpaper {
		privateRouters.Static(values.ImageId.String(), values.Path)
	}

	privateRouters.Use(middleware.JWT)
	privateRouters.POST("/upload", controller.UploadWallpaper)
	privateRouters.GET("", controller.WallpaperCollection)

	profileRouter := privateRouters.Group("/profile")
	profileRouter.GET("", controller.ProfileInfo)
	profileRouter.PUT("/update_profile", controller.UpdateProfileDescription)
	profileRouter.PUT("/upload_profile_picture", controller.PhotoProfileUpload)

	r.Run()
}
