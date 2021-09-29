// Package ext
// 因为swagger的问题所以必须要使用不同的包名
package ext

//Request 请求
type Request struct {
	Email EmailGenerate `json:"email"` // 邮件所需参数
	SMS   SmsGenerator  `json:"sms"`   // 短信需要手机号码
	Image struct{}      `json:"image"` // 图片不需要配置
}

//EmailGenerate 邮箱验证码
type EmailGenerate struct {
	ServerHost string   `json:"serverHost" form:"serverHost"` // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerPort int      `json:"serverPort" form:"serverPort"` // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	FromEmail  string   `json:"fromEmail" form:"fromEmail"`   // FromEmail　发件人邮箱地址
	FromPasswd string   `json:"fromPasswd" form:"fromPasswd"` //发件人邮箱密码（注意，这里是明文形式)
	Recipient  []string `json:"recipient" form:"recipient"`   //收件人邮箱
	CC         []string `json:"cc" form:"cc"`                 //抄送

}

//SmsGenerator 短信验证码参数
type SmsGenerator struct {
	Mobile string `json:"mobile" form:"mobile"` // 手机号码
}
