// Code generated by go-swagger; DO NOT EDIT.

package configuration_parameter_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateConfigUsingPUTReader is a Reader for the UpdateConfigUsingPUT structure.
type UpdateConfigUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateConfigUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateConfigUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateConfigUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateConfigUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateConfigUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateConfigUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateConfigUsingPUTOK creates a UpdateConfigUsingPUTOK with default headers values
func NewUpdateConfigUsingPUTOK() *UpdateConfigUsingPUTOK {
	return &UpdateConfigUsingPUTOK{}
}

/*UpdateConfigUsingPUTOK handles this case with default header values.

OK
*/
type UpdateConfigUsingPUTOK struct {
	Payload interface{}
}

func (o *UpdateConfigUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/configparams/{name}][%d] updateConfigUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateConfigUsingPUTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateConfigUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateConfigUsingPUTCreated creates a UpdateConfigUsingPUTCreated with default headers values
func NewUpdateConfigUsingPUTCreated() *UpdateConfigUsingPUTCreated {
	return &UpdateConfigUsingPUTCreated{}
}

/*UpdateConfigUsingPUTCreated handles this case with default header values.

Created
*/
type UpdateConfigUsingPUTCreated struct {
}

func (o *UpdateConfigUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/configparams/{name}][%d] updateConfigUsingPUTCreated ", 201)
}

func (o *UpdateConfigUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigUsingPUTUnauthorized creates a UpdateConfigUsingPUTUnauthorized with default headers values
func NewUpdateConfigUsingPUTUnauthorized() *UpdateConfigUsingPUTUnauthorized {
	return &UpdateConfigUsingPUTUnauthorized{}
}

/*UpdateConfigUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateConfigUsingPUTUnauthorized struct {
}

func (o *UpdateConfigUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/configparams/{name}][%d] updateConfigUsingPUTUnauthorized ", 401)
}

func (o *UpdateConfigUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigUsingPUTForbidden creates a UpdateConfigUsingPUTForbidden with default headers values
func NewUpdateConfigUsingPUTForbidden() *UpdateConfigUsingPUTForbidden {
	return &UpdateConfigUsingPUTForbidden{}
}

/*UpdateConfigUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type UpdateConfigUsingPUTForbidden struct {
}

func (o *UpdateConfigUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/configparams/{name}][%d] updateConfigUsingPUTForbidden ", 403)
}

func (o *UpdateConfigUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateConfigUsingPUTNotFound creates a UpdateConfigUsingPUTNotFound with default headers values
func NewUpdateConfigUsingPUTNotFound() *UpdateConfigUsingPUTNotFound {
	return &UpdateConfigUsingPUTNotFound{}
}

/*UpdateConfigUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdateConfigUsingPUTNotFound struct {
}

func (o *UpdateConfigUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/configparams/{name}][%d] updateConfigUsingPUTNotFound ", 404)
}

func (o *UpdateConfigUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
