// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type Get400Response struct {
	// HTTP response content type for this operation
	ContentType              string
	Get400200TextPlainObject *string
	// 500 Global
	GlobalTestException *shared.GlobalTestException
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *Get400Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Get400Response) GetGet400200TextPlainObject() *string {
	if o == nil {
		return nil
	}
	return o.Get400200TextPlainObject
}

func (o *Get400Response) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *Get400Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Get400Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
