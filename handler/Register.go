package handler

import (
	"errors"
	"strconv"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/lib/tools"
)

func CreateUser(user models.User) error {
	/*
		this function is for filtering the data from JSON POST
		is the data is valid or not
		this function filtering user_name, email, and phone_number
	*/
	db, err := database.ConnectDB()
	if err != nil {
		return errors.New("internal server error")
	}

	//to validate number if the number contains some symbol
	if !tools.ValidationNumberPhone(strconv.Itoa(user.PhoneNumber)) {
		return errors.New("phone number is invalid")
	}

	//validating email is the string is email or not
	if !tools.ValidateEmail(user.Email) {
		return errors.New("email is invalid")
	}

	//searching if email is in database or not if not the data will pass this filter
	if err := db.Table("users").Where("email = ?", user.Email).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	//searching if user_name is not in database if not the data will pass this filter
	if err := db.Table("users").Where("user_name = ?", user.UserName).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	//searching if the phone number is not in database if not the data will pass this filter
	if err := db.Table("users").Where("phone_number = ?", user.PhoneNumber).First(&user).Error; err == nil {
		panic(err)
		return err
	}

	return nil
}
