package handler

import (
	"github.com/database"
	"github.com/database/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(userInput models.UserLogin) (models.User, error) {
	db, err := database.ConnectDB()
	userDB := models.User{}
	if err != nil {
		panic(err)
		return userDB, err
	}
	if err := db.Table("users").Where("user_name = ?", userInput.UserName).First(&userDB).Error; err != nil {
		panic(err)
		return userDB, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(userInput.Password)); err != nil {
		panic(err)
		return userDB, err
	}
	return userDB, nil
}
