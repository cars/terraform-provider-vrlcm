// Code generated by go-swagger; DO NOT EDIT.

package authzn_login_logout_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// LoginBasicUsingPOSTReader is a Reader for the LoginBasicUsingPOST structure.
type LoginBasicUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginBasicUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginBasicUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewLoginBasicUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginBasicUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLoginBasicUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewLoginBasicUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewLoginBasicUsingPOSTOK creates a LoginBasicUsingPOSTOK with default headers values
func NewLoginBasicUsingPOSTOK() *LoginBasicUsingPOSTOK {
	return &LoginBasicUsingPOSTOK{}
}

/*LoginBasicUsingPOSTOK handles this case with default header values.

OK
*/
type LoginBasicUsingPOSTOK struct {
	Payload string
}

func (o *LoginBasicUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/login][%d] loginBasicUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *LoginBasicUsingPOSTOK) GetPayload() string {
	return o.Payload
}

func (o *LoginBasicUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginBasicUsingPOSTCreated creates a LoginBasicUsingPOSTCreated with default headers values
func NewLoginBasicUsingPOSTCreated() *LoginBasicUsingPOSTCreated {
	return &LoginBasicUsingPOSTCreated{}
}

/*LoginBasicUsingPOSTCreated handles this case with default header values.

Created
*/
type LoginBasicUsingPOSTCreated struct {
}

func (o *LoginBasicUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/login][%d] loginBasicUsingPOSTCreated ", 201)
}

func (o *LoginBasicUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginBasicUsingPOSTUnauthorized creates a LoginBasicUsingPOSTUnauthorized with default headers values
func NewLoginBasicUsingPOSTUnauthorized() *LoginBasicUsingPOSTUnauthorized {
	return &LoginBasicUsingPOSTUnauthorized{}
}

/*LoginBasicUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type LoginBasicUsingPOSTUnauthorized struct {
}

func (o *LoginBasicUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/login][%d] loginBasicUsingPOSTUnauthorized ", 401)
}

func (o *LoginBasicUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginBasicUsingPOSTForbidden creates a LoginBasicUsingPOSTForbidden with default headers values
func NewLoginBasicUsingPOSTForbidden() *LoginBasicUsingPOSTForbidden {
	return &LoginBasicUsingPOSTForbidden{}
}

/*LoginBasicUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type LoginBasicUsingPOSTForbidden struct {
}

func (o *LoginBasicUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/login][%d] loginBasicUsingPOSTForbidden ", 403)
}

func (o *LoginBasicUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginBasicUsingPOSTNotFound creates a LoginBasicUsingPOSTNotFound with default headers values
func NewLoginBasicUsingPOSTNotFound() *LoginBasicUsingPOSTNotFound {
	return &LoginBasicUsingPOSTNotFound{}
}

/*LoginBasicUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type LoginBasicUsingPOSTNotFound struct {
}

func (o *LoginBasicUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/authzn/api/login][%d] loginBasicUsingPOSTNotFound ", 404)
}

func (o *LoginBasicUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
