package tools

import (
	"github.com/database"
	"github.com/database/models"
	uuid2 "github.com/google/uuid"
	uuid "github.com/satori/go.uuid"
	"time"
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

	t := time.Now().Local()

	imageDb := models.WallpaperCollection{
		ImageId:   uuid2.UUID(imageUid),
		ImageName: imageName,
		UserId:    uuid2.UUID(userUid),
		Path:      path,
		CreatedAt: &t,
		UpdatedAt: &t,
	}

	if err := db.Table("wallpaper_collect").Create(&imageDb).Error; err != nil {
		return err
	}

	return nil
}
