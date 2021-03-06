// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// VROPSadapterOperationUsingPOSTReader is a Reader for the VROPSadapterOperationUsingPOST structure.
type VROPSadapterOperationUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VROPSadapterOperationUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVROPSadapterOperationUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewVROPSadapterOperationUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVROPSadapterOperationUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVROPSadapterOperationUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVROPSadapterOperationUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVROPSadapterOperationUsingPOSTOK creates a VROPSadapterOperationUsingPOSTOK with default headers values
func NewVROPSadapterOperationUsingPOSTOK() *VROPSadapterOperationUsingPOSTOK {
	return &VROPSadapterOperationUsingPOSTOK{}
}

/*VROPSadapterOperationUsingPOSTOK handles this case with default header values.

OK
*/
type VROPSadapterOperationUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *VROPSadapterOperationUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/vrops/adapterOperation][%d] vROPSadapterOperationUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *VROPSadapterOperationUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *VROPSadapterOperationUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVROPSadapterOperationUsingPOSTCreated creates a VROPSadapterOperationUsingPOSTCreated with default headers values
func NewVROPSadapterOperationUsingPOSTCreated() *VROPSadapterOperationUsingPOSTCreated {
	return &VROPSadapterOperationUsingPOSTCreated{}
}

/*VROPSadapterOperationUsingPOSTCreated handles this case with default header values.

Created
*/
type VROPSadapterOperationUsingPOSTCreated struct {
}

func (o *VROPSadapterOperationUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/vrops/adapterOperation][%d] vROPSadapterOperationUsingPOSTCreated ", 201)
}

func (o *VROPSadapterOperationUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVROPSadapterOperationUsingPOSTUnauthorized creates a VROPSadapterOperationUsingPOSTUnauthorized with default headers values
func NewVROPSadapterOperationUsingPOSTUnauthorized() *VROPSadapterOperationUsingPOSTUnauthorized {
	return &VROPSadapterOperationUsingPOSTUnauthorized{}
}

/*VROPSadapterOperationUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type VROPSadapterOperationUsingPOSTUnauthorized struct {
}

func (o *VROPSadapterOperationUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/vrops/adapterOperation][%d] vROPSadapterOperationUsingPOSTUnauthorized ", 401)
}

func (o *VROPSadapterOperationUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVROPSadapterOperationUsingPOSTForbidden creates a VROPSadapterOperationUsingPOSTForbidden with default headers values
func NewVROPSadapterOperationUsingPOSTForbidden() *VROPSadapterOperationUsingPOSTForbidden {
	return &VROPSadapterOperationUsingPOSTForbidden{}
}

/*VROPSadapterOperationUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type VROPSadapterOperationUsingPOSTForbidden struct {
}

func (o *VROPSadapterOperationUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/vrops/adapterOperation][%d] vROPSadapterOperationUsingPOSTForbidden ", 403)
}

func (o *VROPSadapterOperationUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVROPSadapterOperationUsingPOSTNotFound creates a VROPSadapterOperationUsingPOSTNotFound with default headers values
func NewVROPSadapterOperationUsingPOSTNotFound() *VROPSadapterOperationUsingPOSTNotFound {
	return &VROPSadapterOperationUsingPOSTNotFound{}
}

/*VROPSadapterOperationUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type VROPSadapterOperationUsingPOSTNotFound struct {
}

func (o *VROPSadapterOperationUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/vrops/adapterOperation][%d] vROPSadapterOperationUsingPOSTNotFound ", 404)
}

func (o *VROPSadapterOperationUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
