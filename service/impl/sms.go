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
	Attribute *config.Attribute
}

func (i *SMS) Limit(ctx *gins.Context) error {
	return nil
}

func (i *SMS) Get() (model.Captcha, error) {
	result := model.Captcha{}
	if i.Attribute.Sender == nil {
		return result, errors.New("短信发送方法为空")
	}

	captchaID, code, err := i.Attribute.Sender.Send(i.Attribute.Length)
	if err != nil {
		return result, err
	}

	if err = i.Store.Set(captchaID, code); err != nil {
		return result, err
	}

	result.CaptchaID = captchaID
	return result, nil
}

func (i *SMS) Verify(code, captchaID string) bool {
	return i.Store.Verify(captchaID, code, false)
}
