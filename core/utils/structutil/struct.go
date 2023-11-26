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

package structutil

import (
	"encoding/json"
	"github.com/attains/attainscloud-sdk-go/core/utils/strutil"
	"net/url"
)

func StructToUrlMap(s interface{}) (map[string]interface{}, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	v := map[string]interface{}{}
	if err = json.Unmarshal(j, &v); err != nil {
		return nil, err
	}
	return v, nil
}

func StructToUrlValues(s interface{}, ignoreEmpty bool) (url.Values, error) {
	m, err := StructToUrlMap(s)
	if err != nil {
		return nil, err
	}
	u := url.Values{}
	if m != nil {
		for k, v := range m {
			if v == "" && ignoreEmpty {
				continue
			}
			u.Set(k, strutil.SafeString(v))
		}
	}
	return u, nil
}
