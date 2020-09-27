// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateGroupByIDUsingPUTReader is a Reader for the UpdateGroupByIDUsingPUT structure.
type UpdateGroupByIDUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGroupByIDUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGroupByIDUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateGroupByIDUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateGroupByIDUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateGroupByIDUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateGroupByIDUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGroupByIDUsingPUTOK creates a UpdateGroupByIDUsingPUTOK with default headers values
func NewUpdateGroupByIDUsingPUTOK() *UpdateGroupByIDUsingPUTOK {
	return &UpdateGroupByIDUsingPUTOK{}
}

/*UpdateGroupByIDUsingPUTOK handles this case with default header values.

OK
*/
type UpdateGroupByIDUsingPUTOK struct {
	Payload interface{}
}

func (o *UpdateGroupByIDUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/groups/{vmid}][%d] updateGroupByIdUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateGroupByIDUsingPUTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateGroupByIDUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGroupByIDUsingPUTCreated creates a UpdateGroupByIDUsingPUTCreated with default headers values
func NewUpdateGroupByIDUsingPUTCreated() *UpdateGroupByIDUsingPUTCreated {
	return &UpdateGroupByIDUsingPUTCreated{}
}

/*UpdateGroupByIDUsingPUTCreated handles this case with default header values.

Created
*/
type UpdateGroupByIDUsingPUTCreated struct {
}

func (o *UpdateGroupByIDUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/groups/{vmid}][%d] updateGroupByIdUsingPUTCreated ", 201)
}

func (o *UpdateGroupByIDUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupByIDUsingPUTUnauthorized creates a UpdateGroupByIDUsingPUTUnauthorized with default headers values
func NewUpdateGroupByIDUsingPUTUnauthorized() *UpdateGroupByIDUsingPUTUnauthorized {
	return &UpdateGroupByIDUsingPUTUnauthorized{}
}

/*UpdateGroupByIDUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateGroupByIDUsingPUTUnauthorized struct {
}

func (o *UpdateGroupByIDUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/groups/{vmid}][%d] updateGroupByIdUsingPUTUnauthorized ", 401)
}

func (o *UpdateGroupByIDUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupByIDUsingPUTForbidden creates a UpdateGroupByIDUsingPUTForbidden with default headers values
func NewUpdateGroupByIDUsingPUTForbidden() *UpdateGroupByIDUsingPUTForbidden {
	return &UpdateGroupByIDUsingPUTForbidden{}
}

/*UpdateGroupByIDUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type UpdateGroupByIDUsingPUTForbidden struct {
}

func (o *UpdateGroupByIDUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/groups/{vmid}][%d] updateGroupByIdUsingPUTForbidden ", 403)
}

func (o *UpdateGroupByIDUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGroupByIDUsingPUTNotFound creates a UpdateGroupByIDUsingPUTNotFound with default headers values
func NewUpdateGroupByIDUsingPUTNotFound() *UpdateGroupByIDUsingPUTNotFound {
	return &UpdateGroupByIDUsingPUTNotFound{}
}

/*UpdateGroupByIDUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdateGroupByIDUsingPUTNotFound struct {
}

func (o *UpdateGroupByIDUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /lcm/authzn/api/groups/{vmid}][%d] updateGroupByIdUsingPUTNotFound ", 404)
}

func (o *UpdateGroupByIDUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}