package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/handler"
)

func CreateUserAuth(c *gin.Context) {
	/*
		this function is for inserting the user register data form JSON to the database
		and make sure the data is has error or not
		if data have an error the JSON will POST the error
	*/
	var user models.User
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

	user.Id = uuid.New()
	user.Password = string(hashPass)
	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()

	pathProfile := "././assets/" + user.Id.String() + "/profile"

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

	if err := db.Table("users").Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})

}
