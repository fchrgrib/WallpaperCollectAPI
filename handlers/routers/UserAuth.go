package routers

import (
	"github.com/controllers/login"
	"github.com/controllers/register"
	"github.com/gin-gonic/gin"
	"github.com/lib/utils/oauth2utility"
)

func UserAuth(routers *gin.Engine) {
	routers.POST("/register-email-default", register.CreateUserAuth)
	routers.POST("/login-email-default", login.EmailLoginDefaultController)

	routers.GET("/login-google-session", login.EmailGoogleLoginController)
	routers.GET("/oauth-google", oauth2utility.AuthGoogleHandler)

	routers.GET("/logout", login.Logout)
}
