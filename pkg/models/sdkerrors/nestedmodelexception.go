// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type NestedModelException struct {
	ServerCode    string   `json:"ServerCode"`
	ServerMessage string   `json:"ServerMessage"`
	Model         Validate `json:"model"`
}

var _ error = &NestedModelException{}

func (e *NestedModelException) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
