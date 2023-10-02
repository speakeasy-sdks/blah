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
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.BodyParams.PostSendDate(ctx, []byte("Hy}S3c@&JC"))
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
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendDateResponse](../../models/operations/postsenddateresponse.md), error**


## PostSendIntegerArray

Send Integer Array

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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.PostSendIntegerArrayRequest](../../models/operations/postsendintegerarrayrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.PostSendIntegerArrayResponse](../../models/operations/postsendintegerarrayresponse.md), error**


## PostSendIntegerEnumArray

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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PostSendIntegerEnumArrayRequest](../../models/operations/postsendintegerenumarrayrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PostSendIntegerEnumArrayResponse](../../models/operations/postsendintegerenumarrayresponse.md), error**


## PostSendModel

Send Model

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.BodyParams.PostSendModel(ctx, shared.Employee{
        Address: "4584 Bailey Ways",
        Age: 85493,
        Birthday: types.MustDateFromString("2021-07-04"),
        Birthtime: types.MustTimeFromString("2021-02-27T01:41:35.355Z"),
        Boss: &shared.Person{
            Address: "881 Ryann Ways",
            Age: 46325,
            Birthday: types.MustDateFromString("2022-11-16"),
            Birthtime: types.MustTimeFromString("2023-02-10T04:24:43.071Z"),
            Name: "structure if",
            PersonType: "Rap",
            UID: "Borders blockchains Plastic",
        },
        Department: "Rustic Buckinghamshire",
        Dependents: []shared.Person{
            shared.Person{
                Address: "252 Hirthe Points",
                Age: 103701,
                Birthday: types.MustDateFromString("2023-06-03"),
                Birthtime: types.MustTimeFromString("2022-12-12T11:03:45.416Z"),
                Name: "arid Northwest",
                PersonType: "Bedfordshire",
                UID: "Venezuelan Minnesota",
            },
        },
        HiredAt: "backing",
        JoiningDay: shared.DaysMonday,
        Name: "Northwest",
        PersonType: "Angus",
        Salary: 677435,
        UID: "siemens",
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Employee](../../models/shared/employee.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendModelResponse](../../models/operations/postsendmodelresponse.md), error**


## PostSendRfc1123DateTime

Send Rfc1123 DateTime

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
    res, err := s.BodyParams.PostSendRfc1123DateTime(ctx, "Analyst Elegant")
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
| `request`                                             | [string](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendRfc1123DateTimeResponse](../../models/operations/postsendrfc1123datetimeresponse.md), error**


## PostSendRfc3339DateTime

Send Rfc3339 DateTime

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
    res, err := s.BodyParams.PostSendRfc3339DateTime(ctx, []byte("c62M[.\"a^h"))
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
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendRfc3339DateTimeResponse](../../models/operations/postsendrfc3339datetimeresponse.md), error**


## PostSendStringArray

sends a string body param

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
    res, err := s.BodyParams.PostSendStringArray(ctx, operations.PostSendStringArrayRequest{
        RequestBody: []string{
            "1080p",
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
| `request`                                                                                      | [operations.PostSendStringArrayRequest](../../models/operations/postsendstringarrayrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PostSendStringArrayResponse](../../models/operations/postsendstringarrayresponse.md), error**


## PostSendStringEnumArray

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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.PostSendStringEnumArrayRequest](../../models/operations/postsendstringenumarrayrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.PostSendStringEnumArrayResponse](../../models/operations/postsendstringenumarrayresponse.md), error**


## PostSendUnixDateTime

Send UnixDateTime

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
    res, err := s.BodyParams.PostSendUnixDateTime(ctx, []byte("/G=C%Fo@*s"))
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
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostSendUnixDateTimeResponse](../../models/operations/postsendunixdatetimeresponse.md), error**


## SendDeleteBodywithModel

Send Delete Body with Model

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

**[*operations.SendDeleteBodywithModelResponse](../../models/operations/senddeletebodywithmodelresponse.md), error**


## SendDeletePlainText

Send Delete PlainText

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

**[*operations.SendDeletePlainTextResponse](../../models/operations/senddeleteplaintextresponse.md), error**


## SendDynamic

Send Dynamic

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
    res, err := s.BodyParams.SendDynamic(ctx, []byte("HTrHHG<o{`"))
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
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SendDynamicResponse](../../models/operations/senddynamicresponse.md), error**


## UpdateModel

Update Model

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/shared"
	"github.com/speakeasy-sdks/blah/pkg/types"
)

func main() {
    s := blah.New()

    ctx := context.Background()
    res, err := s.BodyParams.UpdateModel(ctx, shared.Employee{
        Address: "47839 Padberg Pass",
        Age: 993094,
        Birthday: types.MustDateFromString("2021-03-30"),
        Birthtime: types.MustTimeFromString("2021-07-06T11:57:24.993Z"),
        Boss: &shared.Person{
            Address: "3734 Brisa Pine",
            Age: 27925,
            Birthday: types.MustDateFromString("2023-02-24"),
            Birthtime: types.MustTimeFromString("2021-11-27T19:07:57.711Z"),
            Name: "male",
            PersonType: "Gasoline Rhodium mint",
            UID: "Madagascar",
        },
        Department: "stimulating alarm",
        Dependents: []shared.Person{
            shared.Person{
                Address: "86897 Hadley River",
                Age: 383456,
                Birthday: types.MustDateFromString("2023-01-04"),
                Birthtime: types.MustTimeFromString("2022-10-08T22:52:02.234Z"),
                Name: "green",
                PersonType: "Designer Hyundai bluetooth",
                UID: "portal Specialist Tools",
            },
        },
        HiredAt: "Northeast henry",
        JoiningDay: shared.DaysFriday,
        Name: "Massachusetts generate Integration",
        PersonType: "volt",
        Salary: 930709,
        UID: "integrated",
        WorkingDays: []shared.Days{
            shared.DaysSunday,
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

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.Employee](../../models/shared/employee.md)    | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdateModelResponse](../../models/operations/updatemodelresponse.md), error**


## SendDeleteBody

send Delete Body

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

**[*operations.SendDeleteBodyResponse](../../models/operations/senddeletebodyresponse.md), error**


## UpdateString

update String

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
    res, err := s.BodyParams.UpdateString(ctx, "Helena")
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
| `request`                                             | [string](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.UpdateStringResponse](../../models/operations/updatestringresponse.md), error**

