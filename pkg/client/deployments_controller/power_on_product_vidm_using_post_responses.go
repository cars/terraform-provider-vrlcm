// Code generated by go-swagger; DO NOT EDIT.

package deployments_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// PowerOnProductVidmUsingPOSTReader is a Reader for the PowerOnProductVidmUsingPOST structure.
type PowerOnProductVidmUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PowerOnProductVidmUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPowerOnProductVidmUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPowerOnProductVidmUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPowerOnProductVidmUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPowerOnProductVidmUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPowerOnProductVidmUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPowerOnProductVidmUsingPOSTOK creates a PowerOnProductVidmUsingPOSTOK with default headers values
func NewPowerOnProductVidmUsingPOSTOK() *PowerOnProductVidmUsingPOSTOK {
	return &PowerOnProductVidmUsingPOSTOK{}
}

/*PowerOnProductVidmUsingPOSTOK handles this case with default header values.

OK
*/
type PowerOnProductVidmUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *PowerOnProductVidmUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/vidm/powerOn][%d] powerOnProductVidmUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *PowerOnProductVidmUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *PowerOnProductVidmUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPowerOnProductVidmUsingPOSTCreated creates a PowerOnProductVidmUsingPOSTCreated with default headers values
func NewPowerOnProductVidmUsingPOSTCreated() *PowerOnProductVidmUsingPOSTCreated {
	return &PowerOnProductVidmUsingPOSTCreated{}
}

/*PowerOnProductVidmUsingPOSTCreated handles this case with default header values.

Created
*/
type PowerOnProductVidmUsingPOSTCreated struct {
}

func (o *PowerOnProductVidmUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/vidm/powerOn][%d] powerOnProductVidmUsingPOSTCreated ", 201)
}

func (o *PowerOnProductVidmUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPowerOnProductVidmUsingPOSTUnauthorized creates a PowerOnProductVidmUsingPOSTUnauthorized with default headers values
func NewPowerOnProductVidmUsingPOSTUnauthorized() *PowerOnProductVidmUsingPOSTUnauthorized {
	return &PowerOnProductVidmUsingPOSTUnauthorized{}
}

/*PowerOnProductVidmUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type PowerOnProductVidmUsingPOSTUnauthorized struct {
}

func (o *PowerOnProductVidmUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/vidm/powerOn][%d] powerOnProductVidmUsingPOSTUnauthorized ", 401)
}

func (o *PowerOnProductVidmUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPowerOnProductVidmUsingPOSTForbidden creates a PowerOnProductVidmUsingPOSTForbidden with default headers values
func NewPowerOnProductVidmUsingPOSTForbidden() *PowerOnProductVidmUsingPOSTForbidden {
	return &PowerOnProductVidmUsingPOSTForbidden{}
}

/*PowerOnProductVidmUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type PowerOnProductVidmUsingPOSTForbidden struct {
}

func (o *PowerOnProductVidmUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/vidm/powerOn][%d] powerOnProductVidmUsingPOSTForbidden ", 403)
}

func (o *PowerOnProductVidmUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPowerOnProductVidmUsingPOSTNotFound creates a PowerOnProductVidmUsingPOSTNotFound with default headers values
func NewPowerOnProductVidmUsingPOSTNotFound() *PowerOnProductVidmUsingPOSTNotFound {
	return &PowerOnProductVidmUsingPOSTNotFound{}
}

/*PowerOnProductVidmUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type PowerOnProductVidmUsingPOSTNotFound struct {
}

func (o *PowerOnProductVidmUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/vidm/powerOn][%d] powerOnProductVidmUsingPOSTNotFound ", 404)
}

func (o *PowerOnProductVidmUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}