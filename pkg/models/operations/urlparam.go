// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type URLParamRequest struct {
	URL string `queryParam:"style=form,explode=true,name=url"`
}

func (o *URLParamRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type URLParamResponse struct {
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

func (o *URLParamResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *URLParamResponse) GetGlobalTestException() *sdkerrors.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *URLParamResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *URLParamResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *URLParamResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
