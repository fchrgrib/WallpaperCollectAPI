package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

type UserLogin struct {
	UserName string `json:"user_name""`
	Password string `json:"password" binding:"min=6,max=24,required"`
}

type User struct {
	Id          uuid.UUID `json:"user_id" gorm:"primaryKey"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password" gorm:"varchar(300)" binding:"min=6,max=24,required"`
	Email       string    `json:"email"`
	PhoneNumber int       `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type WallpaperCollection struct {
	UserId    uuid.UUID `json:"user_id"`
	ImageName string    `json:"image_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Path      string `json:"path"`
}
