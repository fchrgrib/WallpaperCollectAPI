package database

import (
	"database/sql"
	"github.com/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	sqlDb, err := sql.Open("mysql", "root:Er6v71yyEsCLkcxgPPfx@tcp(containers-us-west-99.railway.app:7370)/railway")

	if err != nil {
		panic(err)
		return nil, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDb,
	}), &gorm.Config{})
	if err = db.AutoMigrate(&models.User{}, &models.WallpaperCollection{}); err != nil {
		return nil, err
	}
	return db, nil
}
