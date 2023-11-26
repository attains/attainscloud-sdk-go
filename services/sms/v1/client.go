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

package v1

import (
	"context"
	"encoding/json"
	"github.com/attains/attainscloud-sdk-go/core/httpclient"
	"github.com/attains/attainscloud-sdk-go/core/metadata"
	"github.com/attains/attainscloud-sdk-go/core/utils/structutil"
	"net/http"
)

// SmsClient sms service client
type SmsClient struct {
	acHttpClient httpclient.AttainsHttpClient
}

// New Create a sms service client
func New(client httpclient.AttainsHttpClient) *SmsClient {
	return &SmsClient{
		acHttpClient: client,
	}
}

func (s *SmsClient) newRequest(ctx context.Context) httpclient.AttainsRequest {
	return httpclient.NewDefaultAttainsRequest(ctx, nil).WithEndpoint(DefaultEndpoint)
}

func (s *SmsClient) sendRequest(q httpclient.AttainsRequest, r interface{}) error {
	return s.acHttpClient.SendRequest(q, httpclient.NewDefaultAttainsResponse(r))
}

// CreateSignature Create s sms signature.
func (s *SmsClient) CreateSignature(ctx context.Context, args *CreateSignatureArgs) (*CreateSignatureResult, error) {
	var err error
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	q := s.newRequest(ctx).
		WithPath(RequestUriSignatureApply).
		WithMethod(http.MethodPost).
		WithBodyBytes(body)
	r := new(CreateSignatureResult)
	err = s.sendRequest(q, r)
	return r, err
}

// QuerySignature Query a sms signature.
func (s *SmsClient) QuerySignature(ctx context.Context, args *QuerySignatureArgs) (*QuerySignatureResult, error) {
	var err error
	query, err := structutil.StructToUrlValues(args, true)
	if err != nil {
		return nil, err
	}
	q := s.newRequest(ctx).
		WithPath(RequestUriSignatureQuery).
		WithQuery(query)
	r := new(QuerySignatureResult)
	err = s.sendRequest(q, r)
	return r, err
}

// GetSignatureList Get the list of sms signature.
func (s *SmsClient) GetSignatureList(ctx context.Context, args *GetSignatureListArgs) (GetSignatureListResult, error) {
	var err error
	query, err := structutil.StructToUrlValues(args, true)
	if err != nil {
		return nil, err
	}
	q := s.newRequest(ctx).
		WithPath(RequestUriSignatureList).
		WithQuery(query)
	r := &GetSignatureListResult{}
	err = s.sendRequest(q, r)
	return *r, err
}

// CreateTemplate create a sms template.
func (s *SmsClient) CreateTemplate(ctx context.Context, args *CreateTemplateArgs) (*CreateTemplateResult, error) {
	var err error
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	q := s.newRequest(ctx).
		WithPath(RequestUriTemplateCreate).
		WithMethod(http.MethodPost).
		WithBodyBytes(body)
	r := new(CreateTemplateResult)
	err = s.sendRequest(q, r)
	return r, err
}

// QueryTemplate query a sms template.
func (s *SmsClient) QueryTemplate(ctx context.Context, args *QueryTemplateArgs) (*QueryTemplateResult, error) {
	q := s.newRequest(ctx).
		WithPath(RequestUriTemplateQuery + metadata.UriSeparator + args.TemplateId)
	r := new(QueryTemplateResult)
	err := s.sendRequest(q, r)
	return r, err
}

// GetTemplateList get the list of sms template.
func (s *SmsClient) GetTemplateList(ctx context.Context, args *GetTemplateListArgs) (GetTemplateListResult, error) {
	var err error
	query, err := structutil.StructToUrlValues(args, true)
	if err != nil {
		return nil, err
	}
	q := s.newRequest(ctx).
		WithPath(RequestUriTemplateList).
		WithQuery(query)
	r := new(GetTemplateListResult)
	err = s.sendRequest(q, r)
	return *r, err
}

// DeleteTemplate delete a sms template.
func (s *SmsClient) DeleteTemplate(ctx context.Context, args *DeleteTemplateArgs) (*DeleteTemplateResult, error) {
	q := s.newRequest(ctx).
		WithPath(RequestUriTemplateDelete + metadata.UriSeparator + args.TemplateId).
		WithMethod(http.MethodDelete)
	r := new(DeleteTemplateResult)
	err := s.sendRequest(q, r)
	return r, err
}

// SendSms send sms.
func (s *SmsClient) SendSms(ctx context.Context, args *SendSmsArgs) (*SendSmsResult, error) {
	var err error
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	q := s.newRequest(ctx).
		WithPath(RequestUriSendSms).
		WithMethod(http.MethodPost).
		WithBodyBytes(body)
	r := new(SendSmsResult)
	err = s.sendRequest(q, r)
	return r, err
}

// GetBalance get sms account balance.
func (s *SmsClient) GetBalance(ctx context.Context) (int64, error) {
	q := s.newRequest(ctx).WithPath(RequestUriGetBalance)
	r := int64(0)
	err := s.sendRequest(q, &r)
	return r, err
}
