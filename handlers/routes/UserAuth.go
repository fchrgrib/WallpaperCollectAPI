package routes

import (
	"github.com/controllers/logresg"
	"github.com/gin-gonic/gin"
)

func UserAuth(routers *gin.Engine) {
	routers.POST("/register", logresg.CreateUserAuth)
	routers.POST("/login", logresg.LoginController)
	routers.GET("/logout", logresg.Logout)
}
