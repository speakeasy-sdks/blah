# Header
(*Header*)

### Available Operations

* [SendHeaders](#sendheaders) - Send Headers

## SendHeaders

Sends a single header params

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.Header.SendHeaders(ctx, operations.SendHeadersRequest{
        RequestBody: &operations.SendHeadersRequestBody{
            Value: "string",
        },
        CustomHeader: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.SendHeadersRequest](../../pkg/models/operations/sendheadersrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.SendHeadersResponse](../../pkg/models/operations/sendheadersresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
