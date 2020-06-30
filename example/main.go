package main

import (
	"cloopen"
	"log"
)

func main() {
	cfg := cloopen.DefaultConfig().
		// 开发者主账号
		WithAPIAccount("ff8081813fc65581013fc72b94880000").
		// 主账号令牌 TOKEN
		WithAPIToken("4e44a775d79e422a9ee26e2966d2cb66").
		WithSmsHost("192.168.182.100:4200").WithUseSSL(false)
	sms := cloopen.NewJsonClient(cfg).SMS()
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "ff8080816f65d70901705b53e14100de",
		// 手机号码
		To: "18601322882",
		// 模版ID
		TemplateId: "8696",
		// 模版变量内容 非必填
		Datas: []string{"您的验证码是4059"},
	}

	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)

}
