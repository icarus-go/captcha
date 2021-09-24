package response

type Captcha struct {
	CaptchaID string `json:"captchaID"`
	Email     *Email `json:"email,omitempty"`
	Image     *Image `json:"image,omitempty"`
}

type Email struct {
	Code string `json:"code"`
}

type Image struct {
	ImageBase64 string `json:"image"`
}
