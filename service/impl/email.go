package impl

import (
	"github.com/icarus-go/captcha/config"
	"github.com/icarus-go/captcha/ext"
	"github.com/icarus-go/component/gins"
	"github.com/mojocn/base64Captcha"
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
