package database

import (
	"github.com/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("././database/User.db"), &gorm.Config{})
	if err != nil {
		panic(err)
		return nil, err
	}
	if err = db.AutoMigrate(&models.User{}, &models.WallpaperCollection{}); err != nil {
		return nil, err
	}
	return db, nil
}
