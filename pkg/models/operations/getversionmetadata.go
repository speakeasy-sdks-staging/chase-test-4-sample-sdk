// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetVersionMetadataRequest struct {
	// The ID of the Api to retrieve metadata for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to retrieve metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GetVersionMetadataRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetVersionMetadataRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GetVersionMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	VersionMetadata []shared.VersionMetadata
}

func (o *GetVersionMetadataResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetVersionMetadataResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetVersionMetadataResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetVersionMetadataResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetVersionMetadataResponse) GetVersionMetadata() []shared.VersionMetadata {
	if o == nil {
		return nil
	}
	return o.VersionMetadata
}