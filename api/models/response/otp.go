package response

type OTP struct {
	Key string `json:"key"`
	URL string `json:"url"`
}
