# Header
(*.Header*)

### Available Operations

* [SendHeaders](#sendheaders) - Send Headers

## SendHeaders

Sends a single header params

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.SendHeadersRequest](../../models/operations/sendheadersrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.SendHeadersResponse](../../models/operations/sendheadersresponse.md), error**

