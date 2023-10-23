// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"net/http"
)

type DeleteVersionMetadataRequest struct {
	// The ID of the Api to delete metadata for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The key of the metadata to delete.
	MetaKey string `pathParam:"style=simple,explode=false,name=metaKey"`
	// The value of the metadata to delete.
	MetaValue string `pathParam:"style=simple,explode=false,name=metaValue"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *DeleteVersionMetadataRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *DeleteVersionMetadataRequest) GetMetaKey() string {
	if o == nil {
		return ""
	}
	return o.MetaKey
}

func (o *DeleteVersionMetadataRequest) GetMetaValue() string {
	if o == nil {
		return ""
	}
	return o.MetaValue
}

func (o *DeleteVersionMetadataRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type DeleteVersionMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteVersionMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteVersionMetadataResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DeleteVersionMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteVersionMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}