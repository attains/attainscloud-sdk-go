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

package httpclient

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type AttainsRequest interface {
	WithEndpoint(endpoint string) AttainsRequest
	GetEndpoint() string
	WithRequestId(string) AttainsRequest
	GetRequestId() string
	WithMethod(string) AttainsRequest
	WithPath(string) AttainsRequest
	WithQuery(url.Values) AttainsRequest
	WithHeader(http.Header) AttainsRequest
	WithContentType(string) AttainsRequest
	WithBody(io.ReadCloser) AttainsRequest
	WithBodyBytes([]byte) AttainsRequest
	WithProxyUrl(*url.URL) AttainsRequest
	GetProxyUrl() *url.URL
	GetContext() context.Context
	Build() *http.Request
}

type DefaultAttainsRequest struct {
	request   *http.Request
	proxyURL  *url.URL
	endpoint  string
	requestId string
}

func NewDefaultAttainsRequest(ctx context.Context, r *http.Request) AttainsRequest {
	request, _ := http.NewRequestWithContext(ctx, http.MethodGet, "", nil)
	if r != nil {
		request = r
	}
	return &DefaultAttainsRequest{
		request: request,
	}
}

func (d *DefaultAttainsRequest) WithMethod(s string) AttainsRequest {
	d.request.Method = s
	return d
}

func (d *DefaultAttainsRequest) WithEndpoint(endpoint string) AttainsRequest {
	d.endpoint = endpoint
	return d
}

func (d *DefaultAttainsRequest) GetEndpoint() string {
	return d.endpoint
}

func (d *DefaultAttainsRequest) WithPath(s string) AttainsRequest {
	d.request.URL.Path = s
	if idx := strings.Index(s, "?"); idx > -1 {
		d.request.URL.Path = s[:idx]
		query, _ := url.ParseQuery(s[idx:])
		d.request.URL.RawQuery = query.Encode()
	}
	return d
}

func (d *DefaultAttainsRequest) WithQuery(values url.Values) AttainsRequest {
	if values != nil {
		d.request.URL.RawQuery = values.Encode()
	}
	return d
}

func (d *DefaultAttainsRequest) WithHeader(header http.Header) AttainsRequest {
	if header != nil {
		for k := range header {
			if _, exist := d.request.Header[k]; exist {
				for _, s := range header[k] {
					d.request.Header.Add(k, s)
				}
			} else {
				for i, s := range header[k] {
					if i == 0 {
						d.request.Header.Set(k, s)
						continue
					}
					d.request.Header.Add(k, s)
				}
			}
		}
	}
	return d
}

func (d *DefaultAttainsRequest) WithContentType(s string) AttainsRequest {
	d.request.Header.Set("Content-Type", s)
	return d
}

func (d *DefaultAttainsRequest) WithRequestId(requestId string) AttainsRequest {
	d.requestId = requestId
	return d
}

func (d *DefaultAttainsRequest) GetRequestId() string {
	return d.requestId
}

func (d *DefaultAttainsRequest) WithBody(closer io.ReadCloser) AttainsRequest {
	d.request.Body = closer
	return d
}

func (d *DefaultAttainsRequest) WithBodyBytes(stream []byte) AttainsRequest {
	d.request.Body = ioutil.NopCloser(bytes.NewReader(stream))
	return d
}

func (d *DefaultAttainsRequest) WithProxyUrl(u *url.URL) AttainsRequest {
	d.proxyURL = u
	return d
}

func (d *DefaultAttainsRequest) GetProxyUrl() *url.URL {
	return d.proxyURL
}

func (d *DefaultAttainsRequest) GetContext() context.Context {
	return d.request.Context()
}

func (d *DefaultAttainsRequest) Build() *http.Request {
	return d.request
}
