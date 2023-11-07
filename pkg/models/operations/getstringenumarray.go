// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"net/http"
)

type Type string

const (
	TypeString Type = "string"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "string":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type GetStringEnumArrayRequest struct {
	Array bool `queryParam:"style=form,explode=true,name=array"`
	Type  Type `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetStringEnumArrayRequest) GetArray() bool {
	if o == nil {
		return false
	}
	return o.Array
}

func (o *GetStringEnumArrayRequest) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

type GetStringEnumArrayResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// 500 Global
	GlobalTestException *shared.GlobalTestException
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Enums       []shared.Days
}

func (o *GetStringEnumArrayResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetStringEnumArrayResponse) GetGlobalTestException() *shared.GlobalTestException {
	if o == nil {
		return nil
	}
	return o.GlobalTestException
}

func (o *GetStringEnumArrayResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetStringEnumArrayResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetStringEnumArrayResponse) GetEnums() []shared.Days {
	if o == nil {
		return nil
	}
	return o.Enums
}
