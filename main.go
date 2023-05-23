package main

import (
	"github.com/gin-gonic/gin"
	"github.com/handlers/routers"
)

func main() {

	r := gin.Default()
	routers.UserAuth(r)
	routers.WallpaperPage(r)
	routers.Images(r)
	routers.Profile(r)

	r.Run()
}
