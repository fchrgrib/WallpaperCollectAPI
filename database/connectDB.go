package database

import (
	"database/sql"
	"github.com/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	sqlDb, err := sql.Open("mysql", "fchrgrib2310:Fchrgrib2310*@tcp(192.168.43.236:3306)/wallpaperdb?parseTime=true")

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
