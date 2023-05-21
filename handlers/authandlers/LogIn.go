package authandlers

import (
	"github.com/database"
	"github.com/database/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(userInput models.UserLogin) (models.UserLogin, error) {
	db, err := database.ConnectDB()
	var userLogInfo models.UserLogin
	var userDB models.UserLogin

	if err != nil {
		panic(err)
		return userLogInfo, err
	}
	if err := db.Table("user_other_email").Where("user_name = ?", userInput.UserName).First(&userDB).Error; err != nil {
		panic(err)
		return userLogInfo, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(userInput.Password)); err != nil {
		panic(err)
		return userLogInfo, err
	}
	userLogInfo.UserName = userDB.UserName
	userLogInfo.Password = userDB.Password
	return userLogInfo, nil
}
