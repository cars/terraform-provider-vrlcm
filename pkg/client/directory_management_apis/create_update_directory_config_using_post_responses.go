// Code generated by go-swagger; DO NOT EDIT.

package directory_management_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateUpdateDirectoryConfigUsingPOSTReader is a Reader for the CreateUpdateDirectoryConfigUsingPOST structure.
type CreateUpdateDirectoryConfigUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUpdateDirectoryConfigUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUpdateDirectoryConfigUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateUpdateDirectoryConfigUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUpdateDirectoryConfigUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUpdateDirectoryConfigUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUpdateDirectoryConfigUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUpdateDirectoryConfigUsingPOSTOK creates a CreateUpdateDirectoryConfigUsingPOSTOK with default headers values
func NewCreateUpdateDirectoryConfigUsingPOSTOK() *CreateUpdateDirectoryConfigUsingPOSTOK {
	return &CreateUpdateDirectoryConfigUsingPOSTOK{}
}

/*CreateUpdateDirectoryConfigUsingPOSTOK handles this case with default header values.

OK
*/
type CreateUpdateDirectoryConfigUsingPOSTOK struct {
	Payload interface{}
}

func (o *CreateUpdateDirectoryConfigUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/idp/dirConfigs][%d] createUpdateDirectoryConfigUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateUpdateDirectoryConfigUsingPOSTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUpdateDirectoryConfigUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUpdateDirectoryConfigUsingPOSTCreated creates a CreateUpdateDirectoryConfigUsingPOSTCreated with default headers values
func NewCreateUpdateDirectoryConfigUsingPOSTCreated() *CreateUpdateDirectoryConfigUsingPOSTCreated {
	return &CreateUpdateDirectoryConfigUsingPOSTCreated{}
}

/*CreateUpdateDirectoryConfigUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateUpdateDirectoryConfigUsingPOSTCreated struct {
}

func (o *CreateUpdateDirectoryConfigUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/idp/dirConfigs][%d] createUpdateDirectoryConfigUsingPOSTCreated ", 201)
}

func (o *CreateUpdateDirectoryConfigUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUpdateDirectoryConfigUsingPOSTUnauthorized creates a CreateUpdateDirectoryConfigUsingPOSTUnauthorized with default headers values
func NewCreateUpdateDirectoryConfigUsingPOSTUnauthorized() *CreateUpdateDirectoryConfigUsingPOSTUnauthorized {
	return &CreateUpdateDirectoryConfigUsingPOSTUnauthorized{}
}

/*CreateUpdateDirectoryConfigUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateUpdateDirectoryConfigUsingPOSTUnauthorized struct {
}

func (o *CreateUpdateDirectoryConfigUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/idp/dirConfigs][%d] createUpdateDirectoryConfigUsingPOSTUnauthorized ", 401)
}

func (o *CreateUpdateDirectoryConfigUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUpdateDirectoryConfigUsingPOSTForbidden creates a CreateUpdateDirectoryConfigUsingPOSTForbidden with default headers values
func NewCreateUpdateDirectoryConfigUsingPOSTForbidden() *CreateUpdateDirectoryConfigUsingPOSTForbidden {
	return &CreateUpdateDirectoryConfigUsingPOSTForbidden{}
}

/*CreateUpdateDirectoryConfigUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateUpdateDirectoryConfigUsingPOSTForbidden struct {
}

func (o *CreateUpdateDirectoryConfigUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/idp/dirConfigs][%d] createUpdateDirectoryConfigUsingPOSTForbidden ", 403)
}

func (o *CreateUpdateDirectoryConfigUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUpdateDirectoryConfigUsingPOSTNotFound creates a CreateUpdateDirectoryConfigUsingPOSTNotFound with default headers values
func NewCreateUpdateDirectoryConfigUsingPOSTNotFound() *CreateUpdateDirectoryConfigUsingPOSTNotFound {
	return &CreateUpdateDirectoryConfigUsingPOSTNotFound{}
}

/*CreateUpdateDirectoryConfigUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateUpdateDirectoryConfigUsingPOSTNotFound struct {
}

func (o *CreateUpdateDirectoryConfigUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/idp/dirConfigs][%d] createUpdateDirectoryConfigUsingPOSTNotFound ", 404)
}

func (o *CreateUpdateDirectoryConfigUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}