package impl

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/ext"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

var SmsStore *base64Captcha.Store

type SMS struct {
	Attribute *config.Attribute
}

func (i *SMS) Limit(ctx *gins.Context) error {
	return nil
}

func (i *SMS) Get(request *ext.Request) (*ext.Captcha, error) {
	result := ext.Captcha{}

	if SmsStore == nil {
		return nil, errors.New("短信验证码池为空")
	}

	if request.SMS.Mobile == "" {
		return nil, errors.New("手机验证码不允许为空")
	}

	if i.Attribute.Sender == nil {
		return nil, errors.New("短信发送方法为空")
	}

	code, err := i.Attribute.Sender.Send(request.SMS.Mobile, i.Attribute.Length)
	if err != nil {
		return nil, err
	}

	if err = (*SmsStore).Set(request.SMS.Mobile, code); err != nil {
		return nil, err
	}

	result.CaptchaID = request.SMS.Mobile

	return nil, nil
}

func (i *SMS) Verify(code, captchaID string) bool {
	return (*SmsStore).Verify(captchaID, code, false)
}
