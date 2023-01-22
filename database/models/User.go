package models

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

type UserLogin struct {
	Id           int            `json:"id" gorm:"primaryKey"`
	UserName     string         `json:"userName""`
	Password     string         `json:"password" binding:"min=6,max=24,required"`
	Email        string         `json:"email"`
	PhoneNumber  int            `json:"phoneNumber"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Token        string         `json:"token"`
	UpdatedToken string         `json:"updateToken"`
}

type User struct {
	Id          int            `json:"id" gorm:"primaryKey"`
	UserName    string         `json:"userName""`
	Password    string         `json:"password" binding:"min=6,max=24,required"`
	Email       string         `json:"email"`
	PhoneNumber int            `json:"phoneNumber"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type WallpaperCollection struct {
	ImageId   int    `json:"imageId" gorm:"primaryKey"`
	UserId    int    `json:"userId"`
	ImageName string `json:"imageName"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Path      string `json:"path"`
}
