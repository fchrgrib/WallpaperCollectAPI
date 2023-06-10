package profile

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/libs/middleware"
	"github.com/libs/utils/data"
	"github.com/models"
	"net/http"
	"os"
)

func UpdatePhotoProfile(c *gin.Context, router *gin.Engine) {
	var (
		user               models.UserDescDB
		photoProfileUser   models.UserPhotoProfileDB
		photoProfileUpload models.PhotoProfile
	)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	userId, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	if err := db.Table("user").Where("user_id = ?", userId).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}
	if err := db.Table("photo_profile").Where("user_id = ?", userId).First(&photoProfileUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}

	if err := os.Remove(photoProfileUser.Path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	if err := c.ShouldBind(&photoProfileUpload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	uid := uuid.New().String()
	imageName := photoProfileUpload.Image.Filename
	path := "././assets/" + userId + "/profile/" + uid + "_" + imageName

	if err := c.SaveUploadedFile(photoProfileUpload.Image, path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	photoProfileUser.Path = path
	user.PhotoProfile = "https://wallpapercollectapi-production.up.railway.app/photo_profile/" + uid
	if err := db.Save(&photoProfileUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}

	rProfile := router.Group("photo_profile")
	rProfile.Use(middleware.AuthWithToken)
	rProfile.Static(uid, path)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
