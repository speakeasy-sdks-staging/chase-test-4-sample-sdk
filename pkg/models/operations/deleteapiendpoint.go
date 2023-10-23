// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"net/http"
)

type DeleteAPIEndpointRequest struct {
	// The ID of the ApiEndpoint to delete.
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	// The ID of the Api the ApiEndpoint belongs to.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api the ApiEndpoint belongs to.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *DeleteAPIEndpointRequest) GetAPIEndpointID() string {
	if o == nil {
		return ""
	}
	return o.APIEndpointID
}

func (o *DeleteAPIEndpointRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *DeleteAPIEndpointRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type DeleteAPIEndpointResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteAPIEndpointResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteAPIEndpointResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DeleteAPIEndpointResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteAPIEndpointResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}