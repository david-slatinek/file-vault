package request

type Login struct {
	Code string `json:"code" binding:"required" form:"code"`
}
