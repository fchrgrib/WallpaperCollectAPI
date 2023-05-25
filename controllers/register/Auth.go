package register

import (
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/handlers/authandlers"
	"github.com/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func CreateUserAuth(c *gin.Context) {
	/*
		this function is for inserting the user register data form JSON to the database
		and make sure the data is has error or not
		if data have an error the JSON will POST the error
	*/
	var (
		userDesc         models.UserOtherEmailDescDB
		userLog          models.UserOtherEmailDB
		user             models.User
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
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}
	if err := authandlers.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	t := time.Now().Local()

	userDesc.Id = uuid.New().String()
	userDesc.UserName = user.UserName
	userDesc.Email = user.Email
	userDesc.PhoneNumber = user.PhoneNumber
	userDesc.CreatedAt = &t
	userDesc.UpdatedAt = &t
	userDesc.DeletedAt = nil

	userLog.Email = user.Email
	userLog.Password = string(hashPass)

	pathProfile := "././assets/" + userDesc.Id + "/profile"

	user.PhotoProfile = pathProfile
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
	if err := db.Table("user_other_email").Create(&userLog).Error; err != nil {
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
