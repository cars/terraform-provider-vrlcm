// Code generated by go-swagger; DO NOT EDIT.

package authzn_sample_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// IsVcfUserUsingGETReader is a Reader for the IsVcfUserUsingGET structure.
type IsVcfUserUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsVcfUserUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsVcfUserUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIsVcfUserUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIsVcfUserUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIsVcfUserUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsVcfUserUsingGETOK creates a IsVcfUserUsingGETOK with default headers values
func NewIsVcfUserUsingGETOK() *IsVcfUserUsingGETOK {
	return &IsVcfUserUsingGETOK{}
}

/*IsVcfUserUsingGETOK handles this case with default header values.

OK
*/
type IsVcfUserUsingGETOK struct {
	Payload *models.AuthenticatedUserDTO
}

func (o *IsVcfUserUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/vcfuser][%d] isVcfUserUsingGETOK  %+v", 200, o.Payload)
}

func (o *IsVcfUserUsingGETOK) GetPayload() *models.AuthenticatedUserDTO {
	return o.Payload
}

func (o *IsVcfUserUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticatedUserDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsVcfUserUsingGETUnauthorized creates a IsVcfUserUsingGETUnauthorized with default headers values
func NewIsVcfUserUsingGETUnauthorized() *IsVcfUserUsingGETUnauthorized {
	return &IsVcfUserUsingGETUnauthorized{}
}

/*IsVcfUserUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type IsVcfUserUsingGETUnauthorized struct {
}

func (o *IsVcfUserUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/vcfuser][%d] isVcfUserUsingGETUnauthorized ", 401)
}

func (o *IsVcfUserUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsVcfUserUsingGETForbidden creates a IsVcfUserUsingGETForbidden with default headers values
func NewIsVcfUserUsingGETForbidden() *IsVcfUserUsingGETForbidden {
	return &IsVcfUserUsingGETForbidden{}
}

/*IsVcfUserUsingGETForbidden handles this case with default header values.

Forbidden
*/
type IsVcfUserUsingGETForbidden struct {
}

func (o *IsVcfUserUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/vcfuser][%d] isVcfUserUsingGETForbidden ", 403)
}

func (o *IsVcfUserUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsVcfUserUsingGETNotFound creates a IsVcfUserUsingGETNotFound with default headers values
func NewIsVcfUserUsingGETNotFound() *IsVcfUserUsingGETNotFound {
	return &IsVcfUserUsingGETNotFound{}
}

/*IsVcfUserUsingGETNotFound handles this case with default header values.

Not Found
*/
type IsVcfUserUsingGETNotFound struct {
}

func (o *IsVcfUserUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/vcfuser][%d] isVcfUserUsingGETNotFound ", 404)
}

func (o *IsVcfUserUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
