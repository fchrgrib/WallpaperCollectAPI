package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/controller"
	"walpapperCollectRestAPI/lib/middleware"
)

func main() {
	r := gin.Default()
	r.POST("/register", controller.CreateUserAuth)
	r.POST("/login", controller.LoginController)
	r.GET("logout", controller.Logout)

	privateRouts := r.Group("/wallpaper")
	privateRouts.Use(middleware.JWT)
	privateRouts.PUT("/upload", controller.UploadWallpaper)
	privateRouts.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ini private lohhh",
		})
	})

	r.Run()
}
