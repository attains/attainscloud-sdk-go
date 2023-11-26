# 成就云 Golang SDK

### [English Document](README.md)
#### [成就云官方网站](https://cloud.attains.cn?utm_source=github)

# 安装SDK工具包

## 运行环境

GO SDK可以在 `go1.13` 及以上环境下运行。

## 安装SDK

**直接从github下载**

使用`go get`工具从github进行下载：

```shell
go get github.com/attains/attainscloud-sdk-go/core
go get github.com/attains/attainscloud-sdk-go/services/sms
```

**SDK目录结构**

```text
attainscloud-sdk-go
|--core                     // 成就云 Golang SDK 核心包
|--services                 // 成就云 Golang SDK 所有服务
|  |--sms                   // 成就云 短信服务 Sms api
```

## 卸载SDK

预期卸载SDK时，删除下载的源码即可。

# 使用步骤

## 确认 Endpoint

在使用SDK之前，需确认您将接入的成就云产品的Endpoint（服务域名）。

## 创建Client对象

每种具体的服务都有一个`Client`对象，为开发者与对应的服务进行交互封装了一系列易用的方法。

## 调用功能接口

开发者基于创建的对应服务的`Client`对象，即可调用相应的功能接口，使用成就云产品的功能。

## 示例

下面以成就云短信服务（SMS）为例，给出一个基本的使用示例。

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

	// SMS服务的Endpoint
	ENDPOINT := "<domain-name>"

	// 创建SMS服务的Client
	attainsClient := httpclient.NewDefaultAttainsClient(ACCESS_KEY, SECRET_KEY, ENDPOINT)
	client := sms.New(attainsClient)
	
	result, err := client.GetBalance(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}
```

# 错误处理

GO语言以error类型标识错误，定义了如下两种错误类型：

错误类型        |  说明
----------------|-------------------
AttainsClientError  | 用户操作产生的错误
AttainsServiceError | 成就云服务返回的错误

用户使用SDK调用各服务的相关接口，除了返回所需的结果之外还会返回错误，用户可以获取相关错误的详细信息进行处理。实例如下：

```go
// smsClient 为已创建的SMS服务的Client对象
err := smsClient.SendSms(args)
if err != nil {
	switch realErr := err.(type) {
	case *attains.AttainsClientError:
		fmt.Println("client occurs error:", realErr.Error())
	case *attains.AttainsServiceError:
		fmt.Println("service occurs error:", realErr.Error())
	default:
		fmt.Println("unknown error:", err)
	}
}
fmt.Println("send sms succeed")
```

## 客户端异常

客户端异常表示客户端尝试向成就云服务发送请求以及数据传输时遇到的异常。例如，当发送请求时网络连接不可用时，则会返回AttainsClientError。

## 服务端异常

当服务端出现异常时，成就云服务端会返回给用户相应的错误信息，以便定位问题。每种服务端的异常需参考各服务的官网文档。

# 支持产品列表

产品名称   | 产品缩写 | 导入路径                                  |              文档 |                              
-----------|------|----------------------------------------------|-------------|
SMS service | SMS  | github.com/attains/attainscloud-sdk-go/services/sms | [doc](services/sms/README_zh-CN.md) |