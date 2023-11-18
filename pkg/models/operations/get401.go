// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/sdkerrors"
	"net/http"
)

type Get401Response struct {
	// HTTP response content type for this operation
	ContentType string
	// 401 Local
	LocalTestException *sdkerrors.LocalTestException
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Res         *string
}

func (o *Get401Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Get401Response) GetLocalTestException() *sdkerrors.LocalTestException {
	if o == nil {
		return nil
	}
	return o.LocalTestException
}

func (o *Get401Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Get401Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *Get401Response) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
