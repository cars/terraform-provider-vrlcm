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

// EnvironmentTemplateUsingPOSTReader is a Reader for the EnvironmentTemplateUsingPOST structure.
type EnvironmentTemplateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnvironmentTemplateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnvironmentTemplateUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewEnvironmentTemplateUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEnvironmentTemplateUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnvironmentTemplateUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEnvironmentTemplateUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewEnvironmentTemplateUsingPOSTOK creates a EnvironmentTemplateUsingPOSTOK with default headers values
func NewEnvironmentTemplateUsingPOSTOK() *EnvironmentTemplateUsingPOSTOK {
	return &EnvironmentTemplateUsingPOSTOK{}
}

/*EnvironmentTemplateUsingPOSTOK handles this case with default header values.

OK
*/
type EnvironmentTemplateUsingPOSTOK struct {
	Payload *models.EnvironmentRequestDTO
}

func (o *EnvironmentTemplateUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/import/template/{type}][%d] environmentTemplateUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *EnvironmentTemplateUsingPOSTOK) GetPayload() *models.EnvironmentRequestDTO {
	return o.Payload
}

func (o *EnvironmentTemplateUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentRequestDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnvironmentTemplateUsingPOSTCreated creates a EnvironmentTemplateUsingPOSTCreated with default headers values
func NewEnvironmentTemplateUsingPOSTCreated() *EnvironmentTemplateUsingPOSTCreated {
	return &EnvironmentTemplateUsingPOSTCreated{}
}

/*EnvironmentTemplateUsingPOSTCreated handles this case with default header values.

Created
*/
type EnvironmentTemplateUsingPOSTCreated struct {
}

func (o *EnvironmentTemplateUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/import/template/{type}][%d] environmentTemplateUsingPOSTCreated ", 201)
}

func (o *EnvironmentTemplateUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnvironmentTemplateUsingPOSTUnauthorized creates a EnvironmentTemplateUsingPOSTUnauthorized with default headers values
func NewEnvironmentTemplateUsingPOSTUnauthorized() *EnvironmentTemplateUsingPOSTUnauthorized {
	return &EnvironmentTemplateUsingPOSTUnauthorized{}
}

/*EnvironmentTemplateUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type EnvironmentTemplateUsingPOSTUnauthorized struct {
}

func (o *EnvironmentTemplateUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/import/template/{type}][%d] environmentTemplateUsingPOSTUnauthorized ", 401)
}

func (o *EnvironmentTemplateUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnvironmentTemplateUsingPOSTForbidden creates a EnvironmentTemplateUsingPOSTForbidden with default headers values
func NewEnvironmentTemplateUsingPOSTForbidden() *EnvironmentTemplateUsingPOSTForbidden {
	return &EnvironmentTemplateUsingPOSTForbidden{}
}

/*EnvironmentTemplateUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type EnvironmentTemplateUsingPOSTForbidden struct {
}

func (o *EnvironmentTemplateUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/import/template/{type}][%d] environmentTemplateUsingPOSTForbidden ", 403)
}

func (o *EnvironmentTemplateUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnvironmentTemplateUsingPOSTNotFound creates a EnvironmentTemplateUsingPOSTNotFound with default headers values
func NewEnvironmentTemplateUsingPOSTNotFound() *EnvironmentTemplateUsingPOSTNotFound {
	return &EnvironmentTemplateUsingPOSTNotFound{}
}

/*EnvironmentTemplateUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type EnvironmentTemplateUsingPOSTNotFound struct {
}

func (o *EnvironmentTemplateUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/import/template/{type}][%d] environmentTemplateUsingPOSTNotFound ", 404)
}

func (o *EnvironmentTemplateUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
