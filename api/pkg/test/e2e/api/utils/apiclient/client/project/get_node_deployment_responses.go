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

// GetNodeDeploymentReader is a Reader for the GetNodeDeployment structure.
type GetNodeDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNodeDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNodeDeploymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeDeploymentOK creates a GetNodeDeploymentOK with default headers values
func NewGetNodeDeploymentOK() *GetNodeDeploymentOK {
	return &GetNodeDeploymentOK{}
}

/*GetNodeDeploymentOK handles this case with default header values.

NodeDeployment
*/
type GetNodeDeploymentOK struct {
	Payload *models.NodeDeployment
}

func (o *GetNodeDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetNodeDeploymentOK) GetPayload() *models.NodeDeployment {
	return o.Payload
}

func (o *GetNodeDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeployment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDeploymentUnauthorized creates a GetNodeDeploymentUnauthorized with default headers values
func NewGetNodeDeploymentUnauthorized() *GetNodeDeploymentUnauthorized {
	return &GetNodeDeploymentUnauthorized{}
}

/*GetNodeDeploymentUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetNodeDeploymentUnauthorized struct {
}

func (o *GetNodeDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentUnauthorized ", 401)
}

func (o *GetNodeDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeDeploymentForbidden creates a GetNodeDeploymentForbidden with default headers values
func NewGetNodeDeploymentForbidden() *GetNodeDeploymentForbidden {
	return &GetNodeDeploymentForbidden{}
}

/*GetNodeDeploymentForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetNodeDeploymentForbidden struct {
}

func (o *GetNodeDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeploymentForbidden ", 403)
}

func (o *GetNodeDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeDeploymentDefault creates a GetNodeDeploymentDefault with default headers values
func NewGetNodeDeploymentDefault(code int) *GetNodeDeploymentDefault {
	return &GetNodeDeploymentDefault{
		_statusCode: code,
	}
}

/*GetNodeDeploymentDefault handles this case with default header values.

errorResponse
*/
type GetNodeDeploymentDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get node deployment default response
func (o *GetNodeDeploymentDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeDeploymentDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}][%d] getNodeDeployment default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDeploymentDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetNodeDeploymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
