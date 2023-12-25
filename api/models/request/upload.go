package request

import "mime/multipart"

type FileUpload struct {
	File  *multipart.FileHeader `form:"file" binding:"required"`
	Login `form:"code" binding:"required"`
}
