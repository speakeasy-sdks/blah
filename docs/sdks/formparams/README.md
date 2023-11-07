# FormParams
(*.FormParams*)

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
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.SendDate(ctx, &operations.SendDateRequestBody{
        Date: types.MustDateFromString("2021-10-04"),
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
| `request`                                                                        | [operations.SendDateRequestBody](../../models/operations/senddaterequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SendDateResponse](../../models/operations/senddateresponse.md), error**


## SendFile

Send File

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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SendFileRequestBody](../../models/operations/sendfilerequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SendFileResponse](../../models/operations/sendfileresponse.md), error**


## SendIntegerEnumArray

SendIntegerEnumArray

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
    res, err := s.FormParams.SendIntegerEnumArray(ctx, operations.SendIntegerEnumArrayRequest{
        RequestBody: &operations.SendIntegerEnumArrayRequestBody{
            Suites: []shared.SuiteCode{
                shared.SuiteCodeFour,
            },
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.SendIntegerEnumArrayRequest](../../models/operations/sendintegerenumarrayrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.SendIntegerEnumArrayResponse](../../models/operations/sendintegerenumarrayresponse.md), error**


## SendLong

Send Long

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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.SendLongRequestBody](../../models/operations/sendlongrequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.SendLongResponse](../../models/operations/sendlongresponse.md), error**


## SendMixedArray

Send a variety for form params. Returns file count and body params

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/types"
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
                Birthday: types.MustDateFromString("2022-06-13"),
                Birthtime: types.MustTimeFromString("2021-11-15T23:08:18.123Z"),
                Boss: &shared.Person{
                    Address: "7442 Tamara Grove",
                    Age: 318424,
                    Birthday: types.MustDateFromString("2023-02-04"),
                    Birthtime: types.MustTimeFromString("2021-05-07T13:59:39.796Z"),
                    Name: "string",
                    PersonType: "string",
                    UID: "string",
                },
                Department: "string",
                Dependents: []shared.Person{
                    shared.Person{
                        Address: "536 King Locks",
                        Age: 351258,
                        Birthday: types.MustDateFromString("2023-01-25"),
                        Birthtime: types.MustTimeFromString("2022-08-23T17:04:03.001Z"),
                        Name: "string",
                        PersonType: "string",
                        UID: "string",
                    },
                },
                HiredAt: "string",
                JoiningDay: shared.DaysThursday,
                Name: "string",
                PersonType: "string",
                Salary: 244149,
                UID: "string",
                WorkingDays: []shared.Days{
                    shared.DaysWednesday,
                },
            },
        },
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SendMixedArrayRequestBody](../../models/operations/sendmixedarrayrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.SendMixedArrayResponse](../../models/operations/sendmixedarrayresponse.md), error**


## SendModel

Send Model

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.SendModel(ctx, &operations.SendModelRequestBody{
        Model: shared.Employee{
            Address: "0371 Nienow Coves",
            Age: 574270,
            Birthday: types.MustDateFromString("2021-05-28"),
            Birthtime: types.MustTimeFromString("2022-10-14T06:30:12.294Z"),
            Boss: &shared.Person{
                Address: "8002 Pansy Dale",
                Age: 509061,
                Birthday: types.MustDateFromString("2021-10-26"),
                Birthtime: types.MustTimeFromString("2023-06-05T00:58:45.263Z"),
                Name: "string",
                PersonType: "string",
                UID: "string",
            },
            Department: "string",
            Dependents: []shared.Person{
                shared.Person{
                    Address: "8830 McDermott Forks",
                    Age: 465073,
                    Birthday: types.MustDateFromString("2021-02-06"),
                    Birthtime: types.MustTimeFromString("2023-04-01T23:29:40.608Z"),
                    Name: "string",
                    PersonType: "string",
                    UID: "string",
                },
            },
            HiredAt: "string",
            JoiningDay: shared.DaysMonday,
            Name: "string",
            PersonType: "string",
            Salary: 35111,
            UID: "string",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.SendModelRequestBody](../../models/operations/sendmodelrequestbody.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.SendModelResponse](../../models/operations/sendmodelresponse.md), error**


## SendRfc1123DateTime

Send Rfc1123 DateTime

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
    res, err := s.FormParams.SendRfc1123DateTime(ctx, &operations.SendRfc1123DateTimeRequestBody{
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SendRfc1123DateTimeRequestBody](../../models/operations/sendrfc1123datetimerequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.SendRfc1123DateTimeResponse](../../models/operations/sendrfc1123datetimeresponse.md), error**


## SendRfc3339DateTime

Send Rfc3339 DateTime

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
    res, err := s.FormParams.SendRfc3339DateTime(ctx, &operations.SendRfc3339DateTimeRequestBody{
        Datetime: types.MustTimeFromString("2021-01-24T16:40:49.158Z"),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SendRfc3339DateTimeRequestBody](../../models/operations/sendrfc3339datetimerequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.SendRfc3339DateTimeResponse](../../models/operations/sendrfc3339datetimeresponse.md), error**


## SendStringArray

Send String Array

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
    res, err := s.FormParams.SendStringArray(ctx, operations.SendStringArrayRequest{
        RequestBody: &operations.SendStringArrayRequestBody{
            Strings: []string{
                "string",
            },
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.SendStringArrayRequest](../../models/operations/sendstringarrayrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.SendStringArrayResponse](../../models/operations/sendstringarrayresponse.md), error**


## SendStringEnumArray

SendStringEnumArray

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
    res, err := s.FormParams.SendStringEnumArray(ctx, operations.SendStringEnumArrayRequest{
        RequestBody: &operations.SendStringEnumArrayRequestBody{
            Days: []shared.Days{
                shared.DaysMonday,
            },
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.SendStringEnumArrayRequest](../../models/operations/sendstringenumarrayrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.SendStringEnumArrayResponse](../../models/operations/sendstringenumarrayresponse.md), error**


## SendUnixDateTime

Send UnixDateTime

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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.SendUnixDateTimeRequestBody](../../models/operations/sendunixdatetimerequestbody.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.SendUnixDateTimeResponse](../../models/operations/sendunixdatetimeresponse.md), error**


## UpdateModelwithForm

Update Model with Form

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.UpdateModelwithForm(ctx, &operations.UpdateModelwithFormRequestBody{
        Model: shared.Employee{
            Address: "37579 Sporer Fields",
            Age: 899553,
            Birthday: types.MustDateFromString("2023-02-10"),
            Birthtime: types.MustTimeFromString("2021-06-17T01:16:01.265Z"),
            Boss: &shared.Person{
                Address: "1901 Carter Lodge",
                Age: 445689,
                Birthday: types.MustDateFromString("2022-02-09"),
                Birthtime: types.MustTimeFromString("2022-03-17T14:13:39.740Z"),
                Name: "string",
                PersonType: "string",
                UID: "string",
            },
            Department: "string",
            Dependents: []shared.Person{
                shared.Person{
                    Address: "11041 Jacobi Stream",
                    Age: 134413,
                    Birthday: types.MustDateFromString("2021-08-30"),
                    Birthtime: types.MustTimeFromString("2023-11-26T23:43:21.899Z"),
                    Name: "string",
                    PersonType: "string",
                    UID: "string",
                },
            },
            HiredAt: "string",
            JoiningDay: shared.DaysFriday,
            Name: "string",
            PersonType: "string",
            Salary: 697151,
            UID: "string",
            WorkingDays: []shared.Days{
                shared.DaysMonday,
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateModelwithFormRequestBody](../../models/operations/updatemodelwithformrequestbody.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.UpdateModelwithFormResponse](../../models/operations/updatemodelwithformresponse.md), error**


## SendDeleteMultipart

send Delete Multipart

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

**[*operations.SendDeleteMultipartResponse](../../models/operations/senddeletemultipartresponse.md), error**


## SenddeleteForm

send delete Form

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

**[*operations.SenddeleteFormResponse](../../models/operations/senddeleteformresponse.md), error**


## SenddeleteForm1

send delete Form1

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

**[*operations.SenddeleteForm1Response](../../models/operations/senddeleteform1response.md), error**


## UpdateStringwithForm

update String with Form

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
    res, err := s.FormParams.UpdateStringwithForm(ctx, &operations.UpdateStringwithFormRequestBody{
        Value: "string",
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateStringwithFormRequestBody](../../models/operations/updatestringwithformrequestbody.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.UpdateStringwithFormResponse](../../models/operations/updatestringwithformresponse.md), error**

