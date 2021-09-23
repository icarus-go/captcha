package model

type Configuration struct {
	Email Email // 邮件所需参数

	SMS Sms // 短信需要手机号码

	Image struct{} // 图片不需要配置
}

type Email struct {
	ServerHost string // ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerPort int    // ServerPort 邮箱服务器端口，如腾讯企业邮箱为465

	FromEmail  string // FromEmail　发件人邮箱地址
	FromPasswd string //发件人邮箱密码（注意，这里是明文形式)

	Recipient []string //收件人邮箱
	CC        []string //抄送

}

//Sms 短信验证码参数
type Sms struct {
	Mobile string `json:"mobile"`
}
