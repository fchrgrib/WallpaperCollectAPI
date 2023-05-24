package routers

import (
	"github.com/controllers/login"
	"github.com/controllers/register"
	"github.com/gin-gonic/gin"
)

func UserAuth(routers *gin.Engine) {
	routers.POST("/register-email-default", register.CreateUserAuth)
	routers.POST("/login-email-default", login.EmailLoginDefaultController)

	routers.GET("/register-google-session", register.RedirectGoogleRegisterController)
	routers.GET("/oauth-google", register.CreateUserAuthFacebook)

	routers.GET("/register-facebook-session", register.RedirectFacebookRegisterController)
	routers.GET("/oauth-facebook", register.CreateUserAuthGoogle)

	routers.GET("/login-google-session", login.RedirectGoogleLoginController)
	routers.GET("/login-google", login.EmailLoginGoogleController)

	routers.GET("/login-facebook-session", login.RedirectFacebookLoginController)
	routers.GET("/login-facebook", login.EmailLoginFacebookController)

	routers.GET("/logout", login.Logout)
}
