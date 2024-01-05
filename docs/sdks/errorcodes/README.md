# ErrorCodes
(*ErrorCodes*)

### Available Operations

* [Get400](#get400) - Get400
* [Get401](#get401) - Get401
* [Get500](#get500) - Get500
* [Get501](#get501) - Get501
* [Catch412globalerror](#catch412globalerror) - catch 412 global error

## Get400

Get400

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
    res, err := s.ErrorCodes.Get400(ctx)
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


### Response

**[*operations.Get400Response](../../pkg/models/operations/get400response.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Get401

Get401

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
    res, err := s.ErrorCodes.Get401(ctx)
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


### Response

**[*operations.Get401Response](../../pkg/models/operations/get401response.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.LocalTestException   | 401,421,431,432,441            | application/json               |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Get500

Get500

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
    res, err := s.ErrorCodes.Get500(ctx)
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


### Response

**[*operations.Get500Response](../../pkg/models/operations/get500response.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Get501

Get501

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
    res, err := s.ErrorCodes.Get501(ctx)
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


### Response

**[*operations.Get501Response](../../pkg/models/operations/get501response.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412,501                        | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Catch412globalerror

catch 412 global error

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
    res, err := s.ErrorCodes.Catch412globalerror(ctx)
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


### Response

**[*operations.Catch412globalerrorResponse](../../pkg/models/operations/catch412globalerrorresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
