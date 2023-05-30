package register

import (
	"bytes"
	"encoding/json"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/libs/utils/data"
	"github.com/libs/utils/oauth2utility"
	"github.com/models"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
	"time"
)

// CreateUserAuthGoogle TODO make redirect to application
func CreateUserAuthGoogle(c *gin.Context) {
	var (
		userDesc         models.UserOtherEmailDescDB
		userPhotoProfile models.UserPhotoProfileDB
	)

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	token, err := oauth2utility.GetGoogleConfRegis().Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": err,
		})
		return
	}

	client := oauth2utility.GetGoogleConfRegis().Client(oauth2.NoContext, token)

	userProfile, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo?alt=json&access_token=" + token.AccessToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	defer userProfile.Body.Close()

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, userProfile.Body)

	var GoogleUserRes map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return
	}

	user := &models.UserOtherEmailDescDB{
		Email:        GoogleUserRes["email"].(string),
		UserName:     GoogleUserRes["family_name"].(string),
		PhotoProfile: GoogleUserRes["picture"].(string),
		PhoneNumber:  "",
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

// RedirectGoogleRegisterController TODO make real Login with JWT key and Register User
func RedirectGoogleRegisterController(c *gin.Context) {
	state := data.RandToken()
	c.JSON(http.StatusOK, gin.H{
		"url":    oauth2utility.GetGoogleRegisterURL(state),
		"status": "ok",
	})
	return
}
