package tools

import (
	"errors"
	"github.com/database"
	models "github.com/database/models"
)

func GetUserDataWithId(id string) (models.UserOtherEmailDesc, error) {
	var userData models.UserOtherEmailDesc

	db, err := database.ConnectDB()
	if err != nil {
		return userData, errors.New("internal server error")
	}

	if err := db.Table("user").Where("id = ?", id).First(&userData).Error; err != nil {
		return userData, errors.New("internal server error")
	}

	return userData, nil
}
