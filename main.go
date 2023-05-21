package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handlers/routes"
)

func main() {

	r := gin.Default()
	routes.UserAuth(r)
	routes.WallpaperPage(r)
	routes.Images(r)
	routes.Profile(r)

	r.Run("192.168.43.236:8080")
}
