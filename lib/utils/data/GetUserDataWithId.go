package data

import (
	"errors"
	"github.com/database"
	"github.com/models"
)

func GetUserDataWithId(id string) (models.UserOtherEmailDescDB, error) {
	var userData models.UserOtherEmailDescDB

	db, err := database.ConnectDB()
	if err != nil {
		return userData, errors.New("internal server error")
	}

	if err := db.Table("user").Where("id = ?", id).First(&userData).Error; err != nil {
		return userData, errors.New("internal server error")
	}

	return userData, nil
}
