package login

import "github.com/gin-gonic/gin"

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	return
}
