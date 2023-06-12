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

type UserProfileFacebook struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}

type Images struct {
	ImageUrl  string `json:"image_url"`
	ImageName string `json:"image_name"`
	ImageId   string `json:"image_id"`
}

type PhotoProfile struct {
	Image *multipart.FileHeader
}

type Wallpaper struct {
	Image *multipart.FileHeader
}

type GoogleToken struct {
	Token string `json:"token"`
}
