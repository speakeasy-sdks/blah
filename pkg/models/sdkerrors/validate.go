// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type Validate struct {
	Address *string `json:"address,omitempty"`
	Field   string  `json:"field"`
	Name    string  `json:"name"`
}

var _ error = &Validate{}

func (e *Validate) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
