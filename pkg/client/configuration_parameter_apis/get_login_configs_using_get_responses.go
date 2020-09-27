// Code generated by go-swagger; DO NOT EDIT.

package configuration_parameter_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLoginConfigsUsingGETReader is a Reader for the GetLoginConfigsUsingGET structure.
type GetLoginConfigsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoginConfigsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoginConfigsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLoginConfigsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLoginConfigsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLoginConfigsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLoginConfigsUsingGETOK creates a GetLoginConfigsUsingGETOK with default headers values
func NewGetLoginConfigsUsingGETOK() *GetLoginConfigsUsingGETOK {
	return &GetLoginConfigsUsingGETOK{}
}

/*GetLoginConfigsUsingGETOK handles this case with default header values.

OK
*/
type GetLoginConfigsUsingGETOK struct {
	Payload interface{}
}

func (o *GetLoginConfigsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/loginconfig][%d] getLoginConfigsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetLoginConfigsUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLoginConfigsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoginConfigsUsingGETUnauthorized creates a GetLoginConfigsUsingGETUnauthorized with default headers values
func NewGetLoginConfigsUsingGETUnauthorized() *GetLoginConfigsUsingGETUnauthorized {
	return &GetLoginConfigsUsingGETUnauthorized{}
}

/*GetLoginConfigsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetLoginConfigsUsingGETUnauthorized struct {
}

func (o *GetLoginConfigsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/loginconfig][%d] getLoginConfigsUsingGETUnauthorized ", 401)
}

func (o *GetLoginConfigsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoginConfigsUsingGETForbidden creates a GetLoginConfigsUsingGETForbidden with default headers values
func NewGetLoginConfigsUsingGETForbidden() *GetLoginConfigsUsingGETForbidden {
	return &GetLoginConfigsUsingGETForbidden{}
}

/*GetLoginConfigsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetLoginConfigsUsingGETForbidden struct {
}

func (o *GetLoginConfigsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/loginconfig][%d] getLoginConfigsUsingGETForbidden ", 403)
}

func (o *GetLoginConfigsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLoginConfigsUsingGETNotFound creates a GetLoginConfigsUsingGETNotFound with default headers values
func NewGetLoginConfigsUsingGETNotFound() *GetLoginConfigsUsingGETNotFound {
	return &GetLoginConfigsUsingGETNotFound{}
}

/*GetLoginConfigsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetLoginConfigsUsingGETNotFound struct {
}

func (o *GetLoginConfigsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/loginconfig][%d] getLoginConfigsUsingGETNotFound ", 404)
}

func (o *GetLoginConfigsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
