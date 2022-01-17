package service

import (
	"errors"
	"github.com/icarus-go/captcha/config"
	"github.com/icarus-go/captcha/ext"
	"github.com/mojocn/base64Captcha"
	"time"

	"github.com/icarus-go/captcha/service/adapter"
	"github.com/icarus-go/captcha/service/impl"
	"github.com/icarus-go/component/gins"
)

type ICode interface {
	//Limit 限制验证码发送的次数或者频率
	Limit(ctx *gins.Context) error
	//Get 获取验证码
	Get(configuration *ext.Request) (*ext.Captcha, error)
	//Verify 验证 验证码是否正确
	Verify(code, captchaID string) bool
}

func New(cnf *config.Attribute) (ICode, error) {
	if cnf.Kind == "" {
		cnf.Kind = adapter.Image.Value()
	}

	if cnf.Length < 4 {
		cnf.Length = 4
	}

	if cnf.CollectNumber < 1 {
		cnf.CollectNumber = 10000
	}

	if cnf.ExpireSecond < 1 {
		cnf.ExpireSecond = 40 // 默认秒数
	}

	store := base64Captcha.NewMemoryStore(cnf.CollectNumber, cnf.ExpireSecond*time.Second)

	var instance ICode
	switch adapter.Adapter(cnf.Kind) {

	case adapter.SMS:
		sms := new(impl.SMS)
		sms.Attribute = cnf
		impl.SmsStore = &store
		instance = sms
	case adapter.Image:
		image := new(impl.Image)
		image.Attribute = cnf
		impl.ImageStore = store
		instance = image
	case adapter.Email:
		email := new(impl.Email)
		email.Attribute = cnf
		impl.EmailStore = &store
		instance = email
	default:
		return nil, errors.New("unknown adapter")
	}
	return instance, nil
}
