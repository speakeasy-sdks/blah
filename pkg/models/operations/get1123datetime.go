// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type Get1123DateTimeResponse struct {
	// HTTP response content type for this operation
	ContentType                                      string
	Get1123DateTime200TextPlainDateTimeRfc1123String *string
	// 500 Global
	GlobalTestException *shared.GlobalTestException
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *Get1123DateTimeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Get1123DateTimeResponse) GetGet1123DateTime200TextPlainDateTimeRfc1123String() *string {
	if o == nil {
		return nil
	}
	return o.Get1123DateTime200TextPlainDateTimeRfc1123String
}

func (o *Get1123DateTimeResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *Get1123DateTimeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Get1123DateTimeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
