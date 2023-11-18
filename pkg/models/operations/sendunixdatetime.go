// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type SendUnixDateTimeRequestBody struct {
	Datetime float64 `form:"name=datetime"`
}

func (o *SendUnixDateTimeRequestBody) GetDatetime() float64 {
	if o == nil {
		return 0.0
	}
	return o.Datetime
}

type SendUnixDateTimeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// 500 Global
	GlobalTestException *sdkerrors.GlobalTestException
	ServerResponse      *shared.ServerResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SendUnixDateTimeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendUnixDateTimeResponse) GetGlobalTestException() *sdkerrors.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *SendUnixDateTimeResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendUnixDateTimeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendUnixDateTimeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
