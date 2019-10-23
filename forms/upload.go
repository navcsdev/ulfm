package forms

import "mime/multipart"

type UploadForm struct {
	Image *multipart.FileHeader `form:"file" binding:"required"`
	Auth  string                `form:"auth" binding:"required"`
}
