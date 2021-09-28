package impl

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/ext"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
	"time"
)

var SmsStore *base64Captcha.Store

type SMS struct {
	Attribute *config.Attribute
}

func (i *SMS) Limit(ctx *gins.Context) error {
	return nil
}

func (i *SMS) Get(configuration *ext.Request) (*ext.Captcha, error) {
	result := ext.Captcha{}

	if configuration.SMS.Mobile == "" {
		return nil, errors.New("手机号码不允许为空")
	}

	if i.Attribute.Length < 4 {
		i.Attribute.Length = 4
	}

	if i.Attribute.CollectNumber < 1 {
		i.Attribute.CollectNumber = 10000
	}

	if i.Attribute.ExpireSecond < time.Second*1 {
		i.Attribute.ExpireSecond = time.Second * 40
	}

	if i.Attribute.Sender == nil {
		return nil, errors.New("短信发送方法为空")
	}

	if SmsStore == nil {
		return nil, errors.New("短信验证码池为空")
	}

	code, err := i.Attribute.Sender.Send(configuration.SMS.Mobile, i.Attribute.Length)
	if err != nil {
		return nil, err
	}

	if err = (*SmsStore).Set(configuration.SMS.Mobile, code); err != nil {
		return nil, err
	}

	result.CaptchaID = configuration.SMS.Mobile

	return nil, nil
}

func (i *SMS) Verify(code, captchaID string) bool {
	return (*SmsStore).Verify(captchaID, code, false)
}
