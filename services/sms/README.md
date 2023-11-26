# Attains Cloud SMS Service SDK

### [中文文档](README_zh-CN.md)

## Install SDK

**Download directly from GitHub**

Download from GitHub using the 'go get' tool：

```shell
go get github.com/attains/attainscloud-sdk-go/services/sms
```

## Example

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
	// Your Access Key and Secret Key 
	ACCESS_KEY, SECRET_KEY := "<your-access-key>", "<your-access-key>"

	// Endpoint of SMS service
	ENDPOINT := "<domain-name>"

	// Create a client for SMS services 
	attainsClient := httpclient.NewDefaultAttainsClient(ACCESS_KEY, SECRET_KEY, ENDPOINT)
	client := sms.New(attainsClient)
	
	result, err := client.GetBalance(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}
```

# Support API list

SmsClient method        |  Illustrate
----------------        |-------------------
New                     | Create a sms service client object
CreateSignature         | Create s sms signature
QuerySignature          | Query a sms signature.
GetSignatureList        | Get the list of sms signature.
CreateTemplate          | create a sms template
QueryTemplate           | query a sms template.
GetTemplateList         | get the list of sms template.
DeleteTemplate          | delete a sms template.
SendSms                 | send sms.
GetBalance              | get sms account balance.