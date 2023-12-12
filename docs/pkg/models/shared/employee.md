# Employee


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Address`                                               | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Age`                                                   | *int64*                                                 | :heavy_check_mark:                                      | N/A                                                     |
| `Birthday`                                              | [types.Date](../../../types/date.md)                    | :heavy_check_mark:                                      | N/A                                                     |
| `Birthtime`                                             | [time.Time](https://pkg.go.dev/time#Time)               | :heavy_check_mark:                                      | N/A                                                     |
| `Boss`                                                  | [*shared.Person](../../../pkg/models/shared/person.md)  | :heavy_minus_sign:                                      | N/A                                                     |
| `Department`                                            | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Dependents`                                            | [][shared.Person](../../../pkg/models/shared/person.md) | :heavy_check_mark:                                      | N/A                                                     |
| `HiredAt`                                               | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `JoiningDay`                                            | [shared.Days](../../../pkg/models/shared/days.md)       | :heavy_check_mark:                                      | A string enum representing days of the week             |
| `Name`                                                  | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `PersonType`                                            | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Salary`                                                | *int*                                                   | :heavy_check_mark:                                      | N/A                                                     |
| `UID`                                                   | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `WorkingDays`                                           | [][shared.Days](../../../pkg/models/shared/days.md)     | :heavy_check_mark:                                      | N/A                                                     |