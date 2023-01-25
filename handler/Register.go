package handler

import (
	"errors"
	"strconv"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/lib/tools"
)

func CreateUser(user models.User) error {
	db, err := database.ConnectDB()
	if err != nil {
		return errors.New("internal server error")
	}

	if tools.ValidationNumberPhone(strconv.Itoa(user.PhoneNumber)) {
		return errors.New("phone number is invalid")
	}

	if !tools.ValidateEmail(user.Email) {
		return errors.New("email is invalid")
	}

	if err := db.Table("users").Where("email = ?", user.Email).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	if err := db.Table("users").Where("user_name = ?", user.UserName).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	if err := db.Table("users").Where("phone_number = ?", user.PhoneNumber).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	//if service.SendEmail(authFinal.Email, model.EmailTypeVerification) {
	//	authFinal.VerifyEmail = model.EmailNotVerified
	//}
	return nil
}
