# Get400Response


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ContentType`                                                             | *string*                                                                  | :heavy_check_mark:                                                        | HTTP response content type for this operation                             |
| `Get400200TextPlainObject`                                                | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |
| `GlobalTestException`                                                     | [*shared.GlobalTestException](../../models/shared/globaltestexception.md) | :heavy_minus_sign:                                                        | 500 Global                                                                |
| `StatusCode`                                                              | *int*                                                                     | :heavy_check_mark:                                                        | HTTP response status code for this operation                              |
| `RawResponse`                                                             | [*http.Response](https://pkg.go.dev/net/http#Response)                    | :heavy_minus_sign:                                                        | Raw HTTP response; suitable for custom response parsing                   |