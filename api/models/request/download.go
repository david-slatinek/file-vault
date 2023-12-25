package request

type FileDownload struct {
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required,min=12"`
}
