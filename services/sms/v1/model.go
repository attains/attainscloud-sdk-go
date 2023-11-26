/*
 * Copyright 2023 Attains Cloud, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 *
 * visit: https://cloud.attains.cn
 *
 */

// Field Definition：
//    Signature:         signature content. Excluding【 and 】,max length 18 bytes.
//    SignatureType:     signature type. one of "COMMON"、"NOTICE"、"PROMOTION".
//    SignatureStatus:   one of "CHECKING"、"PASSED"、"REJECTED".
//    Description:       the description for this signature.
//    CountryType:       current only support DOMESTIC.
//    Uid:               the user id of applicant for this signature.
//    Review:            review comments for this signature.
//    StartAt:           unix timestamp.
//    EndAt:			 unix timestamp.

// Package v1 model.go
package v1

import (
	"github.com/attains/attainscloud-sdk-go/core/model"
)

// ==========================// Signature ========================== //

// CreateSignatureArgs The args for create a sms signature.
type CreateSignatureArgs struct {
	Signature     string `json:"signature"`
	SignatureType string `json:"signatureType"`
	Description   string `json:"description"`
	CountryType   string `json:"countryType"`
}

// CreateSignatureResult The result for create a sms signature, void.
type CreateSignatureResult model.EmptyS

// SignatureDetail The detail of a sms a signature.
type SignatureDetail struct {
	SignatureStatus string `json:"signatureStatus"`
	Uid             string `json:"uid"`
	Signature       string `json:"signature"`
	CountryType     string `json:"countryType"`
	Review          string `json:"review"`
	CreatedAt       string `json:"createdAt"`
	AuditAt         string `json:"auditAt"`
}

// QuerySignatureArgs The args for query a sms signature.
type QuerySignatureArgs struct {
	Signature string `json:"signature"`
}

// QuerySignatureResult The result for query a sms signature.
type QuerySignatureResult SignatureDetail

// GetSignatureListArgs The args for get list of sms signature.
type GetSignatureListArgs struct {
	Signature       string `json:"signature"`
	SignatureStatus string `json:"signatureStatus"`
	StartAt         int    `json:"startAt"`
	EndAt           int    `json:"endAt"`
	model.Pagination
}

// GetSignatureListResult The result for get list of sms signature.
type GetSignatureListResult []*SignatureDetail

// ========================== Signature //========================== //

// ==========================// Template ========================== //

// CreateTemplateArgs The args for create a sms template
type CreateTemplateArgs struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	SmsType     string `json:"smsType"`
	Description string `json:"description"`
	CountryType string `json:"countryType"`
}

// CreateTemplateResult The result for create a sms template
type CreateTemplateResult struct {
	TemplateId     string `json:"templateId"`
	TemplateStatus string `json:"templateStatus"`
}

// QueryTemplateDetail The detail of a sms template
type QueryTemplateDetail struct {
	TemplateId     string `json:"templateId"`
	TemplateStatus string `json:"templateStatus"`
	Uid            string `json:"uid"`
	Name           string `json:"name"`
	Content        string `json:"content"`
	SmsType        string `json:"smsType"`
	Description    string `json:"description"`
	CountryType    string `json:"countryType"`
	Review         string `json:"review"`
	CreatedAt      string `json:"createdAt"`
	AuditAt        string `json:"auditAt"`
}

// QueryTemplateArgs The args for query a sms template
type QueryTemplateArgs struct {
	TemplateId string `json:"templateId"`
}

// QueryTemplateResult The result for query a sms template
type QueryTemplateResult QueryTemplateDetail

// GetTemplateListArgs The args for get list of sms template
type GetTemplateListArgs struct {
	TemplateStatus string `json:"templateStatus"`
	StartAt        int    `json:"startAt"`
	EndAt          int    `json:"endAt"`
	model.Pagination
}

// GetTemplateListResult The result for get list of sms template
type GetTemplateListResult []*QueryTemplateDetail

// DeleteTemplateArgs The args for delete a sms template
type DeleteTemplateArgs struct {
	TemplateId string `json:"templateId"`
}

// DeleteTemplateResult The result for delete a sms template
type DeleteTemplateResult model.EmptyS

// ========================== Template //========================== //

// ==========================// SendSms ========================== //

// SendSmsArgsItem the detail for send every sms
type SendSmsArgsItem struct {
	Mobile      string            `json:"mobile"`
	ContentVars map[string]string `json:"contentVars"`
}

// SendSmsArgs The args for send sms
type SendSmsArgs struct {
	TemplateId string             `json:"templateId"`
	Signature  string             `json:"signature"`
	Items      []*SendSmsArgsItem `json:"items"`
	SendAt     int64              `json:"sendAt"`
}

// SendSmsResult The result for send sms
type SendSmsResult model.EmptyS

// ========================== SendSms //========================== //
