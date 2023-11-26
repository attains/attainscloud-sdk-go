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

package errors

import "strconv"

type AttainsError interface {
	error
}

type AttainsClientError struct {
	message string
}

func (e *AttainsClientError) Error() string {
	return e.message
}

func NewAttainsClientError(text string) error {
	return &AttainsClientError{
		message: text,
	}
}

type AttainsServiceError struct {
	code    int64
	message string
}

func (e *AttainsServiceError) Error() string {
	ret := "[Code: " + strconv.FormatInt(e.code, 10)
	ret += "; Message: " + e.message + "]"
	return ret
}

func (e *AttainsServiceError) Code() int64 {
	return e.code
}

func NewAttainsServiceError(code int64, text string) error {
	return &AttainsServiceError{
		code:    code,
		message: text,
	}
}

const (
	ErrCodeRequestExpired = -2
)
