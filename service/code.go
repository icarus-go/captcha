package service

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
	"pmo-test4.yz-intelligence.com/kit/captcha/service/adapter"
	"pmo-test4.yz-intelligence.com/kit/captcha/service/impl"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type ICode interface {
	//Limit 限制验证码发送的次数或者频率
	Limit(ctx *gins.Context) error
	//Get 获取验证码
	Get(configuration model.Configuration) (model.Captcha, error)
	//Verify 验证 验证码是否正确
	Verify(code, captchaID string) bool
}

func New(cnf *config.Attribute) (ICode, error) {
	if cnf.Kind == "" {
		cnf.Kind = adapter.Image.Value()
	}

	store := base64Captcha.NewMemoryStore(cnf.CollectNumber, cnf.Expire)

	var instance ICode

	switch adapter.Adapter(cnf.Kind) {

	case adapter.SMS:
		sms := new(impl.SMS)
		sms.Attribute = cnf
		sms.Store = store
		instance = sms
	case adapter.Image:
		image := new(impl.Image)
		image.Attribute = cnf
		image.Store = store
		instance = image
	case adapter.Email:
		email := new(impl.Email)
		email.Attribute = cnf
		email.Store = store
		instance = email
	default:
		return nil, errors.New("unknown adapter")
	}
	return instance, nil
}
