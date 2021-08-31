package controller

import (
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/service"
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

	code, err := service.New(cnf)
	if err != nil {
		return
	}

	result, err := code.Get()
	if err != nil {
		return
	}
	println("image:", result.Image, ", captchaID:", result.CaptchaID)

	value := "" // debug setting
	if verify := code.Verify(value, result); verify {
		println("result: ", verify)
		return
	}
	println("result: false")
}
