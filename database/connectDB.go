package database

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	if err := godotenv.Load("././handlers/env/database.env"); err != nil {
		return nil, err
	}

	//databaseType := os.Getenv("DATABASE_TYPE")
	//userName := os.Getenv("USER_NAME")
	//password := os.Getenv("PASSWORD")
	//host := os.Getenv("HOST")
	//port := os.Getenv("PORT")
	//schema := os.Getenv("SCHEMA")

	sqlDb, err := sql.Open("mysql", "root:0MVkzZzGGYCminYTLrEJ@tcp(containers-us-west-152.railway.app:6553)/railway?parseTime=true")

	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err = db.Table("user").AutoMigrate(&models.UserDescDB{}); err != nil {
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
