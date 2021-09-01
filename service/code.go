package service

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
	"pmo-test4.yz-intelligence.com/kit/captcha/service/adapter"
	"pmo-test4.yz-intelligence.com/kit/captcha/service/impl"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
	"time"
)

type ICode interface {
	//Limit 限制验证码发送的次数或者频率
	Limit(ctx *gins.Context) error
	//Get 获取验证码
	Get() (model.Captcha, error)
	//Verify 验证 验证码是否正确
	Verify(code string, md model.Captcha) bool
}

func New(cnf *config.Attribute) (ICode, error) {
	if cnf.Length < 4 {
		cnf.Length = 4
	}

	if cnf.Height < 1 {
		cnf.Height = 40
	}

	if cnf.Width < 1 {
		cnf.Width = 80
	}

	if cnf.CollectNumber < 1 {
		cnf.CollectNumber = 10000
	}

	if cnf.Expire < time.Second*1 {
		cnf.Expire = time.Second * 10
	}

	if cnf.Kind == "" {
		cnf.Kind = adapter.Image.Value()
	}

	kind := adapter.Adapter(cnf.Kind)

	store := base64Captcha.NewMemoryStore(cnf.CollectNumber, cnf.Expire)

	var instance ICode

	switch kind {

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
