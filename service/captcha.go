package service

import (
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
)

type captcha struct{}

var (
	store   = base64Captcha.DefaultMemStore // 验证码存储库
	Captcha = new(captcha)
)

//Get 获取验证码对象
func (c *captcha) Get(param config.Captcha) (*model.Captcha, error) {
	driver := base64Captcha.NewDriverDigit(param.ImgHeight, param.ImgWidth, param.KeyLong, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver

	cp := base64Captcha.NewCaptcha(driver, store)

	result := new(model.Captcha)

	captchaID, imageBase64, err := cp.Generate()
	if err != nil {
		return nil, err
	}

	result.Image = imageBase64
	result.CaptchaID = captchaID

	return result, nil
}

//Verify 校验
func (c *captcha) Verify(param model.Captcha) bool {
	return store.Verify(param.CaptchaID, param.Image, true)
}
