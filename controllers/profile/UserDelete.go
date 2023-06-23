package profile

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/libs/utils/data"
	"github.com/models"
	"net/http"
	"os"
	"sync"
)

func UserDelete(c *gin.Context) {

	var (
		userDelete           models.UserDescDB
		wallpaperCollections []models.WallpaperCollectionDB
		photoProfileDelete   models.UserPhotoProfileDB
		wg                   sync.WaitGroup
	)

	userId, err := data.GetUserIdFromCookies(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": err.Error(),
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	wg.Add(3)

	go func() {
		defer wg.Done()
		if _ = db.Table("user").Where("id = ?", userId).First(&userDelete); userDelete.Id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "user not found",
			})
			return
		}
	}()

	go func() {
		defer wg.Done()
		if _ = db.Table("photo_profile").Where("user_id = ?", userId).First(&photoProfileDelete); photoProfileDelete.UserId == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "photo profile not found",
			})
			return
		}

		if photoProfileDelete.Path != "" {
			if err := os.Remove(photoProfileDelete.Path); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": "path not found",
				})
				return
			}
		}

		if err := db.Table("photo_profile").Delete(photoProfileDelete).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
			})
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := db.Table("wallpaper_collect").Where("user_id = ?", userId).Find(&wallpaperCollections).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "wallpapers not found",
			})
			return
		}

		if len(wallpaperCollections) != 0 {
			for _, value := range wallpaperCollections {
				if err := os.Remove(value.Path); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"status": "path not found",
					})
					return
				}
			}

			if err := db.Table("wallpaper_collect").Delete(wallpaperCollections).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": err.Error(),
				})
				return
			}
		}
	}()

	wg.Wait()

	if err := db.Table("user").Delete(userDelete).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
