# QueryParam
(*QueryParam*)

### Available Operations

* [Date](#date) - Date
* [DateArray](#datearray) - Date Array
* [IntegerEnumArray](#integerenumarray) - Integer Enum Array
* [MultipleParams](#multipleparams) - MultipleParams
* [NoParams](#noparams) - NoParams
* [NumberArray](#numberarray) - Number Array
* [OptionalDynamicQueryParam](#optionaldynamicqueryparam) - Optional Dynamic Query Param
* [Rfc1123DateTime](#rfc1123datetime) - Rfc1123 DateTime
* [Rfc1123DateTimeArray](#rfc1123datetimearray) - Rfc1123 DateTime Array
* [Rfc3339DateTime](#rfc3339datetime) - Rfc3339 DateTime
* [Rfc3339DateTimeArray](#rfc3339datetimearray) - Rfc3339 DateTime Array
* [SimpleQuery](#simplequery) - SimpleQuery
* [StringArray](#stringarray) - String Array
* [StringEnumArray](#stringenumarray) - String Enum Array
* [StringParam](#stringparam) - StringParam
* [UnixDateTime](#unixdatetime) - Unix DateTime
* [UnixDateTimeArray](#unixdatetimearray) - Unix DateTime Array
* [URLParam](#urlparam) - UrlParam

## Date

Date

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.Date(ctx, operations.DateRequest{
        Date: types.MustDateFromString("2022-12-07"),
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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [operations.DateRequest](../../pkg/models/operations/daterequest.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.DateResponse](../../pkg/models/operations/dateresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## DateArray

Date Array

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.DateArray(ctx, operations.DateArrayRequest{
        Dates: []types.Date{
            types.MustDateFromString("2022-10-22"),
        },
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
| `request`                                                                      | [operations.DateArrayRequest](../../pkg/models/operations/datearrayrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.DateArrayResponse](../../pkg/models/operations/datearrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## IntegerEnumArray

Integer Enum Array

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.IntegerEnumArray(ctx, operations.IntegerEnumArrayRequest{
        Suites: []shared.SuiteCode{
            shared.SuiteCodeThree,
        },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.IntegerEnumArrayRequest](../../pkg/models/operations/integerenumarrayrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.IntegerEnumArrayResponse](../../pkg/models/operations/integerenumarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## MultipleParams

MultipleParams

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
    res, err := s.QueryParam.MultipleParams(ctx, operations.MultipleParamsRequest{
        Number: 907904,
        Precision: 9477.38,
        String: "string",
        URL: "https://definitive-tab.net",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.MultipleParamsRequest](../../pkg/models/operations/multipleparamsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.MultipleParamsResponse](../../pkg/models/operations/multipleparamsresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## NoParams

NoParams

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
    res, err := s.QueryParam.NoParams(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.NoParamsResponse](../../pkg/models/operations/noparamsresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## NumberArray

Number Array

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
    res, err := s.QueryParam.NumberArray(ctx, operations.NumberArrayRequest{
        Integers: []int{
            621871,
        },
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
| `request`                                                                          | [operations.NumberArrayRequest](../../pkg/models/operations/numberarrayrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.NumberArrayResponse](../../pkg/models/operations/numberarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## OptionalDynamicQueryParam

get optional dynamic query parameter

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
    res, err := s.QueryParam.OptionalDynamicQueryParam(ctx, operations.OptionalDynamicQueryParamRequest{
        Name: "string",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.OptionalDynamicQueryParamRequest](../../pkg/models/operations/optionaldynamicqueryparamrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.OptionalDynamicQueryParamResponse](../../pkg/models/operations/optionaldynamicqueryparamresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Rfc1123DateTime

Rfc1123 DateTime

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
    res, err := s.QueryParam.Rfc1123DateTime(ctx, operations.Rfc1123DateTimeRequest{
        Datetime: "string",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.Rfc1123DateTimeRequest](../../pkg/models/operations/rfc1123datetimerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.Rfc1123DateTimeResponse](../../pkg/models/operations/rfc1123datetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Rfc1123DateTimeArray

Rfc1123 DateTime Array

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
    res, err := s.QueryParam.Rfc1123DateTimeArray(ctx, operations.Rfc1123DateTimeArrayRequest{
        Datetimes: []string{
            "string",
        },
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.Rfc1123DateTimeArrayRequest](../../pkg/models/operations/rfc1123datetimearrayrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.Rfc1123DateTimeArrayResponse](../../pkg/models/operations/rfc1123datetimearrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Rfc3339DateTime

Rfc3339 DateTime

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.Rfc3339DateTime(ctx, operations.Rfc3339DateTimeRequest{
        Datetime: types.MustTimeFromString("2023-04-29T21:02:33.424Z"),
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.Rfc3339DateTimeRequest](../../pkg/models/operations/rfc3339datetimerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.Rfc3339DateTimeResponse](../../pkg/models/operations/rfc3339datetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## Rfc3339DateTimeArray

Rfc3339 DateTime Array

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"time"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.Rfc3339DateTimeArray(ctx, operations.Rfc3339DateTimeArrayRequest{
        Datetimes: []time.Time{
            types.MustTimeFromString("2022-09-26T17:09:47.534Z"),
        },
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.Rfc3339DateTimeArrayRequest](../../pkg/models/operations/rfc3339datetimearrayrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.Rfc3339DateTimeArrayResponse](../../pkg/models/operations/rfc3339datetimearrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SimpleQuery

SimpleQuery

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
    res, err := s.QueryParam.SimpleQuery(ctx, operations.SimpleQueryRequest{
        Boolean: false,
        Number: 41509,
        String: "string",
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
| `request`                                                                          | [operations.SimpleQueryRequest](../../pkg/models/operations/simplequeryrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.SimpleQueryResponse](../../pkg/models/operations/simplequeryresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## StringArray

String Array

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
    res, err := s.QueryParam.StringArray(ctx, operations.StringArrayRequest{
        Strings: []string{
            "string",
        },
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
| `request`                                                                          | [operations.StringArrayRequest](../../pkg/models/operations/stringarrayrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.StringArrayResponse](../../pkg/models/operations/stringarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## StringEnumArray

String Enum Array

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.StringEnumArray(ctx, operations.StringEnumArrayRequest{
        Days: []shared.Days{
            shared.DaysThursday,
        },
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.StringEnumArrayRequest](../../pkg/models/operations/stringenumarrayrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.StringEnumArrayResponse](../../pkg/models/operations/stringenumarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## StringParam

StringParam

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
    res, err := s.QueryParam.StringParam(ctx, operations.StringParamRequest{
        String: "string",
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
| `request`                                                                          | [operations.StringParamRequest](../../pkg/models/operations/stringparamrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.StringParamResponse](../../pkg/models/operations/stringparamresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UnixDateTime

Unix DateTime

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
    res, err := s.QueryParam.UnixDateTime(ctx, operations.UnixDateTimeRequest{
        Datetime: 8696.2,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UnixDateTimeRequest](../../pkg/models/operations/unixdatetimerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UnixDateTimeResponse](../../pkg/models/operations/unixdatetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UnixDateTimeArray

Unix DateTime Array

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
    res, err := s.QueryParam.UnixDateTimeArray(ctx, operations.UnixDateTimeArrayRequest{
        Datetimes: []float64{
            145.26,
        },
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UnixDateTimeArrayRequest](../../pkg/models/operations/unixdatetimearrayrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UnixDateTimeArrayResponse](../../pkg/models/operations/unixdatetimearrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## URLParam

UrlParam

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
    res, err := s.QueryParam.URLParam(ctx, operations.URLParamRequest{
        URL: "https://metallic-spirit.info",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.URLParamRequest](../../pkg/models/operations/urlparamrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.URLParamResponse](../../pkg/models/operations/urlparamresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
