package authandlers

import (
	"github.com/database"
	"github.com/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(userInput models.UserLogin) (models.UserLogin, error) {
	db, err := database.ConnectDB()
	var userLogInfo models.UserLogin
	var userDB models.UserLogin

	if err != nil {
		return userLogInfo, err
	}
	if err := db.Table("user_other_email").Where("email = ?", userInput.Email).First(&userDB).Error; err != nil {
		panic(err)
		return userLogInfo, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(userInput.Password)); err != nil {
		panic(err)
		return userLogInfo, err
	}
	userLogInfo.Email = userDB.Email
	userLogInfo.Password = userDB.Password
	return userLogInfo, nil
}
