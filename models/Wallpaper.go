package models

import "mime/multipart"

type Wallpaper struct {
	Image *multipart.FileHeader
}
