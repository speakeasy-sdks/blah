// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type SenddeleteFormResponse struct {
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

func (o *SenddeleteFormResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SenddeleteFormResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *SenddeleteFormResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SenddeleteFormResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SenddeleteFormResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
