package database

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"walpapperCollectRestAPI/database/models"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./User.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&models.User{}, &models.WallpaperCollection{}); err != nil {
		return nil, err
	}
	return db, nil
}
