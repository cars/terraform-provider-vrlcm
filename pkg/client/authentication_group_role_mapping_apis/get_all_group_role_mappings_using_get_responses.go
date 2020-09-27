// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_role_mapping_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllGroupRoleMappingsUsingGETReader is a Reader for the GetAllGroupRoleMappingsUsingGET structure.
type GetAllGroupRoleMappingsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllGroupRoleMappingsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllGroupRoleMappingsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllGroupRoleMappingsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllGroupRoleMappingsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllGroupRoleMappingsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllGroupRoleMappingsUsingGETOK creates a GetAllGroupRoleMappingsUsingGETOK with default headers values
func NewGetAllGroupRoleMappingsUsingGETOK() *GetAllGroupRoleMappingsUsingGETOK {
	return &GetAllGroupRoleMappingsUsingGETOK{}
}

/*GetAllGroupRoleMappingsUsingGETOK handles this case with default header values.

OK
*/
type GetAllGroupRoleMappingsUsingGETOK struct {
	Payload interface{}
}

func (o *GetAllGroupRoleMappingsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/grouprolemapping][%d] getAllGroupRoleMappingsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllGroupRoleMappingsUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAllGroupRoleMappingsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllGroupRoleMappingsUsingGETUnauthorized creates a GetAllGroupRoleMappingsUsingGETUnauthorized with default headers values
func NewGetAllGroupRoleMappingsUsingGETUnauthorized() *GetAllGroupRoleMappingsUsingGETUnauthorized {
	return &GetAllGroupRoleMappingsUsingGETUnauthorized{}
}

/*GetAllGroupRoleMappingsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllGroupRoleMappingsUsingGETUnauthorized struct {
}

func (o *GetAllGroupRoleMappingsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/grouprolemapping][%d] getAllGroupRoleMappingsUsingGETUnauthorized ", 401)
}

func (o *GetAllGroupRoleMappingsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGroupRoleMappingsUsingGETForbidden creates a GetAllGroupRoleMappingsUsingGETForbidden with default headers values
func NewGetAllGroupRoleMappingsUsingGETForbidden() *GetAllGroupRoleMappingsUsingGETForbidden {
	return &GetAllGroupRoleMappingsUsingGETForbidden{}
}

/*GetAllGroupRoleMappingsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllGroupRoleMappingsUsingGETForbidden struct {
}

func (o *GetAllGroupRoleMappingsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/grouprolemapping][%d] getAllGroupRoleMappingsUsingGETForbidden ", 403)
}

func (o *GetAllGroupRoleMappingsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllGroupRoleMappingsUsingGETNotFound creates a GetAllGroupRoleMappingsUsingGETNotFound with default headers values
func NewGetAllGroupRoleMappingsUsingGETNotFound() *GetAllGroupRoleMappingsUsingGETNotFound {
	return &GetAllGroupRoleMappingsUsingGETNotFound{}
}

/*GetAllGroupRoleMappingsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllGroupRoleMappingsUsingGETNotFound struct {
}

func (o *GetAllGroupRoleMappingsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/grouprolemapping][%d] getAllGroupRoleMappingsUsingGETNotFound ", 404)
}

func (o *GetAllGroupRoleMappingsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
