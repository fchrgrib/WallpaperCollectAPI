package models

import (
	"github.com/google/uuid"
	_ "gorm.io/gorm"
	"time"
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
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Path      string `json:"path"`
}

type UserOtherEmailDesc struct {
	Id           uuid.UUID  `json:"user_id" gorm:"primaryKey;column:id"`
	UserName     string     `json:"user_name" gorm:"primaryKey;column:user_name"`
	Email        string     `json:"email" gorm:"unique;column:email"`
	PhoneNumber  string     `json:"phone_number" gorm:"column:phone_number"`
	PhotoProfile string     `json:"photo_profile" gorm:"column:photo_profile"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;type:datetime"`
	UpdatedAt    *time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime"`
	DeletedAt    *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:datetime"`
}

type WallpaperCollectionDB struct {
	ImageId   uuid.UUID  `json:"image_id" gorm:"primaryKey;column:image_id"`
	ImageName string     `json:"image_name" gorm:"column:image_name"`
	UserId    uuid.UUID  `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime"`
	Path      string     `json:"path" gorm:"column:path"`
	User      User       `gorm:"foreignKey:Id;references:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserOtherEmail struct {
	UserName string `json:"user_name" gorm:"column:user_name"`
	Password string `json:"password" binding:"min=6,max=24,required" gorm:"column:password"`
	User     User   `gorm:"foreignKey:UserName;references:user_name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
