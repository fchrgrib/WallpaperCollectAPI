package register

import (
	"bytes"
	"encoding/json"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/models"
	"io"
	"net/http"
	"os"
	"time"
)

// CreateUserAuthGoogle TODO make redirect to application
func CreateUserAuthGoogle(c *gin.Context) {
	var (
		userDesc         models.UserDescDB
		userPhotoProfile models.UserPhotoProfileDB
		googleToken      models.GoogleToken
	)

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&googleToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	userProfile, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + googleToken.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err,
			})
			return
		}
	}(userProfile.Body)

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, userProfile.Body)

	var GoogleUserRes map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return
	}

	user := &models.UserDescDB{
		Email:        GoogleUserRes["email"].(string),
		UserName:     GoogleUserRes["family_name"].(string),
		PhotoProfile: GoogleUserRes["picture"].(string),
		PhoneNumber:  "",
	}

	if err := db.Table("user").Where("email = ?", user.Email).First(&models.UserDescDB{}).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "cannot create new user because email was existed",
		})
		return
	}

	t := time.Now().Local()

	userDesc.Id = uuid.New().String()
	userDesc.UserName = user.UserName
	userDesc.Email = user.Email
	userDesc.PhoneNumber = user.PhoneNumber
	userDesc.PhotoProfile = user.PhotoProfile
	userDesc.CreatedAt = &t
	userDesc.UpdatedAt = &t
	userDesc.DeletedAt = nil

	pathProfile := "././assets/" + userDesc.Id + "/profile"

	userPhotoProfile = models.UserPhotoProfileDB{
		UserId: userDesc.Id,
		Path:   "",
	}

	if err := os.MkdirAll("././assets/"+userDesc.Id+"/wallpaper_collection", os.ModePerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := os.MkdirAll(pathProfile, os.ModePerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		return
	}

	if err := db.Table("user").Create(&userDesc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}
	if err := db.Table("photo_profile").Create(&userPhotoProfile).Error; err != nil {
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
