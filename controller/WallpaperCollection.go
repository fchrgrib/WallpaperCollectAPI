package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"walpapperCollectRestAPI/lib/tools"
)

func WallpaperCollection(c *gin.Context) {

	id, err := tools.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	dir := http.Dir("././assets/" + id)
	c.FileFromFS("", dir)
}
