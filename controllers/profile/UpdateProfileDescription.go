package profile

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/libs/utils/data"
	"github.com/libs/utils/validation"
	"github.com/models"
	"net/http"
	"sync"
	"time"
)

func UpdateProfileDescription(c *gin.Context) {

	var (
		user       models.UserDescDB
		userUpdate models.UserDescDB
		wg         sync.WaitGroup
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

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		if !validation.ValidateEmail(userUpdate.Email) {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"status": "invalid email",
			})
			return
		}
	}()

	go func() {
		wg.Done()
		if !validation.ValidationNumberPhone(userUpdate.PhoneNumber) {
			c.JSON(http.StatusNotAcceptable, gin.H{
				"status": "invalid phone number",
			})
			return
		}
	}()

	go func() {
		wg.Done()
		t := time.Now().Local()
		user.UserName = userUpdate.UserName
		user.Email = userUpdate.Email
		user.PhoneNumber = userUpdate.PhoneNumber
		user.UpdatedAt = &t
		if err := db.Table("user").Save(user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": err.Error(),
			})
			return
		}
	}()

	wg.Wait()

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
