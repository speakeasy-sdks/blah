# Echo
(*Echo*)

### Available Operations

* [Jsonecho](#jsonecho) - Json echo
* [QueryEcho](#queryecho) - QueryEcho

## Jsonecho

Echo's back the request

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.Echo.Jsonecho(ctx, []byte("0x4CcAf4eebe"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.JsonechoResponse](../../pkg/models/operations/jsonechoresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## QueryEcho

QueryEcho

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.Echo.QueryEcho(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EchoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.QueryEchoResponse](../../pkg/models/operations/queryechoresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
