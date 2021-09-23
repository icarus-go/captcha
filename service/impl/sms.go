package impl

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
	"time"
)

type SMS struct {
	Store     base64Captcha.Store
	Attribute *config.Attribute
}

func (i *SMS) Limit(ctx *gins.Context) error {
	return nil
}

func (i *SMS) Get(configuration model.Configuration) (model.Captcha, error) {
	if i.Attribute.Length < 4 {
		i.Attribute.Length = 4
	}

	if i.Attribute.CollectNumber < 1 {
		i.Attribute.CollectNumber = 10000
	}

	if i.Attribute.Expire < time.Second*1 {
		i.Attribute.Expire = time.Second * 40
	}

	result := model.Captcha{}
	if i.Attribute.Sender == nil {
		return result, errors.New("短信发送方法为空")
	}

	captchaID, code, err := i.Attribute.Sender.Send(configuration.SMS.Mobile, i.Attribute.Length)
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
