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
	ExampleGetBalance()
}

// ExampleGetSignatureList Get the list of sms signatures example.
func ExampleGetSignatureList() {
	client := getSmsClient()
	result, err := client.GetSignatureList(context.Background(), &sms.GetSignatureListArgs{
		Signature:       "",
		SignatureStatus: "",
		StartAt:         0,
		EndAt:           0,
		Pagination: model.Pagination{
			PageIndex: 0,
		},
	})
	if err != nil {
		log.Panicln(err)
	}
	if result != nil {
		for _, detail := range result {
			fmt.Println(detail)
		}
	}
}

// ExampleQuerySignature Query a sms signature example.
func ExampleQuerySignature() {
	client := getSmsClient()
	result, err := client.QuerySignature(context.Background(), &sms.QuerySignatureArgs{
		Signature: "<your signature without 【 OR 】>",
	})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}

// ExampleSignatureApply Apply a sms signature example.
func ExampleSignatureApply() {
	client := getSmsClient()
	result, err := client.CreateSignature(context.Background(), &sms.CreateSignatureArgs{
		Signature:     "<your signature without 【 OR 】>",
		SignatureType: "Enterprise",
		Description:   "",
		CountryType:   "DOMESTIC",
	})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}

// ExampleGetTemplateList Get the list of sms templates example.
func ExampleGetTemplateList() {
	client := getSmsClient()
	result, err := client.GetTemplateList(context.Background(), &sms.GetTemplateListArgs{
		TemplateStatus: "",
		StartAt:        0,
		EndAt:          0,
		Pagination:     model.Pagination{},
	})
	if err != nil {
		log.Panicln(err)
	}
	if result != nil {
		for _, detail := range result {
			fmt.Println(detail)
		}
	}
}

// ExampleQueryTemplate Query a sms template example.
func ExampleQueryTemplate() {
	client := getSmsClient()
	result, err := client.QueryTemplate(context.Background(), &sms.QueryTemplateArgs{
		TemplateId: "<your sms template id>",
	})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}

// ExampleCreateTemplate Create a sms template example.
func ExampleCreateTemplate() {
	client := getSmsClient()

	result, err := client.CreateTemplate(context.Background(), &sms.CreateTemplateArgs{
		Name:        "<template name>",
		Content:     "<example: Dear user, your verification code is {code},{expiresIn} minutes valid.>",
		SmsType:     sms.SmsTypeCommonNotice,
		Description: "<my example template>",
		CountryType: "DOMESTIC",
	})
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}

// ExampleDeleteTemplate Delete a sms template example.
func ExampleDeleteTemplate() {
	client := getSmsClient()
	result, err := client.DeleteTemplate(context.Background(), &sms.DeleteTemplateArgs{
		TemplateId: "<your sms template id>",
	})
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("result:", result)
}

// ExampleSendSms Send sms example.
func ExampleSendSms() {
	client := getSmsClient()

	args := &sms.SendSmsArgs{
		TemplateId: "<your sms template id>",
		Signature:  "<your signature without 【 OR 】>",
		Items: []*sms.SendSmsArgsItem{
			&sms.SendSmsArgsItem{
				Mobile: "<your customer mobile number>",
				ContentVars: map[string]string{
					"<var1 name>": "<var1 value>",
					"<var2 name>": "<var2 value>",
				},
			},
		},
		SendAt: 0,
	}

	result, err := client.SendSms(context.Background(), args)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}

// ExampleGetBalance Get sms service balance example.
func ExampleGetBalance() {
	client := getSmsClient()
	result, err := client.GetBalance(context.Background())
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("result:", result)
}

func getSmsClient() *sms.SmsClient {
	ak, sk, endpoint := "<your-access-key>", "<your-secret-key>", "smsv1.bj.api.atains.cloud"
	attainsClient := httpclient.NewDefaultAttainsClient(ak, sk, endpoint)
	return sms.New(attainsClient)
}
