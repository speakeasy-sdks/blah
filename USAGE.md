<!-- Start SDK Example Usage [usage] -->
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
	res, err := s.Echo.Jsonecho(ctx, []byte("0x4CcAf4eebe"))
	if err != nil {
		log.Fatal(err)
	}
	if res.Res != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->