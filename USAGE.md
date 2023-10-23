<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	chasetest4samplesdk "github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := chasetest4samplesdk.New(
		chasetest4samplesdk.WithSecurity(""),
	)

	ctx := context.Background()
	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
		Metadata: map[string][]string{
			"key": []string{
				"string",
			},
		},
		Op: &operations.GetApisOp{
			And: false,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Apis != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->