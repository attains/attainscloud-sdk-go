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

const (
	DefaultEndpoint = "smsv1.bj.api.attains.cloud"
)

const (
	RequestUriSignatureApply = "/signature/apply"
	RequestUriSignatureQuery = "/signature/query"
	RequestUriSignatureList  = "/signature/list"

	RequestUriTemplateCreate = "/template/create"
	RequestUriTemplateQuery  = "/template/query"
	RequestUriTemplateDelete = "/template/delete"
	RequestUriTemplateList   = "/template/list"

	RequestUriSendSms    = "/send"
	RequestUriGetBalance = "/balance"
)

const (
	SignatureStatusChecking = "CHECKING"
	SignatureStatusPassed   = "PASSED"
	SignatureStatusRejected = "REJECTED"
)

const (
	TemplateStatusChecking = "CHECKING"
	TemplateStatusPassed   = "PASSED"
	TemplateStatusRejected = "REJECTED"
)

const (
	SmsTypeCommonNotice = "CommonNotice"
	SmsTypeCommonSale   = "CommonSale"
)
