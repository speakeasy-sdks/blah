// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/utils"
	"net/http"
)

type SendStringEnumArrayRequestBody struct {
	Days []shared.Days `form:"name=days"`
}

func (o *SendStringEnumArrayRequestBody) GetDays() []shared.Days {
	if o == nil {
		return []shared.Days{}
	}
	return o.Days
}

type SendStringEnumArrayRequest struct {
	RequestBody *SendStringEnumArrayRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
	array       bool                            `const:"true" queryParam:"style=form,explode=true,name=array"`
}

func (s SendStringEnumArrayRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SendStringEnumArrayRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SendStringEnumArrayRequest) GetRequestBody() *SendStringEnumArrayRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *SendStringEnumArrayRequest) GetArray() bool {
	return true
}

type SendStringEnumArrayResponse struct {
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

func (o *SendStringEnumArrayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendStringEnumArrayResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *SendStringEnumArrayResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendStringEnumArrayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendStringEnumArrayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
