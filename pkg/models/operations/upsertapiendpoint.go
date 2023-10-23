// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"net/http"
)

type UpsertAPIEndpointRequest struct {
	// A JSON representation of the ApiEndpoint to upsert.
	APIEndpointInput shared.APIEndpointInput `request:"mediaType=application/json"`
	// The ID of the ApiEndpoint to upsert.
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	// The ID of the Api the ApiEndpoint belongs to.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api the ApiEndpoint belongs to.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *UpsertAPIEndpointRequest) GetAPIEndpointInput() shared.APIEndpointInput {
	if o == nil {
		return shared.APIEndpointInput{}
	}
	return o.APIEndpointInput
}

func (o *UpsertAPIEndpointRequest) GetAPIEndpointID() string {
	if o == nil {
		return ""
	}
	return o.APIEndpointID
}

func (o *UpsertAPIEndpointRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *UpsertAPIEndpointRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type UpsertAPIEndpointResponse struct {
	// OK
	APIEndpoint *shared.APIEndpoint
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpsertAPIEndpointResponse) GetAPIEndpoint() *shared.APIEndpoint {
	if o == nil {
		return nil
	}
	return o.APIEndpoint
}

func (o *UpsertAPIEndpointResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertAPIEndpointResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *UpsertAPIEndpointResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertAPIEndpointResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
