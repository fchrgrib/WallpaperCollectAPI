package models

import "mime/multipart"

type PhotoProfile struct {
	Image *multipart.FileHeader
}
