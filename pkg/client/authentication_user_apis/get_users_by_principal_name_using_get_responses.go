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

// GetUsersByPrincipalNameUsingGETReader is a Reader for the GetUsersByPrincipalNameUsingGET structure.
type GetUsersByPrincipalNameUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersByPrincipalNameUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersByPrincipalNameUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersByPrincipalNameUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersByPrincipalNameUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersByPrincipalNameUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersByPrincipalNameUsingGETOK creates a GetUsersByPrincipalNameUsingGETOK with default headers values
func NewGetUsersByPrincipalNameUsingGETOK() *GetUsersByPrincipalNameUsingGETOK {
	return &GetUsersByPrincipalNameUsingGETOK{}
}

/*GetUsersByPrincipalNameUsingGETOK handles this case with default header values.

OK
*/
type GetUsersByPrincipalNameUsingGETOK struct {
	Payload interface{}
}

func (o *GetUsersByPrincipalNameUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/users/upn/{upn}][%d] getUsersByPrincipalNameUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetUsersByPrincipalNameUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetUsersByPrincipalNameUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersByPrincipalNameUsingGETUnauthorized creates a GetUsersByPrincipalNameUsingGETUnauthorized with default headers values
func NewGetUsersByPrincipalNameUsingGETUnauthorized() *GetUsersByPrincipalNameUsingGETUnauthorized {
	return &GetUsersByPrincipalNameUsingGETUnauthorized{}
}

/*GetUsersByPrincipalNameUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUsersByPrincipalNameUsingGETUnauthorized struct {
}

func (o *GetUsersByPrincipalNameUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/users/upn/{upn}][%d] getUsersByPrincipalNameUsingGETUnauthorized ", 401)
}

func (o *GetUsersByPrincipalNameUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersByPrincipalNameUsingGETForbidden creates a GetUsersByPrincipalNameUsingGETForbidden with default headers values
func NewGetUsersByPrincipalNameUsingGETForbidden() *GetUsersByPrincipalNameUsingGETForbidden {
	return &GetUsersByPrincipalNameUsingGETForbidden{}
}

/*GetUsersByPrincipalNameUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetUsersByPrincipalNameUsingGETForbidden struct {
}

func (o *GetUsersByPrincipalNameUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/users/upn/{upn}][%d] getUsersByPrincipalNameUsingGETForbidden ", 403)
}

func (o *GetUsersByPrincipalNameUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersByPrincipalNameUsingGETNotFound creates a GetUsersByPrincipalNameUsingGETNotFound with default headers values
func NewGetUsersByPrincipalNameUsingGETNotFound() *GetUsersByPrincipalNameUsingGETNotFound {
	return &GetUsersByPrincipalNameUsingGETNotFound{}
}

/*GetUsersByPrincipalNameUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetUsersByPrincipalNameUsingGETNotFound struct {
}

func (o *GetUsersByPrincipalNameUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/users/upn/{upn}][%d] getUsersByPrincipalNameUsingGETNotFound ", 404)
}

func (o *GetUsersByPrincipalNameUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}