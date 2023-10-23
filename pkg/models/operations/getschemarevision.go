// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetSchemaRevisionRequest struct {
	// The ID of the Api to retrieve schemas for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The revision ID of the schema to retrieve.
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GetSchemaRevisionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetSchemaRevisionRequest) GetRevisionID() string {
	if o == nil {
		return ""
	}
	return o.RevisionID
}

func (o *GetSchemaRevisionRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GetSchemaRevisionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	Schema *shared.Schema
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSchemaRevisionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSchemaRevisionResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetSchemaRevisionResponse) GetSchema() *shared.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *GetSchemaRevisionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSchemaRevisionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
