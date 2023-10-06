// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type EchoResponseBody struct {
}

// EchoResponse - Raw http Request info
type EchoResponse struct {
	Body    map[string]EchoResponseBody `json:"body,omitempty"`
	Headers map[string]string           `json:"headers,omitempty"`
	Method  *string                     `json:"method,omitempty"`
	// relativePath
	Path        *string                   `json:"path,omitempty"`
	Query       map[string]QueryParameter `json:"query,omitempty"`
	UploadCount *int                      `json:"uploadCount,omitempty"`
}

func (o *EchoResponse) GetBody() map[string]EchoResponseBody {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *EchoResponse) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *EchoResponse) GetMethod() *string {
	if o == nil {
		return nil
	}
	return o.Method
}

func (o *EchoResponse) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *EchoResponse) GetQuery() map[string]QueryParameter {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *EchoResponse) GetUploadCount() *int {
	if o == nil {
		return nil
	}
	return o.UploadCount
}