package tools

import (
	"errors"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
)

func GetUserDataWithId(id string) (models.User, error) {
	var userData models.User

	db, err := database.ConnectDB()
	if err != nil {
		return userData, errors.New("internal server error")
	}

	if err := db.Table("users").Where("id = ?", id).First(&userData).Error; err != nil {
		return userData, errors.New("internal server error")
	}

	return userData, nil
}
