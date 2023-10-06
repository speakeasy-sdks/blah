# ResponseTypes
(*ResponseTypes*)

### Available Operations

* [Get1123DateTime](#get1123datetime) - Get 1123DateTime
* [Get3339Datetime](#get3339datetime) - Get 3339Datetime
* [GetBinary](#getbinary) - Get Binary
* [GetBoolean](#getboolean) - Get Boolean
* [GetDateArray](#getdatearray) - Get Date Array
* [GetDynamic](#getdynamic) - Get Dynamic
* [GetHeaders](#getheaders) - Get Headers
* [GetInteger](#getinteger) - Get Integer
* [GetLong](#getlong) - Get Long
* [GetModelArray](#getmodelarray) - Get Model Array
* [GetPrecision](#getprecision) - Get Precision
* [GetStringEnumArray](#getstringenumarray) - Get String Enum Array
* [GetUnixDateTime](#getunixdatetime) - Get UnixDateTime
* [Getcontenttypeheaders](#getcontenttypeheaders) - get content type headers
* [Returnresponsewithenums](#returnresponsewithenums) - return response with enums

## Get1123DateTime

Get 1123DateTime

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
    res, err := s.ResponseTypes.Get1123DateTime(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Get1123DateTime200TextPlainDateTimeRfc1123String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.Get1123DateTimeResponse](../../models/operations/get1123datetimeresponse.md), error**


## Get3339Datetime

Get 3339Datetime

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
    res, err := s.ResponseTypes.Get3339Datetime(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Get3339Datetime200TextPlainDateTimeString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.Get3339DatetimeResponse](../../models/operations/get3339datetimeresponse.md), error**


## GetBinary

gets a binary object

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
    res, err := s.ResponseTypes.GetBinary(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetBinary200ApplicationOctetStreamBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetBinaryResponse](../../models/operations/getbinaryresponse.md), error**


## GetBoolean

Get Boolean

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
    res, err := s.ResponseTypes.GetBoolean(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetBoolean200TextPlainBoolean != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetBooleanResponse](../../models/operations/getbooleanresponse.md), error**


## GetDateArray

Get Date Array

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
    res, err := s.ResponseTypes.GetDateArray(ctx, operations.GetDateArrayRequest{
        Array: false,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDateArray200ApplicationJSONDateStrings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetDateArrayRequest](../../models/operations/getdatearrayrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetDateArrayResponse](../../models/operations/getdatearrayresponse.md), error**


## GetDynamic

Get Dynamic

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
    res, err := s.ResponseTypes.GetDynamic(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDynamic200TextPlainObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetDynamicResponse](../../models/operations/getdynamicresponse.md), error**


## GetHeaders

Get Headers

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
    res, err := s.ResponseTypes.GetHeaders(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetHeadersResponse](../../models/operations/getheadersresponse.md), error**


## GetInteger

Gets a integer response

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
    res, err := s.ResponseTypes.GetInteger(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetInteger200TextPlainInt32Integer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetIntegerResponse](../../models/operations/getintegerresponse.md), error**


## GetLong

Get Long

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
    res, err := s.ResponseTypes.GetLong(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetLong200TextPlainInt64Integer != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |


### Response

**[*operations.GetLongResponse](../../models/operations/getlongresponse.md), error**


## GetModelArray

Get Model Array

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
    res, err := s.ResponseTypes.GetModelArray(ctx, operations.GetModelArrayRequest{
        Array: false,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.People != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetModelArrayRequest](../../models/operations/getmodelarrayrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetModelArrayResponse](../../models/operations/getmodelarrayresponse.md), error**


## GetPrecision

Get Precision

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
    res, err := s.ResponseTypes.GetPrecision(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetPrecision200TextPlainDoubleNumber != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetPrecisionResponse](../../models/operations/getprecisionresponse.md), error**


## GetStringEnumArray

Get String Enum Array

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
    res, err := s.ResponseTypes.GetStringEnumArray(ctx, operations.GetStringEnumArrayRequest{
        Array: false,
        Type: operations.GetStringEnumArrayTypeString,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Days != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetStringEnumArrayRequest](../../models/operations/getstringenumarrayrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetStringEnumArrayResponse](../../models/operations/getstringenumarrayresponse.md), error**


## GetUnixDateTime

Get UnixDateTime

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
    res, err := s.ResponseTypes.GetUnixDateTime(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetUnixDateTime200TextPlainUnixTimestampNumber != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetUnixDateTimeResponse](../../models/operations/getunixdatetimeresponse.md), error**


## Getcontenttypeheaders

get content type headers

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
    res, err := s.ResponseTypes.Getcontenttypeheaders(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetcontenttypeheadersResponse](../../models/operations/getcontenttypeheadersresponse.md), error**


## Returnresponsewithenums

return response with enums

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
    res, err := s.ResponseTypes.Returnresponsewithenums(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ResponsewithEnum != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.ReturnresponsewithenumsResponse](../../models/operations/returnresponsewithenumsresponse.md), error**

