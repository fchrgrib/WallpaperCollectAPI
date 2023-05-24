package register

import (
	"bytes"
	"encoding/json"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/utils/data"
	"github.com/lib/utils/oauth2utility"
	"github.com/models"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"os"
	"time"
)

// CreateUserAuthFacebook TODO make redirect to application
func CreateUserAuthFacebook(c *gin.Context) {
	var (
		userDesc         models.UserOtherEmailDescDB
		userPhotoProfile models.UserPhotoProfileDB
		userProfiles     models.UserProfileFacebook
	)

	db, err := database.ConnectDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	token, err := oauth2utility.GetFacebookConfRegis().Exchange(oauth2.NoContext, c.Query("code"))

	client := oauth2utility.GetFacebookConfRegis().Client(oauth2.NoContext, token)

	userProfile, err := client.Get("https://graph.facebook.com/v13.0/me?fields=id,name,email,picture")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer userProfile.Body.Close()

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, userProfile.Body)

	if err := json.Unmarshal(resBody.Bytes(), &userProfiles); err != nil {
		return
	}

	t := time.Now().Local()

	userDesc.Id = uuid.New().String()
	userDesc.UserName = userProfiles.Name
	userDesc.Email = userProfiles.Email
	userDesc.PhoneNumber = ""
	userDesc.PhotoProfile = userProfiles.Picture.Data.URL
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

func RedirectFacebookRegisterController(c *gin.Context) {
	state := data.RandToken()
	c.Redirect(http.StatusSeeOther, oauth2utility.GetFacebookRegisterURL(state))
}
