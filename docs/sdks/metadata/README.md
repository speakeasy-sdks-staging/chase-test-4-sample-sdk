# Metadata
(*Metadata*)

## Overview

REST APIs for managing Version Metadata entities

### Available Operations

* [DeleteVersionMetadata](#deleteversionmetadata) - Delete metadata for a particular apiID and versionID.
* [GetVersionMetadata](#getversionmetadata) - Get all metadata for a particular apiID and versionID.
* [InsertVersionMetadata](#insertversionmetadata) - Insert metadata for a particular apiID and versionID.

## DeleteVersionMetadata

Delete metadata for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	chasetest4samplesdk "github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/operations"
)

func main() {
    s := chasetest4samplesdk.New(
        chasetest4samplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Metadata.DeleteVersionMetadata(ctx, operations.DeleteVersionMetadataRequest{
        APIID: "string",
        MetaKey: "string",
        MetaValue: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.DeleteVersionMetadataRequest](../../models/operations/deleteversionmetadatarequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.DeleteVersionMetadataResponse](../../models/operations/deleteversionmetadataresponse.md), error**


## GetVersionMetadata

Get all metadata for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	chasetest4samplesdk "github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/operations"
)

func main() {
    s := chasetest4samplesdk.New(
        chasetest4samplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Metadata.GetVersionMetadata(ctx, operations.GetVersionMetadataRequest{
        APIID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetVersionMetadataRequest](../../models/operations/getversionmetadatarequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetVersionMetadataResponse](../../models/operations/getversionmetadataresponse.md), error**


## InsertVersionMetadata

Insert metadata for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	chasetest4samplesdk "github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/operations"
)

func main() {
    s := chasetest4samplesdk.New(
        chasetest4samplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Metadata.InsertVersionMetadata(ctx, operations.InsertVersionMetadataRequest{
        VersionMetadataInput: shared.VersionMetadataInput{
            MetaKey: "string",
            MetaValue: "string",
        },
        APIID: "string",
        VersionID: "string",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.InsertVersionMetadataRequest](../../models/operations/insertversionmetadatarequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.InsertVersionMetadataResponse](../../models/operations/insertversionmetadataresponse.md), error**

