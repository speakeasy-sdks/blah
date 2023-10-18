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
        File: operations.SendFileRequestBodyFile{
            Content: []byte("uAspk_G{\["),
            File: "Fadel",
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
        File: operations.SendMixedArrayRequestBodyFile{
            Content: []byte("N|}Ave]2L-"),
            File: "Electric",
        },
        Integers: []int{
            816429,
        },
        Models: []shared.Employee{
            shared.Employee{
                Address: "22003 Grimes Mills",
                Age: 951933,
                Birthday: types.MustDateFromString("2021-03-22"),
                Birthtime: types.MustTimeFromString("2021-10-17T14:28:49.414Z"),
                Boss: &shared.Person{
                    Address: "94257 Jacques Forks",
                    Age: 480419,
                    Birthday: types.MustDateFromString("2023-09-25"),
                    Birthtime: types.MustTimeFromString("2022-01-13T00:34:45.421Z"),
                    Name: "visualize",
                    PersonType: "Southeast",
                    UID: "card",
                },
                Department: "Investor",
                Dependents: []shared.Person{
                    shared.Person{
                        Address: "4365 Frami Landing",
                        Age: 988098,
                        Birthday: types.MustDateFromString("2023-07-29"),
                        Birthtime: types.MustTimeFromString("2022-01-20T22:03:14.338Z"),
                        Name: "calculate",
                        PersonType: "Kennedi",
                        UID: "Northeast",
                    },
                },
                HiredAt: "Bangladesh",
                JoiningDay: shared.DaysSaturday,
                Name: "voluptatibus",
                PersonType: "North",
                Salary: 484536,
                UID: "Missouri",
                WorkingDays: []shared.Days{
                    shared.DaysTuesday,
                },
            },
        },
        Strings: []string{
            "helpfully",
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
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.SendModel(ctx, &operations.SendModelRequestBody{
        Model: "Checking",
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
        Datetime: "panel",
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
                "Bedfordshire",
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
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.FormParams.UpdateModelwithForm(ctx, &operations.UpdateModelwithFormRequestBody{
        Model: "Towels",
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
        Value: "UDP",
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

