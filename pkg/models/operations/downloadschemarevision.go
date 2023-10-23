// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"io"
	"net/http"
)

type DownloadSchemaRevisionRequest struct {
	// The ID of the Api to retrieve schemas for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The revision ID of the schema to retrieve.
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *DownloadSchemaRevisionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *DownloadSchemaRevisionRequest) GetRevisionID() string {
	if o == nil {
		return ""
	}
	return o.RevisionID
}

func (o *DownloadSchemaRevisionRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type DownloadSchemaRevisionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Schema io.ReadCloser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DownloadSchemaRevisionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DownloadSchemaRevisionResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DownloadSchemaRevisionResponse) GetSchema() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *DownloadSchemaRevisionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DownloadSchemaRevisionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}