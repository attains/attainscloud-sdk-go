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

package config

import (
	"fmt"
	"github.com/attains/attainscloud-sdk-go/core/auth"
	"github.com/attains/attainscloud-sdk-go/core/logger"
	"github.com/attains/attainscloud-sdk-go/core/retry"
	"reflect"
	"runtime"
)

// Constants and default values

const (
	SdkVersion = "0.0.1"
)

var (
	DefaultUserAgent string
)

func init() {
	DefaultUserAgent = "cloud-sdk-go"
	DefaultUserAgent += "/" + SdkVersion
	DefaultUserAgent += "/" + runtime.Version()
	DefaultUserAgent += "/" + runtime.GOOS
	DefaultUserAgent += "/" + runtime.GOARCH
}

type AttainsCustomConfig struct {
	Endpoint    string
	UserAgent   string
	Credentials *auth.AttainsCredentials
	SignOption  *auth.SignOptions
	Retry       retry.AttainsRetryPolicy
	Logger      logger.Interface
}

type AttainsConfig struct {
	AttainsCustomConfig
	ProxyUrl                  string
	ConnectionTimeoutInMillis int
	RedirectDisabled          bool
}

func (c *AttainsConfig) String() string {
	return fmt.Sprintf(`AttainsConfig [
        Endpoint=%s;
        ProxyUrl=%s;
        UserAgent=%s;
        Credentials=%v;
        SignOption=%v;
        RetryPolicy=%v;
        Logger=%v;
        ConnectionTimeoutInMillis=%v;
		RedirectDisabled=%v
    ]`, c.Endpoint, c.ProxyUrl, c.UserAgent, c.Credentials,
		c.SignOption, reflect.TypeOf(c.Retry).Name(), reflect.TypeOf(c.Logger).Name(), c.ConnectionTimeoutInMillis, c.RedirectDisabled)
}
