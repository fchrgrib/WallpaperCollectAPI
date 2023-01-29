package handler

import (
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"time"
	"walpapperCollectRestAPI/database"
	"walpapperCollectRestAPI/database/models"
)

func AllWallpaperToDB(id string, path string, uid string, imageName string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	imageUid, err := uuid.FromString(uid)
	userUid, err := uuid.FromString(id)
	if err != nil {
		return err
	}

	imageDb := models.WallpaperCollection{
		ImageId:   uuid2.UUID(imageUid),
		ImageName: imageName,
		UserId:    uuid2.UUID(userUid),
		Path:      path,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	if err := db.Table("wallpaper_collections").Create(&imageDb).Error; err != nil {
		return err
	}

	return nil
}
