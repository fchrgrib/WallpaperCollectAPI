package database

import (
	"database/sql"
	"github.com/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	sqlDb, err := sql.Open("mysql", "root:Er6v71yyEsCLkcxgPPfx@tcp(containers-us-west-99.railway.app:7370)/railway?parseTime=true")

	if err != nil {
		panic(err)
		return nil, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err = db.Table("user").AutoMigrate(&models.UserOtherEmailDesc{}); err != nil {
		return nil, err
	}
	if err = db.Table("wallpaper_collect").AutoMigrate(&models.WallpaperCollectionDB{}); err != nil {
		return nil, err
	}
	if err = db.Table("user_other_email").AutoMigrate(&models.UserOtherEmail{}); err != nil {
		return nil, err
	}
	return db, nil
}
