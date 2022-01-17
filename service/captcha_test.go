package service

import (
	"github.com/icarus-go/captcha/config"
	"github.com/icarus-go/captcha/ext"
	"testing"
)

func Test_captcha_CommonGet(t *testing.T) {
	cnf := config.Attribute{}

	cnf.Length = 4
	cnf.CollectNumber = 10000
	cnf.ExpireSecond = 10
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

type SmsSender struct{}

func (*SmsSender) Send(mobile string, length int) (string, error) {
	println("helloWorld")
	return "", nil
}

func Test_captcha_SMSGet(t *testing.T) {
	cnf := config.Attribute{}

	cnf.Length = 4
	cnf.CollectNumber = 10000
	cnf.ExpireSecond = 10
	cnf.Height = 60
	cnf.Width = 300
	cnf.Kind = "sms"
	cnf.Sender = &SmsSender{}

	code, err := New(&cnf)
	if err != nil {
		return
	}

	println(code)

	code.Get(&ext.Request{

		SMS: ext.SmsGenerator{
			Mobile: "13129627708",
		},
	})
	//result, err := code.Get(&ext.Request{
	//	Email: ext.EmailGenerate{},
	//	SMS:   ext.SmsGenerator{},
	//	Image: struct{}{},
	//})
	//if err != nil {
	//	return
	//}
	////println("image:", result.Image, ", captchaID:", result.CaptchaID)
	//
	//value := "" // debug setting
	//if verify := code.Verify(value, result.CaptchaID); verify {
	//	println("result: ", verify)
	//	return
	//}
	println("result: false")
}
