# AttainsCloud SDK Core Package For Golang.

### [中文文档](README_zh-CN.md)


# Install SDK Toolkit

## Run environments

The GO SDK can run in environments of 'go1.13' and above.

## Install SDK

**Download directly from GitHub**

Download from GitHub using the 'go get' tool：

```shell
go get github.com/attains/attainscloud-sdk-go/core
```

**SDK directory structure**

```text
core
|--auth                     // Signature and permission authentication
|--config                   // Configuration
|--errors                   // Error
|--httpclient               // http client
|--logger                   // logger package
|--metadata                 // Predefined Resources
|--model                    // Predefined model
|--retry                    // Retry policy
|--utils                    // Common tool implementation
```

## Uninstalling SDK

When uninstalling the SDK, simply delete the downloaded source code.


## Usage

```go
package main

import (
	"fmt"
	"github.com/attains/attainscloud-sdk-go/core/httpclient"
)

func main() {
	ak, sk, region := "<your-access-key>", "<your-secret-key>", "<service-region>"
	attainsClient := httpclient.NewDefaultAttainsClient(ak, sk, fmt.Sprintf("smsv1.%s.api.attains.cloud", region))
}
```