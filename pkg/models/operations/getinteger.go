// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type GetIntegerResponse struct {
	// HTTP response content type for this operation
	ContentType                        string
	GetInteger200TextPlainInt32Integer *string
	// 500 Global
	GlobalTestException *shared.GlobalTestException
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetIntegerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetIntegerResponse) GetGetInteger200TextPlainInt32Integer() *string {
	if o == nil {
		return nil
	}
	return o.GetInteger200TextPlainInt32Integer
}

func (o *GetIntegerResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *GetIntegerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetIntegerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
