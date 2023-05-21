package controller

import (
	"github.com/database"
	"github.com/database/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/handler"
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
	var user models.User
	var userDesc models.UserOtherEmailDesc
	var userLog models.UserOtherEmail
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
	if err := handler.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	userDesc.Id = uuid.New()
	userDesc.UserName = user.UserName
	userDesc.Email = user.Email
	userDesc.PhoneNumber = user.PhoneNumber
	userDesc.CreatedAt = time.Now().Local()
	userDesc.UpdatedAt = time.Now().Local()
	userDesc.DeletedAt = nil

	userLog.UserName = user.UserName
	userLog.Password = string(hashPass)

	pathProfile := "././assets/" + userDesc.Id.String() + "/profile"

	user.PhotoProfile = pathProfile

	if err := os.MkdirAll("././assets/"+user.Id.String()+"/wallpaper_collection", os.ModePerm); err != nil {
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
		panic(err)
		return
	}
	if err := db.Table("user_other_email").Create(&userLog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
