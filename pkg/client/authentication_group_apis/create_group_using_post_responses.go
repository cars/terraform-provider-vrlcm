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

// CreateGroupUsingPOSTReader is a Reader for the CreateGroupUsingPOST structure.
type CreateGroupUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGroupUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGroupUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateGroupUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateGroupUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGroupUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGroupUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateGroupUsingPOSTOK creates a CreateGroupUsingPOSTOK with default headers values
func NewCreateGroupUsingPOSTOK() *CreateGroupUsingPOSTOK {
	return &CreateGroupUsingPOSTOK{}
}

/*CreateGroupUsingPOSTOK handles this case with default header values.

OK
*/
type CreateGroupUsingPOSTOK struct {
	Payload interface{}
}

func (o *CreateGroupUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/groups][%d] createGroupUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGroupUsingPOSTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateGroupUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupUsingPOSTCreated creates a CreateGroupUsingPOSTCreated with default headers values
func NewCreateGroupUsingPOSTCreated() *CreateGroupUsingPOSTCreated {
	return &CreateGroupUsingPOSTCreated{}
}

/*CreateGroupUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateGroupUsingPOSTCreated struct {
}

func (o *CreateGroupUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/groups][%d] createGroupUsingPOSTCreated ", 201)
}

func (o *CreateGroupUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGroupUsingPOSTUnauthorized creates a CreateGroupUsingPOSTUnauthorized with default headers values
func NewCreateGroupUsingPOSTUnauthorized() *CreateGroupUsingPOSTUnauthorized {
	return &CreateGroupUsingPOSTUnauthorized{}
}

/*CreateGroupUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateGroupUsingPOSTUnauthorized struct {
}

func (o *CreateGroupUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/groups][%d] createGroupUsingPOSTUnauthorized ", 401)
}

func (o *CreateGroupUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGroupUsingPOSTForbidden creates a CreateGroupUsingPOSTForbidden with default headers values
func NewCreateGroupUsingPOSTForbidden() *CreateGroupUsingPOSTForbidden {
	return &CreateGroupUsingPOSTForbidden{}
}

/*CreateGroupUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateGroupUsingPOSTForbidden struct {
}

func (o *CreateGroupUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/groups][%d] createGroupUsingPOSTForbidden ", 403)
}

func (o *CreateGroupUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGroupUsingPOSTNotFound creates a CreateGroupUsingPOSTNotFound with default headers values
func NewCreateGroupUsingPOSTNotFound() *CreateGroupUsingPOSTNotFound {
	return &CreateGroupUsingPOSTNotFound{}
}

/*CreateGroupUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateGroupUsingPOSTNotFound struct {
}

func (o *CreateGroupUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/groups][%d] createGroupUsingPOSTNotFound ", 404)
}

func (o *CreateGroupUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
