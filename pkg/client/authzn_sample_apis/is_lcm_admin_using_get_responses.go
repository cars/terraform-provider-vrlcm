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

// IsLcmAdminUsingGETReader is a Reader for the IsLcmAdminUsingGET structure.
type IsLcmAdminUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsLcmAdminUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsLcmAdminUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIsLcmAdminUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIsLcmAdminUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIsLcmAdminUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsLcmAdminUsingGETOK creates a IsLcmAdminUsingGETOK with default headers values
func NewIsLcmAdminUsingGETOK() *IsLcmAdminUsingGETOK {
	return &IsLcmAdminUsingGETOK{}
}

/*IsLcmAdminUsingGETOK handles this case with default header values.

OK
*/
type IsLcmAdminUsingGETOK struct {
	Payload *models.AuthenticatedUserDTO
}

func (o *IsLcmAdminUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/lcmadmin][%d] isLcmAdminUsingGETOK  %+v", 200, o.Payload)
}

func (o *IsLcmAdminUsingGETOK) GetPayload() *models.AuthenticatedUserDTO {
	return o.Payload
}

func (o *IsLcmAdminUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticatedUserDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsLcmAdminUsingGETUnauthorized creates a IsLcmAdminUsingGETUnauthorized with default headers values
func NewIsLcmAdminUsingGETUnauthorized() *IsLcmAdminUsingGETUnauthorized {
	return &IsLcmAdminUsingGETUnauthorized{}
}

/*IsLcmAdminUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type IsLcmAdminUsingGETUnauthorized struct {
}

func (o *IsLcmAdminUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/lcmadmin][%d] isLcmAdminUsingGETUnauthorized ", 401)
}

func (o *IsLcmAdminUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsLcmAdminUsingGETForbidden creates a IsLcmAdminUsingGETForbidden with default headers values
func NewIsLcmAdminUsingGETForbidden() *IsLcmAdminUsingGETForbidden {
	return &IsLcmAdminUsingGETForbidden{}
}

/*IsLcmAdminUsingGETForbidden handles this case with default header values.

Forbidden
*/
type IsLcmAdminUsingGETForbidden struct {
}

func (o *IsLcmAdminUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/lcmadmin][%d] isLcmAdminUsingGETForbidden ", 403)
}

func (o *IsLcmAdminUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsLcmAdminUsingGETNotFound creates a IsLcmAdminUsingGETNotFound with default headers values
func NewIsLcmAdminUsingGETNotFound() *IsLcmAdminUsingGETNotFound {
	return &IsLcmAdminUsingGETNotFound{}
}

/*IsLcmAdminUsingGETNotFound handles this case with default header values.

Not Found
*/
type IsLcmAdminUsingGETNotFound struct {
}

func (o *IsLcmAdminUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/lcmadmin][%d] isLcmAdminUsingGETNotFound ", 404)
}

func (o *IsLcmAdminUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
