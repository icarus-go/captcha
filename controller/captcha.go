package controller

import (
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/service"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type captcha struct {
	Cnf config.Captcha
}

func NewCaptcha(cnf config.Captcha) *captcha {

	captcha := new(captcha)

	captcha.Cnf = cnf

	return captcha
}

func (c *captcha) Get(ctx *gins.Context) {
	result, err := service.Captcha.Get(c.Cnf)
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	ctx.API.SetData(result)
}
