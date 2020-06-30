# Yuntongxun SMS SDK for GO

[容联云通讯](https://www.yuntongxun.com) SDK

## Quick Start

```go
go get -u github.com/cloopen/go-sms-sdk/cloopen
```

```go
package main

import (
	"github.com/cloopen/go-sms-sdk/cloopen"
	"log"
)

func main() {
	cfg := cloopen.DefaultConfig().
		// 开发者主账号
		WithAPIAccount("xxxxxxxxxxxxx").
		// 主账号令牌 TOKEN
		WithAPIToken("xxxxxxxxxxxxx")
	sms := cloopen.NewJsonClient(cfg).SMS()
  // 下发包体参数
	input := &cloopen.SendRequest{
		// 应用的APPID
		AppId: "xxxxxxxxxxxxx",
		// 手机号码
		To: "18601312882",
		// 模版ID
		TemplateId: "8696",
		// 模版变量内容 非必填
		Datas: []string{"您的验证码是4059"},
	}
  // 下发
	resp, err := sms.Send(input)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Response MsgId: %s \n", resp.TemplateSMS.SmsMessageSid)

}

```

## 使用说明

* 自定义配置及默认

  `WithAPIAccount(xxx)` 配置主账号   **需调用者初始化此值**

  `WithAPIToken(xxx)` 配置主账号令牌  **需调用者初始化此值**

  `WithSmsHost(xxx)` 配置ip:port    **默认 app.cloopen.com:8883**

  `WithUseSSL(true)` 配置是否使用https  **默认启用https**

  `WithHTTPClient(customHttp)` 配置自定义httpClient  **默认使用sdk封装的httpClient**

  `WithHttpConf(&HttpConf{...})` 配置sdk封装的httpClient可调整参数 **默认使用sdk封装的httpClient参数**

  **参考HttpConf默认配置：**

  ```go
  // 时间单位为毫秒
  &HttpConf{
  			Timeout:             300,
  			KeepAlive:           30000,
  			MaxIdleConns:        100,
  			IdleConnTimeout:     30000,
  			TLSHandshakeTimeout: 300,
  		}
  ```

* 方法调用

  `cloopen.NewJsonClient(cfg)`  json 格式包体使用此方法

  `cloopen.NewXmlClient(cfg)`    xml  格式包体使用此方法

## 源码说明

- sdk
  - config.go 接口基础配置
  - client.go  客户端定义、配置
  - fields.go 常量定义
  - sms.go 短信功能
  - util.go 工具函数
- 分支说明
  - master最新稳定发布版本
  - develop待发布版本，贡献的代码请pull request到这里:)