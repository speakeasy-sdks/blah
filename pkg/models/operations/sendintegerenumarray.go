// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/utils"
	"net/http"
)

type SendIntegerEnumArrayRequestBody struct {
	Suites []shared.SuiteCode `form:"name=suites"`
}

func (o *SendIntegerEnumArrayRequestBody) GetSuites() []shared.SuiteCode {
	if o == nil {
		return []shared.SuiteCode{}
	}
	return o.Suites
}

type SendIntegerEnumArrayRequest struct {
	RequestBody *SendIntegerEnumArrayRequestBody `request:"mediaType=application/x-www-form-urlencoded"`
	array       bool                             `const:"true" queryParam:"style=form,explode=true,name=array"`
}

func (s SendIntegerEnumArrayRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SendIntegerEnumArrayRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SendIntegerEnumArrayRequest) GetRequestBody() *SendIntegerEnumArrayRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *SendIntegerEnumArrayRequest) GetArray() bool {
	return true
}

type SendIntegerEnumArrayResponse struct {
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

func (o *SendIntegerEnumArrayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SendIntegerEnumArrayResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *SendIntegerEnumArrayResponse) GetServerResponse() *shared.ServerResponse {
	if o == nil {
		return nil
	}
	return o.ServerResponse
}

func (o *SendIntegerEnumArrayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SendIntegerEnumArrayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
