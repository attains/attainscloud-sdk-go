# 成就云 Golang SDK 核心包

### [English Document](README.md)

# 安装SDK工具包

## 运行环境

GO SDK可以在 `go1.13` 及以上环境下运行。

## 安装SDK

**直接从github下载**

使用`go get`工具从github进行下载：

```shell
go get github.com/attains/attainscloud-sdk-go/core
```

**SDK目录结构**

```text
core
|--auth                     // 签名和权限认证
|--config                   // 配置
|--errors                   // 错误
|--httpclient               // http客户端
|--logger                   // 日志
|--metadata                 // 预定义资源
|--model                    // 预定义model
|--retry                    // 重试策略
|--utils                    // 公用的工具实现
```

## 卸载SDK

预期卸载SDK时，删除下载的源码即可。

## 用法

首先获取 `httpclient` 实例

```go
package main

import (
	"github.com/attains/attainscloud-sdk-go/core/httpclient"
)

func main() {
	ak, sk, region := "<您的 access-key>", "<您的 secret-key>", "<服务地域>"
	attainsClient := httpclient.NewDefaultAttainsClient(ak, sk, fmt.Sprintf("smsv1.%s.api.attains.cloud", region))
}
```

