# 成就云短信服务 SDK

### [English Document](README.md)

## 安装SDK

**直接从github下载**

使用`go get`工具从github进行下载：

```shell
go get github.com/attains/attainscloud-sdk-go/services/sms
```

## 示例

```go
package main

import (
    "context"
    "fmt"
    "github.com/attains/attainscloud-sdk-go/core/httpclient"
    "github.com/attains/attainscloud-sdk-go/core/model"
    sms "github.com/attains/attainscloud-sdk-go/services/sms/v1"
    "log"
)

func main() {
	// 您的 Access Key 和 Secret Key 
	ACCESS_KEY, SECRET_KEY := "<your-access-key>", "<your-access-key>"

	// SMS Service 的 endpoint
	ENDPOINT := "<domain-name>"

	// 创建一个 smsClient
	attainsClient := httpclient.NewDefaultAttainsClient(ACCESS_KEY, SECRET_KEY, ENDPOINT)
	client := sms.New(attainsClient)
	
	result, err := client.GetBalance(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}
```

# 支持的 api 列表

SmsClient 方法           | 说明  
----------------        |-------------------
New                     | 创建一个sms客户端对象
CreateSignature         | 创建(申请)一个短信签名
QuerySignature          | 查询一个短信签名
GetSignatureList        | 获取短信签名列表
CreateTemplate          | 创建一个短信模板
QueryTemplate           | 查询一个短信模板
GetTemplateList         | 获取短信模板列表
DeleteTemplate          | 删除一个短信模板
SendSms                 | 发送短信
GetBalance              | 获取短信账户余额