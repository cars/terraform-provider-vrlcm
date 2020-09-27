// Code generated by go-swagger; DO NOT EDIT.

package authentication_user_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateUserByIDUsingPUTReader is a Reader for the UpdateUserByIDUsingPUT structure.
type UpdateUserByIDUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserByIDUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserByIDUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateUserByIDUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUserByIDUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserByIDUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserByIDUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserByIDUsingPUTOK creates a UpdateUserByIDUsingPUTOK with default headers values
func NewUpdateUserByIDUsingPUTOK() *UpdateUserByIDUsingPUTOK {
	return &UpdateUserByIDUsingPUTOK{}
}

/*UpdateUserByIDUsingPUTOK handles this case with default header values.

OK
*/
type UpdateUserByIDUsingPUTOK struct {
	Payload interface{}
}

func (o *UpdateUserByIDUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/users/{vmid}][%d] updateUserByIdUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateUserByIDUsingPUTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateUserByIDUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserByIDUsingPUTCreated creates a UpdateUserByIDUsingPUTCreated with default headers values
func NewUpdateUserByIDUsingPUTCreated() *UpdateUserByIDUsingPUTCreated {
	return &UpdateUserByIDUsingPUTCreated{}
}

/*UpdateUserByIDUsingPUTCreated handles this case with default header values.

Created
*/
type UpdateUserByIDUsingPUTCreated struct {
}

func (o *UpdateUserByIDUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/users/{vmid}][%d] updateUserByIdUsingPUTCreated ", 201)
}

func (o *UpdateUserByIDUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserByIDUsingPUTUnauthorized creates a UpdateUserByIDUsingPUTUnauthorized with default headers values
func NewUpdateUserByIDUsingPUTUnauthorized() *UpdateUserByIDUsingPUTUnauthorized {
	return &UpdateUserByIDUsingPUTUnauthorized{}
}

/*UpdateUserByIDUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUserByIDUsingPUTUnauthorized struct {
}

func (o *UpdateUserByIDUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/users/{vmid}][%d] updateUserByIdUsingPUTUnauthorized ", 401)
}

func (o *UpdateUserByIDUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserByIDUsingPUTForbidden creates a UpdateUserByIDUsingPUTForbidden with default headers values
func NewUpdateUserByIDUsingPUTForbidden() *UpdateUserByIDUsingPUTForbidden {
	return &UpdateUserByIDUsingPUTForbidden{}
}

/*UpdateUserByIDUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type UpdateUserByIDUsingPUTForbidden struct {
}

func (o *UpdateUserByIDUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/users/{vmid}][%d] updateUserByIdUsingPUTForbidden ", 403)
}

func (o *UpdateUserByIDUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserByIDUsingPUTNotFound creates a UpdateUserByIDUsingPUTNotFound with default headers values
func NewUpdateUserByIDUsingPUTNotFound() *UpdateUserByIDUsingPUTNotFound {
	return &UpdateUserByIDUsingPUTNotFound{}
}

/*UpdateUserByIDUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdateUserByIDUsingPUTNotFound struct {
}

func (o *UpdateUserByIDUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/users/{vmid}][%d] updateUserByIdUsingPUTNotFound ", 404)
}

func (o *UpdateUserByIDUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
