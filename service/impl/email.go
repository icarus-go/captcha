package impl

import (
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model/request"
	"pmo-test4.yz-intelligence.com/kit/captcha/model/response"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type Email struct {
	Store     base64Captcha.Store
	Attribute *config.Attribute
}

func (i *Email) Limit(ctx *gins.Context) error {
	return nil
}

func (i *Email) Get(configuration *request.Configuration) (response.Captcha, error) {

	return response.Captcha{}, nil
}

func (i *Email) Verify(code, captchaID string) bool {
	return i.Store.Verify(captchaID, code, true)
}
