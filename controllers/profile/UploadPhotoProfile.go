package profile

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/libs/middleware"
	"github.com/libs/utils/data"
	models2 "github.com/models"
	"net/http"
	"os"
)

func PhotoProfileUpload(c *gin.Context, router *gin.Engine) {

	var (
		ppUpload       models2.PhotoProfile
		user           models2.UserDescDB
		photoProfileDB models2.UserPhotoProfileDB
	)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := c.ShouldBind(&ppUpload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}
	userId, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	uid := uuid.New().String()
	imageName := ppUpload.Image.Filename
	path := "././assets/" + userId + "/profile/" + uid + "_" + imageName

	if err := c.SaveUploadedFile(ppUpload.Image, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("user").Where("id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	user.PhotoProfile = "https://wallpapercollectapi-production-c728.up.railway.app/photo_profile/" + uid
	if err := db.Table("user").Save(user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("photo_profile").Where("user_id = ?", userId).First(&photoProfileDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	photoProfileDB.Path = path
	if err := db.Table("photo_profile").Save(&photoProfileDB).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	file, err := os.Open(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	fileStat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if fileStat.Size() != 0 {
		rProfile := router.Group("photo_profile")
		rProfile.Use(middleware.AuthWithToken)
		rProfile.GET(uid, func(c *gin.Context) {
			c.File(path)
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
