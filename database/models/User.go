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
	Id           uuid.UUID      `json:"user_id" gorm:"primaryKey"`
	UserName     string         `json:"user_name"`
	Password     string         `json:"password" gorm:"varchar(300)" binding:"min=6,max=24,required"`
	Email        string         `json:"email"`
	PhoneNumber  int            `json:"phone_number"`
	Description  string         `json:"description"`
	PhotoProfile string         `json:"photo_profile"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
type UserUpdate struct {
	UserName    string         `json:"user_name"`
	Password    string         `json:"password" gorm:"varchar(300)" binding:"min=6,max=24,required"`
	Email       string         `json:"email"`
	PhoneNumber int            `json:"phone_number"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type WallpaperCollection struct {
	ImageId   uuid.UUID `json:"image_id" gorm:"primaryKey"`
	ImageName string    `json:"image_name"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Path      string `json:"path"`
}
