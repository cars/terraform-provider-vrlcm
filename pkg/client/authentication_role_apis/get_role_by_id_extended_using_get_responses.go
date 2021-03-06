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

// GetRoleByIDExtendedUsingGETReader is a Reader for the GetRoleByIDExtendedUsingGET structure.
type GetRoleByIDExtendedUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleByIDExtendedUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoleByIDExtendedUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRoleByIDExtendedUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoleByIDExtendedUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoleByIDExtendedUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoleByIDExtendedUsingGETOK creates a GetRoleByIDExtendedUsingGETOK with default headers values
func NewGetRoleByIDExtendedUsingGETOK() *GetRoleByIDExtendedUsingGETOK {
	return &GetRoleByIDExtendedUsingGETOK{}
}

/*GetRoleByIDExtendedUsingGETOK handles this case with default header values.

OK
*/
type GetRoleByIDExtendedUsingGETOK struct {
	Payload interface{}
}

func (o *GetRoleByIDExtendedUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/{vmid}/extended][%d] getRoleByIdExtendedUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetRoleByIDExtendedUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetRoleByIDExtendedUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleByIDExtendedUsingGETUnauthorized creates a GetRoleByIDExtendedUsingGETUnauthorized with default headers values
func NewGetRoleByIDExtendedUsingGETUnauthorized() *GetRoleByIDExtendedUsingGETUnauthorized {
	return &GetRoleByIDExtendedUsingGETUnauthorized{}
}

/*GetRoleByIDExtendedUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRoleByIDExtendedUsingGETUnauthorized struct {
}

func (o *GetRoleByIDExtendedUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/{vmid}/extended][%d] getRoleByIdExtendedUsingGETUnauthorized ", 401)
}

func (o *GetRoleByIDExtendedUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleByIDExtendedUsingGETForbidden creates a GetRoleByIDExtendedUsingGETForbidden with default headers values
func NewGetRoleByIDExtendedUsingGETForbidden() *GetRoleByIDExtendedUsingGETForbidden {
	return &GetRoleByIDExtendedUsingGETForbidden{}
}

/*GetRoleByIDExtendedUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetRoleByIDExtendedUsingGETForbidden struct {
}

func (o *GetRoleByIDExtendedUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/{vmid}/extended][%d] getRoleByIdExtendedUsingGETForbidden ", 403)
}

func (o *GetRoleByIDExtendedUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRoleByIDExtendedUsingGETNotFound creates a GetRoleByIDExtendedUsingGETNotFound with default headers values
func NewGetRoleByIDExtendedUsingGETNotFound() *GetRoleByIDExtendedUsingGETNotFound {
	return &GetRoleByIDExtendedUsingGETNotFound{}
}

/*GetRoleByIDExtendedUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetRoleByIDExtendedUsingGETNotFound struct {
}

func (o *GetRoleByIDExtendedUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/{vmid}/extended][%d] getRoleByIdExtendedUsingGETNotFound ", 404)
}

func (o *GetRoleByIDExtendedUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
