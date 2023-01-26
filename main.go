package main

import (
	"github.com/gin-gonic/gin"
	"walpapperCollectRestAPI/controller"
)

func main() {
	routs := gin.Default()

	routs.POST("/register", controller.CreateUserAuth)
	routs.POST("/login", controller.LoginController)
	routs.Run()
}
