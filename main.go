package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/controller"
	"walpapperCollectRestAPI/lib/middleware"
)

func main() {
	routs := gin.Default()

	routs.POST("/register", controller.CreateUserAuth)
	routs.POST("/login", controller.LoginController)

	privateRouts := routs.Group("/private")
	privateRouts.Use(middleware.JwtTokenCheck)
	privateRouts.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ini private lohhh",
		})
	})

	routs.Run()
}
