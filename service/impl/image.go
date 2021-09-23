package impl

import (
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/model"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

type Image struct {
	Attribute *config.Attribute
	Store     base64Captcha.Store
}

func (i *Image) Limit(ctx *gins.Context) error {
	return nil
}

func (i *Image) Get() (model.Captcha, error) {
	driver := base64Captcha.NewDriverDigit(i.Attribute.Height, i.Attribute.Width, i.Attribute.Length, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver

	cp := base64Captcha.NewCaptcha(driver, i.Store)

	result := model.Captcha{}

	captchaID, imageBase64, err := cp.Generate()
	if err != nil {
		return result, err
	}

	result.Image = imageBase64
	result.CaptchaID = captchaID
	return result, nil
}

func (i *Image) Verify(code, captchaID string) bool {
	return i.Store.Verify(captchaID, code, true)
}
