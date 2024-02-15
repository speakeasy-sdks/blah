# BodyParams
(*BodyParams*)

### Available Operations

* [PostSendDate](#postsenddate) - Send Date
* [PostSendIntegerArray](#postsendintegerarray) - Send Integer Array
* [PostSendIntegerEnumArray](#postsendintegerenumarray) - SendIntegerEnumArray
* [PostSendModel](#postsendmodel) - Send Model
* [PostSendRfc1123DateTime](#postsendrfc1123datetime) - Send Rfc1123 DateTime
* [PostSendRfc3339DateTime](#postsendrfc3339datetime) - Send Rfc3339 DateTime
* [PostSendStringArray](#postsendstringarray) - Send String Array
* [PostSendStringEnumArray](#postsendstringenumarray) - SendStringEnumArray
* [PostSendUnixDateTime](#postsendunixdatetime) - Send UnixDateTime
* [SendDeleteBodywithModel](#senddeletebodywithmodel) - Send Delete Body with Model
* [SendDeletePlainText](#senddeleteplaintext) - Send Delete PlainText
* [SendDynamic](#senddynamic) - Send Dynamic
* [UpdateModel](#updatemodel) - Update Model
* [SendDeleteBody](#senddeletebody) - send Delete Body
* [UpdateString](#updatestring) - update String

## PostSendDate

Send Date

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
    res, err := s.BodyParams.PostSendDate(ctx, []byte("0x9EFb4f7198"))
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
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendDateResponse](../../pkg/models/operations/postsenddateresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendIntegerArray

Send Integer Array

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
    res, err := s.BodyParams.PostSendIntegerArray(ctx, operations.PostSendIntegerArrayRequest{
        RequestBody: []int{
            520832,
        },
        Array: false,
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
| `request`                                                                                            | [operations.PostSendIntegerArrayRequest](../../pkg/models/operations/postsendintegerarrayrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.PostSendIntegerArrayResponse](../../pkg/models/operations/postsendintegerarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendIntegerEnumArray

SendIntegerEnumArray

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
    res, err := s.BodyParams.PostSendIntegerEnumArray(ctx, operations.PostSendIntegerEnumArrayRequest{
        RequestBody: []shared.SuiteCode{
            shared.SuiteCodeTwo,
        },
        Array: false,
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.PostSendIntegerEnumArrayRequest](../../pkg/models/operations/postsendintegerenumarrayrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.PostSendIntegerEnumArrayResponse](../../pkg/models/operations/postsendintegerenumarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendModel

Send Model

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.BodyParams.PostSendModel(ctx, shared.Employee{
        Address: "4584 Bailey Ways",
        Age: 85493,
        Birthday: types.MustDateFromString("2022-07-04"),
        Birthtime: types.MustTimeFromString("2022-02-27T02:56:38.456Z"),
        Department: "<value>",
        Dependents: []shared.Person{
            shared.Person{
                Address: "881 Ryann Ways",
                Age: 46325,
                Birthday: types.MustDateFromString("2023-11-16"),
                Birthtime: types.MustTimeFromString("2024-02-10T21:17:33.741Z"),
                Name: "<value>",
                PersonType: "<value>",
                UID: "<value>",
            },
        },
        HiredAt: "<value>",
        JoiningDay: shared.DaysWednesday,
        Name: "<value>",
        PersonType: "<value>",
        Salary: 383188,
        UID: "<value>",
        WorkingDays: []shared.Days{
            shared.DaysFriday,
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

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.Employee](../../pkg/models/shared/employee.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.PostSendModelResponse](../../pkg/models/operations/postsendmodelresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendRfc1123DateTime

Send Rfc1123 DateTime

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
    res, err := s.BodyParams.PostSendRfc1123DateTime(ctx, "<value>")
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
| `request`                                             | [string](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendRfc1123DateTimeResponse](../../pkg/models/operations/postsendrfc1123datetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendRfc3339DateTime

Send Rfc3339 DateTime

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
    res, err := s.BodyParams.PostSendRfc3339DateTime(ctx, []byte("0xf54ad30feB"))
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
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendRfc3339DateTimeResponse](../../pkg/models/operations/postsendrfc3339datetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendStringArray

sends a string body param

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
    res, err := s.BodyParams.PostSendStringArray(ctx, operations.PostSendStringArrayRequest{
        RequestBody: []string{
            "<value>",
        },
        Array: false,
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostSendStringArrayRequest](../../pkg/models/operations/postsendstringarrayrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PostSendStringArrayResponse](../../pkg/models/operations/postsendstringarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendStringEnumArray

SendStringEnumArray

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
    res, err := s.BodyParams.PostSendStringEnumArray(ctx, operations.PostSendStringEnumArrayRequest{
        RequestBody: []shared.Days{
            shared.DaysSunday,
        },
        Array: false,
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
| `request`                                                                                                  | [operations.PostSendStringEnumArrayRequest](../../pkg/models/operations/postsendstringenumarrayrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PostSendStringEnumArrayResponse](../../pkg/models/operations/postsendstringenumarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## PostSendUnixDateTime

Send UnixDateTime

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
    res, err := s.BodyParams.PostSendUnixDateTime(ctx, []byte("0x396818C72D"))
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
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendUnixDateTimeResponse](../../pkg/models/operations/postsendunixdatetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendDeleteBodywithModel

Send Delete Body with Model

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
    res, err := s.BodyParams.SendDeleteBodywithModel(ctx)
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

**[*operations.SendDeleteBodywithModelResponse](../../pkg/models/operations/senddeletebodywithmodelresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendDeletePlainText

Send Delete PlainText

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
    res, err := s.BodyParams.SendDeletePlainText(ctx)
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

**[*operations.SendDeletePlainTextResponse](../../pkg/models/operations/senddeleteplaintextresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendDynamic

Send Dynamic

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
    res, err := s.BodyParams.SendDynamic(ctx, []byte("0x9cD9996CFe"))
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
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SendDynamicResponse](../../pkg/models/operations/senddynamicresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UpdateModel

Update Model

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.BodyParams.UpdateModel(ctx, shared.Employee{
        Address: "47839 Padberg Pass",
        Age: 993094,
        Birthday: types.MustDateFromString("2022-03-30"),
        Birthtime: types.MustTimeFromString("2022-07-06T16:02:40.468Z"),
        Department: "<value>",
        Dependents: []shared.Person{
            shared.Person{
                Address: "3734 Brisa Pine",
                Age: 27925,
                Birthday: types.MustDateFromString("2024-02-25"),
                Birthtime: types.MustTimeFromString("2022-11-28T02:22:58.969Z"),
                Name: "<value>",
                PersonType: "<value>",
                UID: "<value>",
            },
        },
        HiredAt: "<value>",
        JoiningDay: shared.DaysMonday,
        Name: "<value>",
        PersonType: "<value>",
        Salary: 714606,
        UID: "<value>",
        WorkingDays: []shared.Days{
            shared.DaysWednesday,
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

| Parameter                                              | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `ctx`                                                  | [context.Context](https://pkg.go.dev/context#Context)  | :heavy_check_mark:                                     | The context to use for the request.                    |
| `request`                                              | [shared.Employee](../../pkg/models/shared/employee.md) | :heavy_check_mark:                                     | The request object to use for the request.             |


### Response

**[*operations.UpdateModelResponse](../../pkg/models/operations/updatemodelresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendDeleteBody

send Delete Body

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
    res, err := s.BodyParams.SendDeleteBody(ctx)
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

**[*operations.SendDeleteBodyResponse](../../pkg/models/operations/senddeletebodyresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UpdateString

update String

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
    res, err := s.BodyParams.UpdateString(ctx, "<value>")
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
| `request`                                             | [string](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdateStringResponse](../../pkg/models/operations/updatestringresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
