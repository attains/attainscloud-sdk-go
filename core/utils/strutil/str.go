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
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

func UriEncode(uri string, encodeSlash bool) string {
	var buf bytes.Buffer
	for _, b := range []byte(uri) {
		if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') ||
			b == '-' || b == '_' || b == '.' || b == '~' || (b == '/' && !encodeSlash) {
			buf.WriteByte(b)
		} else {
			buf.WriteString(fmt.Sprintf("%%%02X", b))
		}
	}
	return buf.String()
}

func CalculateContentMD5(data io.Reader, size int64) (string, error) {
	harsher := md5.New()
	n, err := io.CopyN(harsher, data, size)
	if err != nil {
		return "", fmt.Errorf("calculate content-md5 occurs error: %v", err)
	}
	if n != size {
		return "", fmt.Errorf("calculate content-md5 writing size %d != size %d", n, size)
	}
	return base64.StdEncoding.EncodeToString(harsher.Sum(nil)), nil
}

func NewUUID() string {
	var buf [16]byte
	for {
		if _, err := rand.Read(buf[:]); err == nil {
			break
		}
	}
	buf[6] = (buf[6] & 0x0f) | (4 << 4)
	buf[8] = (buf[8] & 0xbf) | 0x80

	res := make([]byte, 36)
	hex.Encode(res[0:8], buf[0:4])
	res[8] = '-'
	hex.Encode(res[9:13], buf[4:6])
	res[13] = '-'
	hex.Encode(res[14:18], buf[6:8])
	res[18] = '-'
	hex.Encode(res[19:23], buf[8:10])
	res[23] = '-'
	hex.Encode(res[24:], buf[10:])
	return string(res)
}

func NewRequestId() string {
	return NewUUID()
}
