package models

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
)

type UserLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password" binding:"min=6,max=24,required"`
}

type User struct {
	Id           uuid.UUID `json:"user_id" gorm:"primaryKey"`
	UserName     string    `json:"user_name" gorm:"primaryKey"`
	Email        string    `json:"email"`
	Password     string    `json:"password" binding:"min=6,max=24,required"`
	PhoneNumber  string    `json:"phone_number"`
	PhotoProfile string    `json:"photo_profile"`
}

type WallpaperCollection struct {
	ImageId   uuid.UUID `json:"image_id" gorm:"primaryKey"`
	ImageName string    `json:"image_name"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Path      string `json:"path"`
}

type UserOtherEmailDesc struct {
	Id           uuid.UUID `json:"user_id" gorm:"primaryKey;column:id"`
	UserName     string    `json:"user_name" gorm:"primaryKey;column:user_name"`
	Email        string    `json:"email" gorm:"unique;column:email"`
	PhoneNumber  string    `json:"phone_number" gorm:"column:phone_number"`
	PhotoProfile string    `json:"photo_profile" gorm:"column:photo_profile"`
	CreatedAt    string    `json:"created_at" gorm:"column:created_at;type:time"`
	UpdatedAt    string    `json:"updated_at" gorm:"column:updated_at;type:time"`
	DeletedAt    string    `json:"deleted_at" gorm:"column:deleted_at;type:time"`
}

type WallpaperCollectionDB struct {
	ImageId   uuid.UUID `json:"image_id" gorm:"primaryKey;column:image_id"`
	ImageName string    `json:"image_name" gorm:"column:image_name"`
	UserId    uuid.UUID `json:"user_id" gorm:"column:user_id"`
	CreatedAt string    `gorm:"column:created_at;type:time"`
	UpdatedAt string    `gorm:"column:updated_at;type:time"`
	DeletedAt string    `gorm:"column:deleted_at;type:time"`
	Path      string    `json:"path" gorm:"column:path"`
	User      User      `gorm:"foreignKey:Id;references:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserOtherEmail struct {
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"password" binding:"min=6,max=24,required" gorm:"column:password"`
	User     User   `gorm:"foreignKey:UserName;references:user_name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
