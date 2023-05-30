package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectDB() (*gorm.DB, error) {
	if err := godotenv.Load("././handlers/env/database.env"); err != nil {
		return nil, err
	}

	databaseType := os.Getenv("DATABASE_TYPE")
	userName := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	schema := os.Getenv("SCHEMA")

	sqlDb, err := sql.Open(databaseType, userName+":"+password+"@tcp("+host+":"+port+")/"+schema+"?parseTime=true")

	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err = db.Table("user").AutoMigrate(&models.UserOtherEmailDescDB{}); err != nil {
		return nil, err
	}
	if err = db.Table("wallpaper_collect").AutoMigrate(&models.WallpaperCollectionDB{}); err != nil {
		return nil, err
	}
	if err = db.Table("user_other_email").AutoMigrate(&models.UserOtherEmailDB{}); err != nil {
		return nil, err
	}
	if err = db.Table("photo_profile").AutoMigrate(&models.UserPhotoProfileDB{}); err != nil {
		return nil, err
	}
	return db, nil
}
