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

// UpdateEnvironmentUsingPUTReader is a Reader for the UpdateEnvironmentUsingPUT structure.
type UpdateEnvironmentUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEnvironmentUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEnvironmentUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateEnvironmentUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateEnvironmentUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEnvironmentUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEnvironmentUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEnvironmentUsingPUTOK creates a UpdateEnvironmentUsingPUTOK with default headers values
func NewUpdateEnvironmentUsingPUTOK() *UpdateEnvironmentUsingPUTOK {
	return &UpdateEnvironmentUsingPUTOK{}
}

/*UpdateEnvironmentUsingPUTOK handles this case with default header values.

OK
*/
type UpdateEnvironmentUsingPUTOK struct {
	Payload *models.EnvironmentDTO
}

func (o *UpdateEnvironmentUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/environments][%d] updateEnvironmentUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateEnvironmentUsingPUTOK) GetPayload() *models.EnvironmentDTO {
	return o.Payload
}

func (o *UpdateEnvironmentUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEnvironmentUsingPUTCreated creates a UpdateEnvironmentUsingPUTCreated with default headers values
func NewUpdateEnvironmentUsingPUTCreated() *UpdateEnvironmentUsingPUTCreated {
	return &UpdateEnvironmentUsingPUTCreated{}
}

/*UpdateEnvironmentUsingPUTCreated handles this case with default header values.

Created
*/
type UpdateEnvironmentUsingPUTCreated struct {
}

func (o *UpdateEnvironmentUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/environments][%d] updateEnvironmentUsingPUTCreated ", 201)
}

func (o *UpdateEnvironmentUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentUsingPUTUnauthorized creates a UpdateEnvironmentUsingPUTUnauthorized with default headers values
func NewUpdateEnvironmentUsingPUTUnauthorized() *UpdateEnvironmentUsingPUTUnauthorized {
	return &UpdateEnvironmentUsingPUTUnauthorized{}
}

/*UpdateEnvironmentUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateEnvironmentUsingPUTUnauthorized struct {
}

func (o *UpdateEnvironmentUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/environments][%d] updateEnvironmentUsingPUTUnauthorized ", 401)
}

func (o *UpdateEnvironmentUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentUsingPUTForbidden creates a UpdateEnvironmentUsingPUTForbidden with default headers values
func NewUpdateEnvironmentUsingPUTForbidden() *UpdateEnvironmentUsingPUTForbidden {
	return &UpdateEnvironmentUsingPUTForbidden{}
}

/*UpdateEnvironmentUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type UpdateEnvironmentUsingPUTForbidden struct {
}

func (o *UpdateEnvironmentUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/environments][%d] updateEnvironmentUsingPUTForbidden ", 403)
}

func (o *UpdateEnvironmentUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEnvironmentUsingPUTNotFound creates a UpdateEnvironmentUsingPUTNotFound with default headers values
func NewUpdateEnvironmentUsingPUTNotFound() *UpdateEnvironmentUsingPUTNotFound {
	return &UpdateEnvironmentUsingPUTNotFound{}
}

/*UpdateEnvironmentUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdateEnvironmentUsingPUTNotFound struct {
}

func (o *UpdateEnvironmentUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/environments][%d] updateEnvironmentUsingPUTNotFound ", 404)
}

func (o *UpdateEnvironmentUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
