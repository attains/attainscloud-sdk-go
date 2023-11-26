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
	"context"
	"encoding/json"
	"fmt"
	"github.com/attains/attainscloud-sdk-go/core/errors"
	"github.com/attains/attainscloud-sdk-go/core/logger"
	"io/ioutil"
	"net/http"
	"reflect"
)

var (
	dynamicResponseStructs = []reflect.StructField{
		{
			Name: "Code",
			Type: reflect.TypeOf(int64(0)),
			Tag:  "json:\"code\"",
		},
		{
			Name: "Message",
			Type: reflect.TypeOf(""),
			Tag:  "json:\"message\"",
		},
	}
)

type AttainsResponse interface {
	SetResponse(*http.Response) AttainsResponse
	GetResponse() *http.Response
	WithLogger(logger.Interface) AttainsResponse
	ParseResponse(ctx context.Context) error
}

type DefaultAttainsResponse struct {
	response *http.Response
	logger   logger.Interface
	target   interface{}
}

func NewDefaultAttainsResponse(result interface{}) AttainsResponse {
	return &DefaultAttainsResponse{
		target: result,
	}
}

func (d *DefaultAttainsResponse) SetResponse(response *http.Response) AttainsResponse {
	d.response = response
	return d
}

func (d *DefaultAttainsResponse) GetResponse() *http.Response {
	return d.response
}

func (d *DefaultAttainsResponse) WithLogger(p logger.Interface) AttainsResponse {
	d.logger = p
	return d
}

func (d *DefaultAttainsResponse) ParseResponse(ctx context.Context) error {
	buf, err := ioutil.ReadAll(d.response.Body)
	d.logger.Debug(ctx, "Get raw response(%s),err(%v)", string(buf), err)
	if err != nil {
		return err
	}
	d.response.Body.Close()
	d.response.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	if vt := reflect.TypeOf(d.target); vt.Kind() != reflect.Ptr {
		return errors.NewAttainsClientError(fmt.Sprintf("result (%s) must be an pointer", vt.String()))
	}
	sfs := append(dynamicResponseStructs, reflect.StructField{
		Name: "Data",
		Type: reflect.TypeOf(d.target),
		Tag:  "json:\"data\"",
	})
	so := reflect.New(reflect.StructOf(sfs))
	if err := json.Unmarshal(buf, so.Interface()); err != nil {
		return err
	}
	if code := so.Elem().FieldByName("Code").Int(); code != int64(http.StatusOK) {
		message := so.Elem().FieldByName("Message").String()
		return errors.NewAttainsServiceError(code, message)
	}
	if em := so.Elem().FieldByName("Data"); !em.IsZero() {
		reflect.ValueOf(d.target).Elem().Set(em.Elem())
	}

	return nil
}
