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

// Package auth signer.go - implement the specific sign algorithm of Attains Cloud V1 protocol
package auth

import (
	"fmt"
	"github.com/attains/attainscloud-sdk-go/core/logger"
	"github.com/attains/attainscloud-sdk-go/core/metadata"
	"github.com/attains/attainscloud-sdk-go/core/utils/cryptoutil"
	"github.com/attains/attainscloud-sdk-go/core/utils/strutil"
	"github.com/attains/attainscloud-sdk-go/core/utils/timeutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var (
	AttainsAuthVersion   = "attains-auth-v1"
	SignJoiner           = "\n"
	SignHeaderJoiner     = ";"
	DefaultExpireSeconds = 1800
	DefaultHeadersToSign = map[string]struct{}{
		strings.ToLower(metadata.RequestKeyHost):          {},
		strings.ToLower(metadata.RequestKeyContentLength): {},
		strings.ToLower(metadata.RequestKeyContentType):   {},
		strings.ToLower(metadata.RequestKeyContentMd5):    {},
	}
)

// Signer abstracts the entity that implements the `Sign` method
type Signer interface {
	// Sign the given Request with the Credentials and SignOptions
	Sign(*http.Request, logger.Interface, *AttainsCredentials, *SignOptions) error
}

// SignOptions defines the data structure used by Signer
type SignOptions struct {
	HeadersToSign map[string]struct{}
	Timestamp     int64
	ExpireSeconds int
}

func (opt *SignOptions) String() string {
	return fmt.Sprintf(`SignOptions [
        HeadersToSign=%s;
        Timestamp=%d;
        ExpireSeconds=%d
    ]`, opt.HeadersToSign, opt.Timestamp, opt.ExpireSeconds)
}

// AttainsV1Signer implements the v1 sign algorithm
type AttainsV1Signer struct{}

// Sign - generate the authorization string from the AttainsCredentials and SignOptions
//
// PARAMS:
//   - req: *http.Request for this sign
//   - cred: *AttainsCredentials to access the service
//   - opt: *SignOptions for this sign algorithm
func (a *AttainsV1Signer) Sign(req *http.Request, ll logger.Interface, cred *AttainsCredentials, opt *SignOptions) error {
	if req == nil {
		return fmt.Errorf("request should not be null for sign")
	}
	if cred == nil {
		return fmt.Errorf("credentials should not be null for sign")
	}
	if ll == nil {
		return fmt.Errorf("logger cannot be null")
	}

	// Prepare parameters
	accessKeyId := cred.AccessKeyId
	secretAccessKey := cred.SecretAccessKey
	signDate := timeutil.FormatISO8601Date(timeutil.NowUTCSeconds())
	// Modify the sign time if it is not the default value but specified by client
	if opt.Timestamp != 0 {
		signDate = timeutil.FormatISO8601Date(opt.Timestamp)
	}

	// Prepare the canonical request components
	signKeyInfo := fmt.Sprintf("%s/%s/%s/%d",
		AttainsAuthVersion,
		accessKeyId,
		signDate,
		opt.ExpireSeconds)
	ll.Debug(req.Context(), "signKeyInfo: %v", signKeyInfo)
	signKey := cryptoutil.HmacSha256Hex(secretAccessKey, signKeyInfo)

	ll.Debug(req.Context(), "signKey: %v", signKey)

	canonicalUri := getCanonicalURIPath(req.URL.Path)
	canonicalQueryString := getCanonicalQueryString(req.URL.Query())
	canonicalHeaders, signedHeadersArr := getCanonicalHeaders(req.Header, opt.HeadersToSign)

	// Generate signed headers string
	signedHeaders := ""
	if len(signedHeadersArr) > 0 {
		sort.Strings(signedHeadersArr)
		signedHeaders = strings.Join(signedHeadersArr, SignHeaderJoiner)
	}

	// Generate signature
	canonicalParts := []string{req.Method, canonicalUri, canonicalQueryString, canonicalHeaders}
	canonicalReq := strings.Join(canonicalParts, SignJoiner)
	ll.Debug(req.Context(), "CanonicalRequest data: %v\n", canonicalReq)
	signature := cryptoutil.HmacSha256Hex(signKey, canonicalReq)

	// Generate auth string and add to the reqeust header
	authStr := signKeyInfo + "/" + signedHeaders + "/" + signature
	ll.Info(req.Context(), "Authorization=%s", authStr)

	req.Header.Set(metadata.RequestKeyAuthorization, authStr)

	return nil
}

func getCanonicalURIPath(path string) string {
	if len(path) == 0 {
		return metadata.UriSeparator
	}
	canonicalPath := path
	if strings.HasPrefix(path, metadata.UriSeparator) {
		canonicalPath = path[1:]
	}
	canonicalPath = strutil.UriEncode(canonicalPath, false)
	return metadata.UriSeparator + canonicalPath
}

func getCanonicalQueryString(query url.Values) string {
	if len(query) == 0 {
		return ""
	}

	result := make([]string, 0)
	for k, vv := range query {
		if strings.ToLower(k) == strings.ToLower(metadata.RequestKeyAuthorization) {
			continue
		}
		for _, v := range vv {
			item := ""
			if len(v) == 0 {
				item = fmt.Sprintf("%s=", strutil.UriEncode(k, true))
			} else {
				item = fmt.Sprintf("%s=%s", strutil.UriEncode(k, true), strutil.UriEncode(v, true))
			}
			result = append(result, item)
		}
	}
	sort.Strings(result)
	return strings.Join(result, "&")
}

func getCanonicalHeaders(headers http.Header, headersToSign map[string]struct{}) (string, []string) {
	canonicalHeaders := make([]string, 0)
	signHeaders := make([]string, 0, len(headersToSign))
	for k, vv := range headers {
		headKey := strings.ToLower(k)
		if headKey == strings.ToLower(metadata.RequestKeyAuthorization) {
			continue
		}
		_, headExists := headersToSign[headKey]
		if headExists || (strings.HasPrefix(headKey, metadata.RequestKeyAttainsPrefix) && (headKey != metadata.RequestKeyAttainsRequestId)) {
			for _, v := range vv {
				headVal := strings.TrimSpace(v)
				encoded := strutil.UriEncode(headKey, true) + ":" + strutil.UriEncode(headVal, true)
				canonicalHeaders = append(canonicalHeaders, encoded)
				signHeaders = append(signHeaders, headKey)
			}
		}
	}
	sort.Strings(canonicalHeaders)
	sort.Strings(signHeaders)
	return strings.Join(canonicalHeaders, SignJoiner), signHeaders
}
