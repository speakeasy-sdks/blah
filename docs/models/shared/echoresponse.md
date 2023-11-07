# EchoResponse

Raw http Request info


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `Body`                                                                    | map[string][shared.Body](../../models/shared/body.md)                     | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Headers`                                                                 | map[string]*string*                                                       | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Method`                                                                  | **string*                                                                 | :heavy_minus_sign:                                                        | N/A                                                                       |
| `Path`                                                                    | **string*                                                                 | :heavy_minus_sign:                                                        | relativePath                                                              |
| `Query`                                                                   | map[string][shared.QueryParameter](../../models/shared/queryparameter.md) | :heavy_minus_sign:                                                        | N/A                                                                       |
| `UploadCount`                                                             | **int*                                                                    | :heavy_minus_sign:                                                        | N/A                                                                       |