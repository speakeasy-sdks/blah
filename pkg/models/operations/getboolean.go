// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type GetBooleanResponse struct {
	// HTTP response content type for this operation
	ContentType                   string
	GetBoolean200TextPlainBoolean *string
	// 500 Global
	GlobalTestException *shared.GlobalTestException
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetBooleanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetBooleanResponse) GetGetBoolean200TextPlainBoolean() *string {
	if o == nil {
		return nil
	}
	return o.GetBoolean200TextPlainBoolean
}

func (o *GetBooleanResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *GetBooleanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetBooleanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
