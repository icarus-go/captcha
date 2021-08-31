package impl

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type SMS struct {
	Store     base64Captcha.Store
	Attribute config.Attribute
}

func (i *SMS) Limit(ctx *gins.Context) error {
	return nil
}

func (SMS) Get() (model.Captcha, error) {
	// todo 发送验证码，并且返回captchaID
	return model.Captcha{}, errors.New("SMS not implemented")
}

func (i *SMS) Verify(code string, md model.Captcha) bool {
	return i.Store.Verify(md.CaptchaID, code, true)
}
