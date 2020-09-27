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

// GetAllExternalRolesUsingGETReader is a Reader for the GetAllExternalRolesUsingGET structure.
type GetAllExternalRolesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllExternalRolesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllExternalRolesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllExternalRolesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllExternalRolesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllExternalRolesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllExternalRolesUsingGETOK creates a GetAllExternalRolesUsingGETOK with default headers values
func NewGetAllExternalRolesUsingGETOK() *GetAllExternalRolesUsingGETOK {
	return &GetAllExternalRolesUsingGETOK{}
}

/*GetAllExternalRolesUsingGETOK handles this case with default header values.

OK
*/
type GetAllExternalRolesUsingGETOK struct {
	Payload interface{}
}

func (o *GetAllExternalRolesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/external][%d] getAllExternalRolesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllExternalRolesUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAllExternalRolesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllExternalRolesUsingGETUnauthorized creates a GetAllExternalRolesUsingGETUnauthorized with default headers values
func NewGetAllExternalRolesUsingGETUnauthorized() *GetAllExternalRolesUsingGETUnauthorized {
	return &GetAllExternalRolesUsingGETUnauthorized{}
}

/*GetAllExternalRolesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllExternalRolesUsingGETUnauthorized struct {
}

func (o *GetAllExternalRolesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/external][%d] getAllExternalRolesUsingGETUnauthorized ", 401)
}

func (o *GetAllExternalRolesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllExternalRolesUsingGETForbidden creates a GetAllExternalRolesUsingGETForbidden with default headers values
func NewGetAllExternalRolesUsingGETForbidden() *GetAllExternalRolesUsingGETForbidden {
	return &GetAllExternalRolesUsingGETForbidden{}
}

/*GetAllExternalRolesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllExternalRolesUsingGETForbidden struct {
}

func (o *GetAllExternalRolesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/external][%d] getAllExternalRolesUsingGETForbidden ", 403)
}

func (o *GetAllExternalRolesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllExternalRolesUsingGETNotFound creates a GetAllExternalRolesUsingGETNotFound with default headers values
func NewGetAllExternalRolesUsingGETNotFound() *GetAllExternalRolesUsingGETNotFound {
	return &GetAllExternalRolesUsingGETNotFound{}
}

/*GetAllExternalRolesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllExternalRolesUsingGETNotFound struct {
}

func (o *GetAllExternalRolesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/roles/external][%d] getAllExternalRolesUsingGETNotFound ", 404)
}

func (o *GetAllExternalRolesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
