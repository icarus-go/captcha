package ext

type Captcha struct {
	CaptchaID string `json:"captchaID"`
	Email     Email  `json:"email"`
	Image     Image  `json:"image"`
}

type Email struct {
	Code string `json:"code"`
}

type Image struct {
	ImageBase64 string `json:"image"`
}
