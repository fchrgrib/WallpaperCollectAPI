package handler

import (
	"errors"
	"github.com/database"
	"github.com/database/models"
	"github.com/lib/tools"
)

func CreateUser(user models.User) error {
	/***
		this function is for filtering the data from JSON POST
		is the data is valid or not
		this function filtering user_name, email, and phone_number
	***/
	db, err := database.ConnectDB()
	if err != nil {
		return errors.New("internal server error")
	}

	//to validate number if the number contains some symbol
	if !tools.ValidationNumberPhone(user.PhoneNumber) {
		return errors.New("phone number is invalid")
	}

	//validating email is the string is email or not
	if !tools.ValidateEmail(user.Email) {
		return errors.New("email is invalid")
	}

	//searching if email, user_name, phone_number are in database or not if not the data will pass this filter
	if err := db.Table("user").Where("email = ? AND user_name = ? AND phone_number = ?", user.Email, user.UserName, user.PhoneNumber).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	return nil
}
