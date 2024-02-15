# FormParams
(*FormParams*)

### Available Operations

* [SendDate](#senddate) - Send Date
* [SendFile](#sendfile) - Send File
* [SendIntegerEnumArray](#sendintegerenumarray) - SendIntegerEnumArray
* [SendLong](#sendlong) - Send Long
* [SendMixedArray](#sendmixedarray) - SendMixedArray
* [SendModel](#sendmodel) - Send Model
* [SendRfc1123DateTime](#sendrfc1123datetime) - Send Rfc1123 DateTime
* [SendRfc3339DateTime](#sendrfc3339datetime) - Send Rfc3339 DateTime
* [SendStringArray](#sendstringarray) - Send String Array
* [SendStringEnumArray](#sendstringenumarray) - SendStringEnumArray
* [SendUnixDateTime](#sendunixdatetime) - Send UnixDateTime
* [UpdateModelwithForm](#updatemodelwithform) - Update Model with Form
* [SendDeleteMultipart](#senddeletemultipart) - send Delete Multipart
* [SenddeleteForm](#senddeleteform) - send delete Form
* [SenddeleteForm1](#senddeleteform1) - send delete Form1
* [UpdateStringwithForm](#updatestringwithform) - update String with Form

## SendDate

Send Date

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
    res, err := s.FormParams.SendDate(ctx, &operations.SendDateRequestBody{
        Date: types.MustDateFromString("2022-10-05"),
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
| `request`                                                                            | [operations.SendDateRequestBody](../../pkg/models/operations/senddaterequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SendDateResponse](../../pkg/models/operations/senddateresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendFile

Send File

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
    res, err := s.FormParams.SendFile(ctx, &operations.SendFileRequestBody{
        File: operations.File{
            Content: []byte("0xD7DCBe9Fdd"),
            FileName: "crewmen.gif",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.SendFileRequestBody](../../pkg/models/operations/sendfilerequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SendFileResponse](../../pkg/models/operations/sendfileresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendIntegerEnumArray

SendIntegerEnumArray

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
    res, err := s.FormParams.SendIntegerEnumArray(ctx, operations.SendIntegerEnumArrayRequest{
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
| `request`                                                                                            | [operations.SendIntegerEnumArrayRequest](../../pkg/models/operations/sendintegerenumarrayrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.SendIntegerEnumArrayResponse](../../pkg/models/operations/sendintegerenumarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendLong

Send Long

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
    res, err := s.FormParams.SendLong(ctx, &operations.SendLongRequestBody{
        Value: 982003,
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
| `request`                                                                            | [operations.SendLongRequestBody](../../pkg/models/operations/sendlongrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.SendLongResponse](../../pkg/models/operations/sendlongresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendMixedArray

Send a variety for form params. Returns file count and body params

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.SendMixedArray(ctx, &operations.SendMixedArrayRequestBody{
        File: operations.SendMixedArrayFile{
            Content: []byte("0xaFF7EAe4a2"),
            FileName: "networks_east_robust.pdf",
        },
        Integers: []int{
            311835,
        },
        Models: []shared.Employee{
            shared.Employee{
                Address: "259 Diego Estates",
                Age: 994876,
                Birthday: types.MustDateFromString("2023-06-13"),
                Birthtime: types.MustTimeFromString("2022-11-16T06:07:45.702Z"),
                Department: "<value>",
                Dependents: []shared.Person{
                    shared.Person{
                        Address: "7442 Tamara Grove",
                        Age: 318424,
                        Birthday: types.MustDateFromString("2024-02-04"),
                        Birthtime: types.MustTimeFromString("2022-05-07T16:46:07.722Z"),
                        Name: "<value>",
                        PersonType: "<value>",
                        UID: "<value>",
                    },
                },
                HiredAt: "<value>",
                JoiningDay: shared.DaysFriday,
                Name: "<value>",
                PersonType: "<value>",
                Salary: 523994,
                UID: "<value>",
                WorkingDays: []shared.Days{
                    shared.DaysTuesday,
                },
            },
        },
        Strings: []string{
            "<value>",
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
| `request`                                                                                        | [operations.SendMixedArrayRequestBody](../../pkg/models/operations/sendmixedarrayrequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.SendMixedArrayResponse](../../pkg/models/operations/sendmixedarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendModel

Send Model

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.SendModel(ctx, &operations.SendModelRequestBody{
        Model: shared.Employee{
            Address: "0371 Nienow Coves",
            Age: 574270,
            Birthday: types.MustDateFromString("2022-05-28"),
            Birthtime: types.MustTimeFromString("2023-10-14T20:46:40.251Z"),
            Department: "<value>",
            Dependents: []shared.Person{
                shared.Person{
                    Address: "8002 Pansy Dale",
                    Age: 509061,
                    Birthday: types.MustDateFromString("2022-10-26"),
                    Birthtime: types.MustTimeFromString("2024-06-04T20:22:38.619Z"),
                    Name: "<value>",
                    PersonType: "<value>",
                    UID: "<value>",
                },
            },
            HiredAt: "<value>",
            JoiningDay: shared.DaysTuesday,
            Name: "<value>",
            PersonType: "<value>",
            Salary: 879548,
            UID: "<value>",
            WorkingDays: []shared.Days{
                shared.DaysSaturday,
            },
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
| `request`                                                                              | [operations.SendModelRequestBody](../../pkg/models/operations/sendmodelrequestbody.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.SendModelResponse](../../pkg/models/operations/sendmodelresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendRfc1123DateTime

Send Rfc1123 DateTime

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
    res, err := s.FormParams.SendRfc1123DateTime(ctx, &operations.SendRfc1123DateTimeRequestBody{
        Datetime: "<value>",
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
| `request`                                                                                                  | [operations.SendRfc1123DateTimeRequestBody](../../pkg/models/operations/sendrfc1123datetimerequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.SendRfc1123DateTimeResponse](../../pkg/models/operations/sendrfc1123datetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendRfc3339DateTime

Send Rfc3339 DateTime

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
    res, err := s.FormParams.SendRfc3339DateTime(ctx, &operations.SendRfc3339DateTimeRequestBody{
        Datetime: types.MustTimeFromString("2022-01-24T17:11:58.792Z"),
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
| `request`                                                                                                  | [operations.SendRfc3339DateTimeRequestBody](../../pkg/models/operations/sendrfc3339datetimerequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.SendRfc3339DateTimeResponse](../../pkg/models/operations/sendrfc3339datetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendStringArray

Send String Array

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
    res, err := s.FormParams.SendStringArray(ctx, operations.SendStringArrayRequest{
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SendStringArrayRequest](../../pkg/models/operations/sendstringarrayrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SendStringArrayResponse](../../pkg/models/operations/sendstringarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendStringEnumArray

SendStringEnumArray

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
    res, err := s.FormParams.SendStringEnumArray(ctx, operations.SendStringEnumArrayRequest{
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
| `request`                                                                                          | [operations.SendStringEnumArrayRequest](../../pkg/models/operations/sendstringenumarrayrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.SendStringEnumArrayResponse](../../pkg/models/operations/sendstringenumarrayresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendUnixDateTime

Send UnixDateTime

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
    res, err := s.FormParams.SendUnixDateTime(ctx, &operations.SendUnixDateTimeRequestBody{
        Datetime: 6863.02,
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
| `request`                                                                                            | [operations.SendUnixDateTimeRequestBody](../../pkg/models/operations/sendunixdatetimerequestbody.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.SendUnixDateTimeResponse](../../pkg/models/operations/sendunixdatetimeresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UpdateModelwithForm

Update Model with Form

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/blah"
	"context"
	"github.com/speakeasy-sdks/blah/pkg/types"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.UpdateModelwithForm(ctx, &operations.UpdateModelwithFormRequestBody{
        Model: shared.Employee{
            Address: "37579 Sporer Fields",
            Age: 899553,
            Birthday: types.MustDateFromString("2024-02-11"),
            Birthtime: types.MustTimeFromString("2022-06-17T04:55:42.416Z"),
            Department: "<value>",
            Dependents: []shared.Person{
                shared.Person{
                    Address: "1901 Carter Lodge",
                    Age: 445689,
                    Birthday: types.MustDateFromString("2023-02-10"),
                    Birthtime: types.MustTimeFromString("2023-03-17T23:53:04.324Z"),
                    Name: "<value>",
                    PersonType: "<value>",
                    UID: "<value>",
                },
            },
            HiredAt: "<value>",
            JoiningDay: shared.DaysMonday,
            Name: "<value>",
            PersonType: "<value>",
            Salary: 103269,
            UID: "<value>",
            WorkingDays: []shared.Days{
                shared.DaysSunday,
            },
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateModelwithFormRequestBody](../../pkg/models/operations/updatemodelwithformrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.UpdateModelwithFormResponse](../../pkg/models/operations/updatemodelwithformresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SendDeleteMultipart

send Delete Multipart

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
    res, err := s.FormParams.SendDeleteMultipart(ctx)
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

**[*operations.SendDeleteMultipartResponse](../../pkg/models/operations/senddeletemultipartresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SenddeleteForm

send delete Form

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
    res, err := s.FormParams.SenddeleteForm(ctx)
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

**[*operations.SenddeleteFormResponse](../../pkg/models/operations/senddeleteformresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## SenddeleteForm1

send delete Form1

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
    res, err := s.FormParams.SenddeleteForm1(ctx)
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

**[*operations.SenddeleteForm1Response](../../pkg/models/operations/senddeleteform1response.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |

## UpdateStringwithForm

update String with Form

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
    res, err := s.FormParams.UpdateStringwithForm(ctx, &operations.UpdateStringwithFormRequestBody{
        Value: "<value>",
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
| `request`                                                                                                    | [operations.UpdateStringwithFormRequestBody](../../pkg/models/operations/updatestringwithformrequestbody.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.UpdateStringwithFormResponse](../../pkg/models/operations/updatestringwithformresponse.md), error**
| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 4xx-5xx                        | */*                            |
