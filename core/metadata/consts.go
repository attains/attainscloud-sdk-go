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

package metadata

import "time"

const (
	UriSeparator = "/"
)

const (
	DefaultMaxIdleConnsPerHost       = 500
	DefaultResponseHeaderTimeout     = 60 * time.Second
	DefaultDialTimeout               = 30 * time.Second
	DefaultKeepAliveTimeout          = 30 * time.Second
	DefaultContentType               = "application/json;charset=utf-8"
	DefaultConnectionTimeoutInMillis = 1200 * 1000
)

const (
	RequestProtocolHttps    = "https"
	RequestProtocolHttp     = "http"
	RequestKeyAuthorization = "Authorization"
	RequestKeyContentLength = "Content-Length"
	RequestKeyContentMd5    = "Content-Md5"
	RequestKeyContentType   = "Content-Type"
	RequestKeyHost          = "HOST"
	RequestKeyUserAgent     = "User-Agent"

	RequestKeyAttainsPrefix    = "x-attains-"
	RequestKeyAttainsRequestId = "x-attains-request-id"
	RequestKeyAttainsDate      = "x-attains-date"
)
