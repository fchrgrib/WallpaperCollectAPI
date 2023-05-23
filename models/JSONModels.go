package models

import "mime/multipart"

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"min=6,max=24,required"`
}

type User struct {
	Id           string `json:"user_id" gorm:"primaryKey"`
	UserName     string `json:"user_name" gorm:"primaryKey"`
	Email        string `json:"email"`
	Password     string `json:"password" binding:"min=6,max=24,required"`
	PhoneNumber  string `json:"phone_number"`
	PhotoProfile string `json:"photo_profile"`
}

type PhotoProfile struct {
	Image *multipart.FileHeader
}

type Wallpaper struct {
	Image *multipart.FileHeader
}
