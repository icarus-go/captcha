package ext

type Captcha struct {
	CaptchaID string `json:"captchaID"`
	Email
	Image
}

type Email struct {
	Code string `json:"code" swaggerType:"string"`
}

type Image struct {
	ImageBase64 string `json:"image" swaggertype:"string"`
}
