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

package timeutil

import "time"

const (
	ISO8601Format = "2006-01-02T15:04:05Z"
)

func NowUTCSeconds() int64 { return time.Now().UTC().Unix() }

func FormatISO8601Date(timestampSecond int64) string {
	tm := time.Unix(timestampSecond, 0).UTC()
	return tm.Format(ISO8601Format)
}
