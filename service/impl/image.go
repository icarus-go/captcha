package impl

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"pmo-test4.yz-intelligence.com/kit/captcha/config"
	"pmo-test4.yz-intelligence.com/kit/captcha/ext"
	"pmo-test4.yz-intelligence.com/kit/component/gins"
)

var ImageStore base64Captcha.Store

type Image struct {
	Attribute *config.Attribute
}

func (i *Image) Limit(ctx *gins.Context) error {
	return nil
}

func (i *Image) Get(configuration *ext.Request) (*ext.Captcha, error) {
	if i.Attribute.Length < 4 {
		i.Attribute.Length = 4
	}

	if i.Attribute.CollectNumber < 1 {
		i.Attribute.CollectNumber = 10000
	}

	if i.Attribute.ExpireSecond < 1 {
		i.Attribute.ExpireSecond = 30
	}

	if ImageStore == nil {
		return nil, errors.New("图形验证码池为空")
	}

	driver := base64Captcha.NewDriverDigit(i.Attribute.Height, i.Attribute.Width, i.Attribute.Length, 0.7, 80) // 字符,公式,验证码配置, 生成默认数字的driver

	id, content, answer := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		return nil, err
	}

	if err = ImageStore.Set(id, answer); err != nil {
		return nil, err
	}

	//captchaID, imageBase64, err := cp.Generate()
	//if err != nil {
	//	return result, err
	//}
	//i.Store = &cp.Store

	return &ext.Captcha{
		CaptchaID: id,
		Email:     ext.Email{},
		Image:     ext.Image{ImageBase64: item.EncodeB64string()},
	}, nil
}

func (i *Image) Verify(code, captchaID string) bool {
	return ImageStore.Verify(captchaID, code, true)
}
