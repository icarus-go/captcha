package response

type Captcha struct {
	CaptchaID string `json:"captchaID"`
	Image     string `json:"image"`
}
