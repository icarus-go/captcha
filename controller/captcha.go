package controller

import (
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
	"pmo-test4.yz-intelligence.com/kit/captcha/service"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type captcha struct{}

var Captcha = new(captcha)

func (*captcha) Get(ctx *gins.Context) {
	var md *model.Captcha
	if err := ctx.ShouldBindJSON(&md); err != nil {
		ctx.API.SetError(err)
		return
	}
	result, err := service.Captcha.Get(*md)
	if err != nil {
		ctx.API.SetError(err)
		return
	}

	ctx.API.SetData(result)
}
