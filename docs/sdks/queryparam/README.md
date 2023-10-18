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
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.Date(ctx, operations.DateRequest{
        Date: types.MustDateFromString("2021-12-06"),
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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [operations.DateRequest](../../models/operations/daterequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.DateResponse](../../models/operations/dateresponse.md), error**


## DateArray

Date Array

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.DateArray(ctx, operations.DateArrayRequest{
        Dates: []types.Date{
            types.MustDateFromString("2021-10-21"),
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.DateArrayRequest](../../models/operations/datearrayrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.DateArrayResponse](../../models/operations/datearrayresponse.md), error**


## IntegerEnumArray

Integer Enum Array

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.IntegerEnumArrayRequest](../../models/operations/integerenumarrayrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.IntegerEnumArrayResponse](../../models/operations/integerenumarrayresponse.md), error**


## MultipleParams

MultipleParams

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
    res, err := s.QueryParam.MultipleParams(ctx, operations.MultipleParamsRequest{
        Number: 907904,
        Precision: 9477.38,
        String: "Cadillac",
        URL: "https://showy-macro.biz",
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
| `request`                                                                            | [operations.MultipleParamsRequest](../../models/operations/multipleparamsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.MultipleParamsResponse](../../models/operations/multipleparamsresponse.md), error**


## NoParams

NoParams

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
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

**[*operations.NoParamsResponse](../../models/operations/noparamsresponse.md), error**


## NumberArray

Number Array

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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.NumberArrayRequest](../../models/operations/numberarrayrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.NumberArrayResponse](../../models/operations/numberarrayresponse.md), error**


## OptionalDynamicQueryParam

get optional dynamic query parameter

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
    res, err := s.QueryParam.OptionalDynamicQueryParam(ctx, operations.OptionalDynamicQueryParamRequest{
        Name: "Diesel",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.OptionalDynamicQueryParamRequest](../../models/operations/optionaldynamicqueryparamrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.OptionalDynamicQueryParamResponse](../../models/operations/optionaldynamicqueryparamresponse.md), error**


## Rfc1123DateTime

Rfc1123 DateTime

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
    res, err := s.QueryParam.Rfc1123DateTime(ctx, operations.Rfc1123DateTimeRequest{
        Datetime: "Tricycle",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.Rfc1123DateTimeRequest](../../models/operations/rfc1123datetimerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.Rfc1123DateTimeResponse](../../models/operations/rfc1123datetimeresponse.md), error**


## Rfc1123DateTimeArray

Rfc1123 DateTime Array

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
    res, err := s.QueryParam.Rfc1123DateTimeArray(ctx, operations.Rfc1123DateTimeArrayRequest{
        Datetimes: []string{
            "Games",
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.Rfc1123DateTimeArrayRequest](../../models/operations/rfc1123datetimearrayrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.Rfc1123DateTimeArrayResponse](../../models/operations/rfc1123datetimearrayresponse.md), error**


## Rfc3339DateTime

Rfc3339 DateTime

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.Rfc3339DateTime(ctx, operations.Rfc3339DateTimeRequest{
        Datetime: types.MustTimeFromString("2022-04-29T10:26:48.393Z"),
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.Rfc3339DateTimeRequest](../../models/operations/rfc3339datetimerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.Rfc3339DateTimeResponse](../../models/operations/rfc3339datetimeresponse.md), error**


## Rfc3339DateTimeArray

Rfc3339 DateTime Array

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.QueryParam.Rfc3339DateTimeArray(ctx, operations.Rfc3339DateTimeArrayRequest{
        Datetimes: []time.Time{
            types.MustTimeFromString("2021-09-26T11:16:44.151Z"),
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.Rfc3339DateTimeArrayRequest](../../models/operations/rfc3339datetimearrayrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.Rfc3339DateTimeArrayResponse](../../models/operations/rfc3339datetimearrayresponse.md), error**


## SimpleQuery

SimpleQuery

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
    res, err := s.QueryParam.SimpleQuery(ctx, operations.SimpleQueryRequest{
        Boolean: false,
        Number: 41509,
        String: "Rap",
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
| `request`                                                                      | [operations.SimpleQueryRequest](../../models/operations/simplequeryrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.SimpleQueryResponse](../../models/operations/simplequeryresponse.md), error**


## StringArray

String Array

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
    res, err := s.QueryParam.StringArray(ctx, operations.StringArrayRequest{
        Strings: []string{
            "Nicaragua",
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
| `request`                                                                      | [operations.StringArrayRequest](../../models/operations/stringarrayrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.StringArrayResponse](../../models/operations/stringarrayresponse.md), error**


## StringEnumArray

String Enum Array

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.StringEnumArrayRequest](../../models/operations/stringenumarrayrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.StringEnumArrayResponse](../../models/operations/stringenumarrayresponse.md), error**


## StringParam

StringParam

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
    res, err := s.QueryParam.StringParam(ctx, operations.StringParamRequest{
        String: "Southwest",
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
| `request`                                                                      | [operations.StringParamRequest](../../models/operations/stringparamrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.StringParamResponse](../../models/operations/stringparamresponse.md), error**


## UnixDateTime

Unix DateTime

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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UnixDateTimeRequest](../../models/operations/unixdatetimerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UnixDateTimeResponse](../../models/operations/unixdatetimeresponse.md), error**


## UnixDateTimeArray

Unix DateTime Array

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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UnixDateTimeArrayRequest](../../models/operations/unixdatetimearrayrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UnixDateTimeArrayResponse](../../models/operations/unixdatetimearrayresponse.md), error**


## URLParam

UrlParam

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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.URLParamRequest](../../models/operations/urlparamrequest.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[*operations.URLParamResponse](../../models/operations/urlparamresponse.md), error**

