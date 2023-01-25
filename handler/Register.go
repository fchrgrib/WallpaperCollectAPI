package handler

import (
	"net/http"
	"strconv"
	"time"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
	"walpapperCollectRestAPI/lib/tools"
)

func CreateUser(user models.User) (httpResponse models.HTTPResponse, httpStatusCode int) {
	db, err := database.ConnectDB()
	if err != nil {
		return
	}

	if tools.ValidationNumberPhone(strconv.Itoa(user.PhoneNumber)) {
		httpResponse.Message = "you input the wrong number phone format"
		httpStatusCode = http.StatusBadRequest
		return
	}

	if tools.ValidateEmail(user.Email) {
		httpResponse.Message = "you input the wrong e-mail"
		httpStatusCode = http.StatusBadRequest
		return
	}

	if err := db.Table("User").First(&user).Error; err == nil {
		httpResponse.Message = "e-mail has already exist"
		httpStatusCode = http.StatusBadRequest
		return
	}

	if err = db.Table("User").Where("userName = ?", user.UserName).First(&user).Error; err == nil {
		httpResponse.Message = "username has already exist"
		httpStatusCode = http.StatusBadRequest
		return
	}

	if err = db.Table("User").Where("phoneNumber = ?", user.PhoneNumber).First(&user).Error; err == nil {
		httpResponse.Message = "phone number has already exist"
		httpStatusCode = http.StatusBadRequest
		return
	}

	//if service.SendEmail(authFinal.Email, model.EmailTypeVerification) {
	//	authFinal.VerifyEmail = model.EmailNotVerified
	//}

	user.CreatedAt = time.Now().Local()
	user.UpdatedAt = time.Now().Local()

	dbCreate := db.Begin()

	if err = dbCreate.Create(&user).Error; err != nil {
		httpResponse.Message = "phone number has already exist"
		httpStatusCode = http.StatusBadRequest
		return
	}
	dbCreate.Commit()

	httpResponse.Message = "statusOK"
	httpStatusCode = http.StatusCreated
	return

}
