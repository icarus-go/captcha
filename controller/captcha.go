package controller

import (
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/service"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type captcha struct {
	Cnf config.Attribute
}

func NewCaptcha(cnf config.Attribute) *captcha {
	captcha := new(captcha)
	captcha.Cnf = cnf
	return captcha
}

//CommonGet
func (c *captcha) CommonGet(ctx *gins.Context) {
	code, err := service.New(&c.Cnf)
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	if err = code.Limit(ctx); err != nil {
		ctx.API.SetError(err)
		return
	}

	result, err := code.Get()
	if err != nil {
		ctx.API.SetError(err)
		return
	}
	ctx.API.SetData(result)
}
