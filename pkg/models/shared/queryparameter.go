// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// QueryParameter - Query parameter key value pair echoed back from the server.
type QueryParameter struct {
	Key *string `json:"key,omitempty"`
}

func (o *QueryParameter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}