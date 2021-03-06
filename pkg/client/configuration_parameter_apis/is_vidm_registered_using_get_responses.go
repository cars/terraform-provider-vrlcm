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

// IsVidmRegisteredUsingGETReader is a Reader for the IsVidmRegisteredUsingGET structure.
type IsVidmRegisteredUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsVidmRegisteredUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsVidmRegisteredUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIsVidmRegisteredUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIsVidmRegisteredUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIsVidmRegisteredUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsVidmRegisteredUsingGETOK creates a IsVidmRegisteredUsingGETOK with default headers values
func NewIsVidmRegisteredUsingGETOK() *IsVidmRegisteredUsingGETOK {
	return &IsVidmRegisteredUsingGETOK{}
}

/*IsVidmRegisteredUsingGETOK handles this case with default header values.

OK
*/
type IsVidmRegisteredUsingGETOK struct {
	Payload interface{}
}

func (o *IsVidmRegisteredUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/isvidmregistered][%d] isVidmRegisteredUsingGETOK  %+v", 200, o.Payload)
}

func (o *IsVidmRegisteredUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *IsVidmRegisteredUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsVidmRegisteredUsingGETUnauthorized creates a IsVidmRegisteredUsingGETUnauthorized with default headers values
func NewIsVidmRegisteredUsingGETUnauthorized() *IsVidmRegisteredUsingGETUnauthorized {
	return &IsVidmRegisteredUsingGETUnauthorized{}
}

/*IsVidmRegisteredUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type IsVidmRegisteredUsingGETUnauthorized struct {
}

func (o *IsVidmRegisteredUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/isvidmregistered][%d] isVidmRegisteredUsingGETUnauthorized ", 401)
}

func (o *IsVidmRegisteredUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsVidmRegisteredUsingGETForbidden creates a IsVidmRegisteredUsingGETForbidden with default headers values
func NewIsVidmRegisteredUsingGETForbidden() *IsVidmRegisteredUsingGETForbidden {
	return &IsVidmRegisteredUsingGETForbidden{}
}

/*IsVidmRegisteredUsingGETForbidden handles this case with default header values.

Forbidden
*/
type IsVidmRegisteredUsingGETForbidden struct {
}

func (o *IsVidmRegisteredUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/isvidmregistered][%d] isVidmRegisteredUsingGETForbidden ", 403)
}

func (o *IsVidmRegisteredUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsVidmRegisteredUsingGETNotFound creates a IsVidmRegisteredUsingGETNotFound with default headers values
func NewIsVidmRegisteredUsingGETNotFound() *IsVidmRegisteredUsingGETNotFound {
	return &IsVidmRegisteredUsingGETNotFound{}
}

/*IsVidmRegisteredUsingGETNotFound handles this case with default header values.

Not Found
*/
type IsVidmRegisteredUsingGETNotFound struct {
}

func (o *IsVidmRegisteredUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/isvidmregistered][%d] isVidmRegisteredUsingGETNotFound ", 404)
}

func (o *IsVidmRegisteredUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
