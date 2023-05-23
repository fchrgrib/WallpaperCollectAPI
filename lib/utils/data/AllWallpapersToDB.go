package data

import (
	"github.com/database"
	"github.com/models"
	"time"
)

func AllWallpaperToDB(id string, path string, uid string, imageName string) error {
	db, err := database.ConnectDB()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	t := time.Now().Local()

	imageDb := models.WallpaperCollectionDB{
		ImageId:   uid,
		ImageName: imageName,
		UserId:    id,
		Path:      path,
		CreatedAt: &t,
		UpdatedAt: &t,
	}

	if err := db.Table("wallpaper_collect").Create(&imageDb).Error; err != nil {
		return err
	}

	return nil
}
