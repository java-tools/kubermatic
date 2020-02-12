// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// RevokeClusterViewerTokenReader is a Reader for the RevokeClusterViewerToken structure.
type RevokeClusterViewerTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeClusterViewerTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeClusterViewerTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeClusterViewerTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeClusterViewerTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRevokeClusterViewerTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeClusterViewerTokenOK creates a RevokeClusterViewerTokenOK with default headers values
func NewRevokeClusterViewerTokenOK() *RevokeClusterViewerTokenOK {
	return &RevokeClusterViewerTokenOK{}
}

/*RevokeClusterViewerTokenOK handles this case with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterViewerTokenOK struct {
}

func (o *RevokeClusterViewerTokenOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenOK ", 200)
}

func (o *RevokeClusterViewerTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterViewerTokenUnauthorized creates a RevokeClusterViewerTokenUnauthorized with default headers values
func NewRevokeClusterViewerTokenUnauthorized() *RevokeClusterViewerTokenUnauthorized {
	return &RevokeClusterViewerTokenUnauthorized{}
}

/*RevokeClusterViewerTokenUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterViewerTokenUnauthorized struct {
}

func (o *RevokeClusterViewerTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenUnauthorized ", 401)
}

func (o *RevokeClusterViewerTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterViewerTokenForbidden creates a RevokeClusterViewerTokenForbidden with default headers values
func NewRevokeClusterViewerTokenForbidden() *RevokeClusterViewerTokenForbidden {
	return &RevokeClusterViewerTokenForbidden{}
}

/*RevokeClusterViewerTokenForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterViewerTokenForbidden struct {
}

func (o *RevokeClusterViewerTokenForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerTokenForbidden ", 403)
}

func (o *RevokeClusterViewerTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterViewerTokenDefault creates a RevokeClusterViewerTokenDefault with default headers values
func NewRevokeClusterViewerTokenDefault(code int) *RevokeClusterViewerTokenDefault {
	return &RevokeClusterViewerTokenDefault{
		_statusCode: code,
	}
}

/*RevokeClusterViewerTokenDefault handles this case with default header values.

errorResponse
*/
type RevokeClusterViewerTokenDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the revoke cluster viewer token default response
func (o *RevokeClusterViewerTokenDefault) Code() int {
	return o._statusCode
}

func (o *RevokeClusterViewerTokenDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/viewertoken][%d] revokeClusterViewerToken default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeClusterViewerTokenDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RevokeClusterViewerTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
