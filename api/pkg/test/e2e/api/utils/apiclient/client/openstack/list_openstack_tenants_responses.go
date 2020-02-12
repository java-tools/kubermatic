// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListOpenstackTenantsReader is a Reader for the ListOpenstackTenants structure.
type ListOpenstackTenantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackTenantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListOpenstackTenantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListOpenstackTenantsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackTenantsOK creates a ListOpenstackTenantsOK with default headers values
func NewListOpenstackTenantsOK() *ListOpenstackTenantsOK {
	return &ListOpenstackTenantsOK{}
}

/*ListOpenstackTenantsOK handles this case with default header values.

OpenstackTenant
*/
type ListOpenstackTenantsOK struct {
	Payload []*models.OpenstackTenant
}

func (o *ListOpenstackTenantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/tenants][%d] listOpenstackTenantsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackTenantsOK) GetPayload() []*models.OpenstackTenant {
	return o.Payload
}

func (o *ListOpenstackTenantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackTenantsDefault creates a ListOpenstackTenantsDefault with default headers values
func NewListOpenstackTenantsDefault(code int) *ListOpenstackTenantsDefault {
	return &ListOpenstackTenantsDefault{
		_statusCode: code,
	}
}

/*ListOpenstackTenantsDefault handles this case with default header values.

errorResponse
*/
type ListOpenstackTenantsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list openstack tenants default response
func (o *ListOpenstackTenantsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackTenantsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/tenants][%d] listOpenstackTenants default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackTenantsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListOpenstackTenantsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
