// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type SendMixedArrayRequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

func (o *SendMixedArrayRequestBodyFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *SendMixedArrayRequestBodyFile) GetFile() string {
	if o == nil {
		return ""
	}
	return o.File
}

type SendMixedArrayRequestBody struct {
	File     SendMixedArrayRequestBodyFile `multipartForm:"file"`
	Integers []int                         `multipartForm:"name=integers"`
	Models   []shared.Employee             `multipartForm:"name=models,json"`
	Strings  []string                      `multipartForm:"name=strings"`
}

func (o *SendMixedArrayRequestBody) GetFile() SendMixedArrayRequestBodyFile {
	if o == nil {
		return SendMixedArrayRequestBodyFile{}
	}
	return o.File
}

func (o *SendMixedArrayRequestBody) GetIntegers() []int {
	if o == nil {
		return []int{}
	}
	return o.Integers
}

func (o *SendMixedArrayRequestBody) GetModels() []shared.Employee {
	if o == nil {
		return []shared.Employee{}
	}
	return o.Models
}

func (o *SendMixedArrayRequestBody) GetStrings() []string {
	if o == nil {
		return []string{}
	}
	return o.Strings
}

type SendMixedArrayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// 500 Global
	GlobalTestException *shared.GlobalTestException
	ServerResponse      *shared.ServerResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SendMixedArrayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendMixedArrayResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *SendMixedArrayResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendMixedArrayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendMixedArrayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
