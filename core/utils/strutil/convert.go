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

package strutil

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// ToStringFunc try to convert value to string, return error on fail
type ToStringFunc func(v interface{}) (string, error)

var ErrConvType = errors.New("convert value type error")

// SprintToStrFunc convert any value to string by fmt.Sprint
var SprintToStrFunc = func(v interface{}) (string, error) {
	if v == nil {
		return "", nil
	}
	return fmt.Sprint(v), nil
}

// ToStringWithFunc convert value to string, with a func to fallback handle.
//
// On not convert:
//   - If fbFn is nil, will return ErrConvType.
//   - If fbFn is not nil, will call it to convert.
func ToStringWithFunc(val interface{}, fbFn ToStringFunc) (str string, err error) {
	switch value := val.(type) {
	case int:
		str = strconv.Itoa(value)
	case int8:
		str = strconv.Itoa(int(value))
	case int16:
		str = strconv.Itoa(int(value))
	case int32: // same as `rune`
		str = strconv.Itoa(int(value))
	case int64:
		str = strconv.FormatInt(value, 10)
	case uint:
		str = strconv.FormatUint(uint64(value), 10)
	case uint8:
		str = strconv.FormatUint(uint64(value), 10)
	case uint16:
		str = strconv.FormatUint(uint64(value), 10)
	case uint32:
		str = strconv.FormatUint(uint64(value), 10)
	case uint64:
		str = strconv.FormatUint(value, 10)
	case float32:
		str = strconv.FormatFloat(float64(value), 'f', -1, 32)
	case float64:
		str = strconv.FormatFloat(value, 'f', -1, 64)
	case bool:
		str = strconv.FormatBool(value)
	case string:
		str = value
	case []byte:
		str = string(value)
	case time.Duration:
		str = strconv.FormatInt(int64(value), 10)
	case fmt.Stringer:
		str = value.String()
	case error:
		str = value.Error()
	default:
		if fbFn == nil {
			err = ErrConvType
		} else {
			str, err = fbFn(value)
		}
	}
	return
}

// SafeString convert value to string, will ignore error
func SafeString(in interface{}) string {
	val, _ := ToStringWithFunc(in, SprintToStrFunc)
	return val
}
