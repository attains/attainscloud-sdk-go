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
	"fmt"
	"github.com/attains/attainscloud-sdk-go/core/auth"
	"github.com/attains/attainscloud-sdk-go/core/config"
	"github.com/attains/attainscloud-sdk-go/core/errors"
	"github.com/attains/attainscloud-sdk-go/core/logger"
	"github.com/attains/attainscloud-sdk-go/core/metadata"
	"github.com/attains/attainscloud-sdk-go/core/retry"
	"github.com/attains/attainscloud-sdk-go/core/utils/strutil"
	"github.com/attains/attainscloud-sdk-go/core/utils/timeutil"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AttainsHttpClient interface {
	// SendRequest send a request
	SendRequest(AttainsRequest, AttainsResponse) error
	GetLogger() logger.Interface
}

type DefaultAttainsHttpClient struct {
	httpClient *http.Client
	transport  *http.Transport
	conf       *config.AttainsConfig
	signer     auth.Signer
	custom     bool
}

func newAttainsHttpClient(signer auth.Signer, conf *config.AttainsConfig, httpClient *http.Client, transport *http.Transport, custom bool) AttainsHttpClient {
	client := &DefaultAttainsHttpClient{
		httpClient: httpClient,
		transport:  transport,
		conf:       conf,
		signer:     signer,
		custom:     custom,
	}
	client.httpClient.Transport = client.transport
	return client
}

func NewCustomHttpClient(signer auth.Signer, conf *config.AttainsCustomConfig, httpClient *http.Client, transport *http.Transport) AttainsHttpClient {
	return newAttainsHttpClient(signer, &config.AttainsConfig{
		AttainsCustomConfig: *conf,
	}, httpClient, transport, true)
}

func NewAttainsHttpClient(signer auth.Signer, conf *config.AttainsConfig) AttainsHttpClient {
	httpClient := &http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	httpClient.Timeout = time.Duration(conf.ConnectionTimeoutInMillis/1000) * time.Second
	if conf.RedirectDisabled {
		httpClient.CheckRedirect = func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	transport := &http.Transport{
		Proxy: nil,
		DialContext: (&net.Dialer{
			Timeout:   metadata.DefaultDialTimeout,
			KeepAlive: metadata.DefaultKeepAliveTimeout,
		}).DialContext,
		DialTLS:                nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    10 * time.Second,
		DisableKeepAlives:      false,
		DisableCompression:     false,
		MaxIdleConns:           100,
		MaxIdleConnsPerHost:    metadata.DefaultMaxIdleConnsPerHost,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        90 * time.Second,
		ResponseHeaderTimeout:  metadata.DefaultResponseHeaderTimeout,
		ExpectContinueTimeout:  1 * time.Second,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      true,
	}

	return newAttainsHttpClient(signer, conf, httpClient, transport, false)
}

func NewDefaultAttainsClient(ak, sk string, endpoints string) AttainsHttpClient {
	conf := &config.AttainsConfig{
		AttainsCustomConfig: config.AttainsCustomConfig{
			Endpoint:  endpoints,
			UserAgent: config.DefaultUserAgent,
			Credentials: &auth.AttainsCredentials{
				AccessKeyId:     ak,
				SecretAccessKey: sk,
			},
			SignOption: &auth.SignOptions{
				HeadersToSign: auth.DefaultHeadersToSign,
				Timestamp:     0,
				ExpireSeconds: auth.DefaultExpireSeconds,
			},
			Retry:  retry.NewAttainsBackoffRetryPolicy(3, 20000, 300),
			Logger: logger.Default,
		},
		ProxyUrl:                  "",
		ConnectionTimeoutInMillis: metadata.DefaultConnectionTimeoutInMillis,
		RedirectDisabled:          false,
	}
	signer := &auth.AttainsV1Signer{}
	return NewAttainsHttpClient(signer, conf)
}

func (d *DefaultAttainsHttpClient) SendRequest(request AttainsRequest, response AttainsResponse) error {
	response.WithLogger(d.GetLogger())
	d.GetLogger().Debug(request.GetContext(), "Start send request")
	if !d.custom {
		if proxyUrl := request.GetProxyUrl(); proxyUrl != nil {
			d.transport.Proxy = func(request *http.Request) (*url.URL, error) {
				return proxyUrl, nil
			}
		} else {
			d.transport.Proxy = nil
		}
	}

	req := request.Build()

	endpoint := d.conf.Endpoint
	if endpoint == "" {
		endpoint = request.GetEndpoint()
	}

	if !strings.HasPrefix(endpoint, metadata.RequestProtocolHttps+"://") && !strings.HasPrefix(endpoint, metadata.RequestProtocolHttp+"://") {
		endpoint = metadata.RequestProtocolHttp + "://" + endpoint
	}
	u, _ := url.Parse(endpoint)
	if u.Scheme != metadata.RequestProtocolHttps && u.Scheme != metadata.RequestProtocolHttp {
		u.Scheme = metadata.RequestProtocolHttp
	}
	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host

	if strings.LastIndex(u.Host, ":") > strings.LastIndex(u.Host, "]") {
		u.Host = strings.TrimSuffix(u.Host, ":")
	}
	req.Host = u.Host
	req.Header.Set(metadata.RequestKeyHost, req.Host)

	d.GetLogger().Debug(request.GetContext(), "Request url: %s", req.URL)
	d.GetLogger().Debug(request.GetContext(), "Request host: %s", req.Host)

	if contentType := req.Header.Get(metadata.RequestKeyContentType); contentType == "" {
		req.Header.Set(metadata.RequestKeyContentType, metadata.DefaultContentType)
	}
	if userAgent := req.Header.Get(metadata.RequestKeyUserAgent); userAgent == "" {
		if d.conf.UserAgent != "" {
			req.Header.Set(metadata.RequestKeyUserAgent, d.conf.UserAgent)
		} else {
			req.Header.Set(metadata.RequestKeyUserAgent, config.DefaultUserAgent)
		}
	}
	req.Header.Set(metadata.RequestKeyAttainsDate, timeutil.FormatISO8601Date(timeutil.NowUTCSeconds()))

	requestId := request.GetRequestId()
	if len(requestId) == 0 {
		// Construct the request ID with UUID
		requestId = strutil.NewRequestId()
	}
	req.Header.Set(metadata.RequestKeyAttainsRequestId, requestId)

	d.GetLogger().Debug(request.GetContext(), "Request header: %v", req.Header)

	if req.Body != nil {
		_, contentMd5Exist := req.Header[metadata.RequestKeyContentMd5]
		_, contentLengthExist := req.Header[metadata.RequestKeyContentLength]
		if !contentMd5Exist || !contentLengthExist {
			body, err := ioutil.ReadAll(req.Body)
			defer req.Body.Close()
			if err != nil {
				return err
			}
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			buf := bytes.NewReader(body)
			size := buf.Len()
			if !contentMd5Exist {
				contentMd5, err := strutil.CalculateContentMD5(buf, int64(size))
				if err != nil {
					return err
				}
				req.Header.Set(metadata.RequestKeyContentMd5, contentMd5)
			}
			if !contentLengthExist {
				req.Header.Set(metadata.RequestKeyContentLength, fmt.Sprintf("%d", size))
			}
		}
	}

	signErr := d.signer.Sign(req, d.GetLogger(), d.conf.Credentials, d.conf.SignOption)
	if signErr != nil {
		return signErr
	}

	retries := 0
	if req.Body != nil {
		defer req.Body.Close() // Manually close the ReadCloser body for retry
	}

	for {
		d.GetLogger().Debug(request.GetContext(), "%dth try send request", retries)

		var retryBuf bytes.Buffer
		var teeReader io.Reader
		if d.conf.Retry.ShouldRetry(nil, 0) && req.Body != nil {
			teeReader = io.TeeReader(req.Body, &retryBuf)
			req.Body = ioutil.NopCloser(teeReader)
		}

		for s, i := range req.Header {
			fmt.Println(fmt.Sprintf("header(%s=%v)", s, i))
		}

		httpResponse, err := d.httpClient.Do(req)

		if err != nil {
			d.transport.CloseIdleConnections()
			return err
		}
		if httpResponse.StatusCode >= 400 && (req.Method == http.MethodPost || req.Method == http.MethodPut) {
			d.transport.CloseIdleConnections()
		}
		if err != nil {
			if d.conf.Retry.ShouldRetry(err, retries) {
				delayInMills := d.conf.Retry.GetDelayBeforeNextRetryInMillis(err, retries)
				time.Sleep(delayInMills)
			} else {
				return errors.NewAttainsClientError(fmt.Sprintf("execute http request failed! Retried %d times, error: %v", retries, err))
			}
			retries++
			if req.Body != nil {
				_, _ = ioutil.ReadAll(teeReader)
				req.Body = ioutil.NopCloser(&retryBuf)
			}
			continue
		}
		err = response.SetResponse(httpResponse).ParseResponse(request.GetContext())
		if err != nil {
			if serviceErr, ok := err.(*errors.AttainsServiceError); ok {
				if d.conf.Retry.ShouldRetry(serviceErr, retries) {
					delayInMills := d.conf.Retry.GetDelayBeforeNextRetryInMillis(serviceErr, retries)
					time.Sleep(delayInMills)
				} else {
					return serviceErr
				}
				retries++
				if req.Body != nil {
					_, _ = ioutil.ReadAll(teeReader)
					req.Body = ioutil.NopCloser(&retryBuf)
				}
				continue
			}
			return err
		}

		return nil
	}
}

func (d *DefaultAttainsHttpClient) GetLogger() logger.Interface {
	if d.conf.Logger == nil {
		return logger.Default
	}
	return d.conf.Logger
}
