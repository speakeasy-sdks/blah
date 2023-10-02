// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type SendHeadersRequestBody struct {
	// Represents the value of the custom header
	Value string `form:"name=value"`
}

func (o *SendHeadersRequestBody) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type SendHeadersRequest struct {
	RequestBody  *SendHeadersRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
	CustomHeader string                  `header:"style=simple,explode=false,name=custom-header"`
}

func (o *SendHeadersRequest) GetRequestBody() *SendHeadersRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *SendHeadersRequest) GetCustomHeader() string {
	if o == nil {
		return ""
	}
	return o.CustomHeader
}

type SendHeadersResponse struct {
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

func (o *SendHeadersResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendHeadersResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *SendHeadersResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendHeadersResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendHeadersResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
