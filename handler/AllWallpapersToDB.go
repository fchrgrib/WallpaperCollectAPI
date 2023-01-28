package handler

import (
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"time"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
)

func AllWallpaperToDB(id string, path string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	uid, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	imageDb := models.WallpaperCollection{
		UserId:    uuid2.UUID(uid),
		Path:      path,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	if err := db.Table("wallpaper_collections").Create(&imageDb).Error; err != nil {
		return err
	}

	return nil
}
