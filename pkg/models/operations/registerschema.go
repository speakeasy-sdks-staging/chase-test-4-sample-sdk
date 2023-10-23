// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks-staging/chase-test-4-sample-sdk/pkg/models/shared"
	"net/http"
)

type RegisterSchemaRequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

func (o *RegisterSchemaRequestBodyFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *RegisterSchemaRequestBodyFile) GetFile() string {
	if o == nil {
		return ""
	}
	return o.File
}

// RegisterSchemaRequestBody - The schema file to upload provided as a multipart/form-data file segment.
type RegisterSchemaRequestBody struct {
	File RegisterSchemaRequestBodyFile `multipartForm:"file"`
}

func (o *RegisterSchemaRequestBody) GetFile() RegisterSchemaRequestBodyFile {
	if o == nil {
		return RegisterSchemaRequestBodyFile{}
	}
	return o.File
}

type RegisterSchemaRequest struct {
	// The schema file to upload provided as a multipart/form-data file segment.
	RequestBody RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
	// The ID of the Api to get the schema for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *RegisterSchemaRequest) GetRequestBody() RegisterSchemaRequestBody {
	if o == nil {
		return RegisterSchemaRequestBody{}
	}
	return o.RequestBody
}

func (o *RegisterSchemaRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *RegisterSchemaRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type RegisterSchemaResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RegisterSchemaResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RegisterSchemaResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *RegisterSchemaResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RegisterSchemaResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}