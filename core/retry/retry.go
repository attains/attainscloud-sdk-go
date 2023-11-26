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

package retry

import (
	"github.com/attains/attainscloud-sdk-go/core/errors"
	"net"
	"net/http"
	"time"
)

type AttainsRetryPolicy interface {
	ShouldRetry(errors.AttainsError, int) bool
	GetDelayBeforeNextRetryInMillis(errors.AttainsError, int) time.Duration
}

type AttainsNoRetryPolicy struct {
}

func (a AttainsNoRetryPolicy) ShouldRetry(errors.AttainsError, int) bool {
	return false
}

func (a AttainsNoRetryPolicy) GetDelayBeforeNextRetryInMillis(errors.AttainsError, int) time.Duration {
	return 0
}

func NewAttainsNoRetryPolicy() AttainsRetryPolicy {
	return &AttainsNoRetryPolicy{}
}

type AttainsBackoffRetryPolicy struct {
	maxErrorRetry        int
	maxDelayInMillis     int64
	baseIntervalInMillis int64
}

func (a AttainsBackoffRetryPolicy) ShouldRetry(err errors.AttainsError, attempts int) bool {
	// Do not retry any more when retry the max times
	if attempts >= a.maxErrorRetry {
		return false
	}

	if err == nil {
		return true
	}

	// Always retry on IO error
	if _, ok := err.(net.Error); ok {
		return true
	}

	// Only retry on a service error
	if realErr, ok := err.(*errors.AttainsServiceError); ok {
		code := realErr.Code()
		switch code {
		case http.StatusInternalServerError:
			return true
		case http.StatusBadGateway:
			return true
		case http.StatusServiceUnavailable:
			return true
		case http.StatusBadRequest:
			return true
		case errors.ErrCodeRequestExpired:
			return true
		default:
		}
	}
	return false
}

func (a AttainsBackoffRetryPolicy) GetDelayBeforeNextRetryInMillis(_ errors.AttainsError, attempts int) time.Duration {
	if attempts < 0 {
		return 0 * time.Millisecond
	}
	delayInMillis := (1 << uint64(attempts)) * a.baseIntervalInMillis
	if delayInMillis > a.maxDelayInMillis {
		return time.Duration(a.maxDelayInMillis) * time.Millisecond
	}
	return time.Duration(delayInMillis) * time.Millisecond
}

func NewAttainsBackoffRetryPolicy(maxRetry int, maxDelay, base int64) AttainsRetryPolicy {
	return &AttainsBackoffRetryPolicy{maxRetry, maxDelay, base}
}
