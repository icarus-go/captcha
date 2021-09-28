package service

import (
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/ext"
	"testing"
)

func Test_captcha_CommonGet(t *testing.T) {
	cnf := config.Attribute{}

	cnf.Length = 4
	cnf.CollectNumber = 10000
	cnf.Expire = 10
	cnf.Height = 60
	cnf.Width = 300
	cnf.Kind = "image"

	code, err := New(&cnf)
	if err != nil {
		return
	}

	result, err := code.Get(nil)
	if err != nil {
		return
	}
	//println("image:", result.Image, ", captchaID:", result.CaptchaID)

	value := "" // debug setting
	if verify := code.Verify(value, result.CaptchaID); verify {
		println("result: ", verify)
		return
	}
	println("result: false")
}

type Sms struct{}

func (Sms) Send(length int) (string, string, error) {
	return "13129627708", "100010", nil
}

func Test_captcha_SMSGet(t *testing.T) {
	cnf := config.Attribute{}

	cnf.Length = 4
	cnf.CollectNumber = 10000
	cnf.Expire = 10
	cnf.Height = 60
	cnf.Width = 300
	cnf.Kind = "ext"
	//cnf.Sender = Sms{}

	code, err := New(&cnf)
	if err != nil {
		return
	}

	result, err := code.Get(&ext.Request{
		Email: ext.EmailGenerate{},
		SMS:   ext.SmsGenerator{},
		Image: struct{}{},
	})
	if err != nil {
		return
	}
	//println("image:", result.Image, ", captchaID:", result.CaptchaID)

	value := "" // debug setting
	if verify := code.Verify(value, result.CaptchaID); verify {
		println("result: ", verify)
		return
	}
	println("result: false")
}
