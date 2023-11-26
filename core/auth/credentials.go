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

// Package auth credentials.go - the credentials data structure definition
package auth

import "errors"

// AttainsCredentials define the data structure for authorization
type AttainsCredentials struct {
	AccessKeyId     string // access key id to the service
	SecretAccessKey string // secret access key to the service
}

func (a *AttainsCredentials) String() string {
	return "ak: " + a.AccessKeyId + ", sk: " + a.SecretAccessKey
}

func NewAttainsCredentials(ak, sk string) (*AttainsCredentials, error) {
	if len(ak) == 0 {
		return nil, errors.New("accessKeyId should not be empty")
	}
	if len(sk) == 0 {
		return nil, errors.New("secretKey should not be empty")
	}

	return &AttainsCredentials{ak, sk}, nil
}
