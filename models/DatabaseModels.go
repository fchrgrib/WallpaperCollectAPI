package models

import (
	_ "gorm.io/gorm"
	"time"
)

type UserPhotoProfileDB struct {
	UserId string `json:"user_id" gorm:"column:user_id;type:varchar(155);primaryKey"`
	Path   string `json:"path" gorm:"column:path"`
}

type UserOtherEmailDescDB struct {
	Id           string     `json:"user_id" gorm:"primaryKey;column:id;type:varchar(155);index"`
	UserName     string     `json:"user_name" gorm:"column:user_name;type:varchar(155);index"`
	Email        string     `json:"email" gorm:"primaryKey;column:email;index;type:varchar(155)"`
	PhoneNumber  string     `json:"phone_number" gorm:"column:phone_number"`
	PhotoProfile string     `json:"photo_profile" gorm:"column:photo_profile"`
	CreatedAt    *time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:datetime"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:datetime"`
}

type WallpaperCollectionDB struct {
	ImageId   string     `json:"image_id" gorm:"primaryKey;column:image_id"`
	ImageName string     `json:"image_name" gorm:"column:image_name"`
	UserId    string     `json:"user_id" gorm:"column:user_id;type:varchar(155)"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime"`
	Path      string     `json:"path" gorm:"column:path"`
}

type UserOtherEmailDB struct {
	Email    string `json:"email" gorm:"column:email;type:varchar(155);primaryKey" binding:"required"`
	Password string `json:"password" binding:"min=6,max=24,required" gorm:"column:password"`
}
