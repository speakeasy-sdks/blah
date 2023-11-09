# github.com/speakeasy-sdks/blah

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/blah.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/blah/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/blah
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
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
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Echo](docs/sdks/echo/README.md)

* [Jsonecho](docs/sdks/echo/README.md#jsonecho) - Json echo
* [QueryEcho](docs/sdks/echo/README.md#queryecho) - QueryEcho

### [BodyParams](docs/sdks/bodyparams/README.md)

* [PostSendDate](docs/sdks/bodyparams/README.md#postsenddate) - Send Date
* [PostSendIntegerArray](docs/sdks/bodyparams/README.md#postsendintegerarray) - Send Integer Array
* [PostSendIntegerEnumArray](docs/sdks/bodyparams/README.md#postsendintegerenumarray) - SendIntegerEnumArray
* [PostSendModel](docs/sdks/bodyparams/README.md#postsendmodel) - Send Model
* [PostSendRfc1123DateTime](docs/sdks/bodyparams/README.md#postsendrfc1123datetime) - Send Rfc1123 DateTime
* [PostSendRfc3339DateTime](docs/sdks/bodyparams/README.md#postsendrfc3339datetime) - Send Rfc3339 DateTime
* [PostSendStringArray](docs/sdks/bodyparams/README.md#postsendstringarray) - Send String Array
* [PostSendStringEnumArray](docs/sdks/bodyparams/README.md#postsendstringenumarray) - SendStringEnumArray
* [PostSendUnixDateTime](docs/sdks/bodyparams/README.md#postsendunixdatetime) - Send UnixDateTime
* [SendDeleteBodywithModel](docs/sdks/bodyparams/README.md#senddeletebodywithmodel) - Send Delete Body with Model
* [SendDeletePlainText](docs/sdks/bodyparams/README.md#senddeleteplaintext) - Send Delete PlainText
* [SendDynamic](docs/sdks/bodyparams/README.md#senddynamic) - Send Dynamic
* [UpdateModel](docs/sdks/bodyparams/README.md#updatemodel) - Update Model
* [SendDeleteBody](docs/sdks/bodyparams/README.md#senddeletebody) - send Delete Body
* [UpdateString](docs/sdks/bodyparams/README.md#updatestring) - update String

### [ErrorCodes](docs/sdks/errorcodes/README.md)

* [Get400](docs/sdks/errorcodes/README.md#get400) - Get400
* [Get401](docs/sdks/errorcodes/README.md#get401) - Get401
* [Get500](docs/sdks/errorcodes/README.md#get500) - Get500
* [Get501](docs/sdks/errorcodes/README.md#get501) - Get501
* [Catch412globalerror](docs/sdks/errorcodes/README.md#catch412globalerror) - catch 412 global error

### [FormParams](docs/sdks/formparams/README.md)

* [SendDate](docs/sdks/formparams/README.md#senddate) - Send Date
* [SendFile](docs/sdks/formparams/README.md#sendfile) - Send File
* [SendIntegerEnumArray](docs/sdks/formparams/README.md#sendintegerenumarray) - SendIntegerEnumArray
* [SendLong](docs/sdks/formparams/README.md#sendlong) - Send Long
* [SendMixedArray](docs/sdks/formparams/README.md#sendmixedarray) - SendMixedArray
* [SendModel](docs/sdks/formparams/README.md#sendmodel) - Send Model
* [SendRfc1123DateTime](docs/sdks/formparams/README.md#sendrfc1123datetime) - Send Rfc1123 DateTime
* [SendRfc3339DateTime](docs/sdks/formparams/README.md#sendrfc3339datetime) - Send Rfc3339 DateTime
* [SendStringArray](docs/sdks/formparams/README.md#sendstringarray) - Send String Array
* [SendStringEnumArray](docs/sdks/formparams/README.md#sendstringenumarray) - SendStringEnumArray
* [SendUnixDateTime](docs/sdks/formparams/README.md#sendunixdatetime) - Send UnixDateTime
* [UpdateModelwithForm](docs/sdks/formparams/README.md#updatemodelwithform) - Update Model with Form
* [SendDeleteMultipart](docs/sdks/formparams/README.md#senddeletemultipart) - send Delete Multipart
* [SenddeleteForm](docs/sdks/formparams/README.md#senddeleteform) - send delete Form
* [SenddeleteForm1](docs/sdks/formparams/README.md#senddeleteform1) - send delete Form1
* [UpdateStringwithForm](docs/sdks/formparams/README.md#updatestringwithform) - update String with Form

### [Header](docs/sdks/header/README.md)

* [SendHeaders](docs/sdks/header/README.md#sendheaders) - Send Headers

### [QueryParam](docs/sdks/queryparam/README.md)

* [Date](docs/sdks/queryparam/README.md#date) - Date
* [DateArray](docs/sdks/queryparam/README.md#datearray) - Date Array
* [IntegerEnumArray](docs/sdks/queryparam/README.md#integerenumarray) - Integer Enum Array
* [MultipleParams](docs/sdks/queryparam/README.md#multipleparams) - MultipleParams
* [NoParams](docs/sdks/queryparam/README.md#noparams) - NoParams
* [NumberArray](docs/sdks/queryparam/README.md#numberarray) - Number Array
* [OptionalDynamicQueryParam](docs/sdks/queryparam/README.md#optionaldynamicqueryparam) - Optional Dynamic Query Param
* [Rfc1123DateTime](docs/sdks/queryparam/README.md#rfc1123datetime) - Rfc1123 DateTime
* [Rfc1123DateTimeArray](docs/sdks/queryparam/README.md#rfc1123datetimearray) - Rfc1123 DateTime Array
* [Rfc3339DateTime](docs/sdks/queryparam/README.md#rfc3339datetime) - Rfc3339 DateTime
* [Rfc3339DateTimeArray](docs/sdks/queryparam/README.md#rfc3339datetimearray) - Rfc3339 DateTime Array
* [SimpleQuery](docs/sdks/queryparam/README.md#simplequery) - SimpleQuery
* [StringArray](docs/sdks/queryparam/README.md#stringarray) - String Array
* [StringEnumArray](docs/sdks/queryparam/README.md#stringenumarray) - String Enum Array
* [StringParam](docs/sdks/queryparam/README.md#stringparam) - StringParam
* [UnixDateTime](docs/sdks/queryparam/README.md#unixdatetime) - Unix DateTime
* [UnixDateTimeArray](docs/sdks/queryparam/README.md#unixdatetimearray) - Unix DateTime Array
* [URLParam](docs/sdks/queryparam/README.md#urlparam) - UrlParam

### [ResponseTypes](docs/sdks/responsetypes/README.md)

* [Get1123DateTime](docs/sdks/responsetypes/README.md#get1123datetime) - Get 1123DateTime
* [Get3339Datetime](docs/sdks/responsetypes/README.md#get3339datetime) - Get 3339Datetime
* [GetBinary](docs/sdks/responsetypes/README.md#getbinary) - Get Binary
* [GetBoolean](docs/sdks/responsetypes/README.md#getboolean) - Get Boolean
* [GetDateArray](docs/sdks/responsetypes/README.md#getdatearray) - Get Date Array
* [GetDynamic](docs/sdks/responsetypes/README.md#getdynamic) - Get Dynamic
* [GetHeaders](docs/sdks/responsetypes/README.md#getheaders) - Get Headers
* [GetInteger](docs/sdks/responsetypes/README.md#getinteger) - Get Integer
* [GetLong](docs/sdks/responsetypes/README.md#getlong) - Get Long
* [GetModelArray](docs/sdks/responsetypes/README.md#getmodelarray) - Get Model Array
* [GetPrecision](docs/sdks/responsetypes/README.md#getprecision) - Get Precision
* [GetStringEnumArray](docs/sdks/responsetypes/README.md#getstringenumarray) - Get String Enum Array
* [GetUnixDateTime](docs/sdks/responsetypes/README.md#getunixdatetime) - Get UnixDateTime
* [Getcontenttypeheaders](docs/sdks/responsetypes/README.md#getcontenttypeheaders) - get content type headers
* [Returnresponsewithenums](docs/sdks/responsetypes/README.md#returnresponsewithenums) - return response with enums
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:


<!-- End Pagination -->

<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                   | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| sdkerrors.NestedModelException | 412                            | application/json               |
| sdkerrors.GlobalTestException  | 500                            | application/json               |
| sdkerrors.SDKError             | 400-600                        | */*                            |


## Example

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
	s := blah.New()

	ctx := context.Background()
	res, err := s.Echo.Jsonecho(ctx, []byte("0x4CcAf4eebe"))
	if err != nil {

		var e *sdkerrors.NestedModelException
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.GlobalTestException
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `http://localhost:3000` | None |

For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
	s := blah.New(
		blah.WithServerIndex(0),
	)

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


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/blah"
	"github.com/speakeasy-sdks/blah/pkg/models/operations"
	"log"
)

func main() {
	s := blah.New(
		blah.WithServerURL("http://localhost:3000"),
	)

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

## Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:

```go
package main

import (
	"context"
	"github.com/speakeasy-sdks/blah"
	"log"
)

func main() {
	s := blah.New()

	ctx := context.Background()
	res, err := s.ResponseTypes.GetLong(ctx, operations.WithServerURL("http://localhost:3000"))
	if err != nil {
		log.Fatal(err)
	}

	if res.Res != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
