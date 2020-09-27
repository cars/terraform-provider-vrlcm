// Code generated by go-swagger; DO NOT EDIT.

package authentication_role_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllRolesUsingGETReader is a Reader for the GetAllRolesUsingGET structure.
type GetAllRolesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllRolesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllRolesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllRolesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllRolesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllRolesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllRolesUsingGETOK creates a GetAllRolesUsingGETOK with default headers values
func NewGetAllRolesUsingGETOK() *GetAllRolesUsingGETOK {
	return &GetAllRolesUsingGETOK{}
}

/*GetAllRolesUsingGETOK handles this case with default header values.

OK
*/
type GetAllRolesUsingGETOK struct {
	Payload interface{}
}

func (o *GetAllRolesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles][%d] getAllRolesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllRolesUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAllRolesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllRolesUsingGETUnauthorized creates a GetAllRolesUsingGETUnauthorized with default headers values
func NewGetAllRolesUsingGETUnauthorized() *GetAllRolesUsingGETUnauthorized {
	return &GetAllRolesUsingGETUnauthorized{}
}

/*GetAllRolesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllRolesUsingGETUnauthorized struct {
}

func (o *GetAllRolesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles][%d] getAllRolesUsingGETUnauthorized ", 401)
}

func (o *GetAllRolesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllRolesUsingGETForbidden creates a GetAllRolesUsingGETForbidden with default headers values
func NewGetAllRolesUsingGETForbidden() *GetAllRolesUsingGETForbidden {
	return &GetAllRolesUsingGETForbidden{}
}

/*GetAllRolesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllRolesUsingGETForbidden struct {
}

func (o *GetAllRolesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles][%d] getAllRolesUsingGETForbidden ", 403)
}

func (o *GetAllRolesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllRolesUsingGETNotFound creates a GetAllRolesUsingGETNotFound with default headers values
func NewGetAllRolesUsingGETNotFound() *GetAllRolesUsingGETNotFound {
	return &GetAllRolesUsingGETNotFound{}
}

/*GetAllRolesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllRolesUsingGETNotFound struct {
}

func (o *GetAllRolesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles][%d] getAllRolesUsingGETNotFound ", 404)
}

func (o *GetAllRolesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
