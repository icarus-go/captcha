package impl

import (
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/ext"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

var EmailStore *base64Captcha.Store

type Email struct {
	Attribute *config.Attribute
}

func (i *Email) Limit(ctx *gins.Context) error {
	return nil
}

func (i *Email) Get(configuration *ext.Request) (*ext.Captcha, error) {

	return &ext.Captcha{}, nil
}

func (i *Email) Verify(code, captchaID string) bool {
	return (*EmailStore).Verify(captchaID, code, true)
}
