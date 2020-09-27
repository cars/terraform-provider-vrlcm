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

// IsContentDeveloperUsingGETReader is a Reader for the IsContentDeveloperUsingGET structure.
type IsContentDeveloperUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IsContentDeveloperUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIsContentDeveloperUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIsContentDeveloperUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIsContentDeveloperUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIsContentDeveloperUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIsContentDeveloperUsingGETOK creates a IsContentDeveloperUsingGETOK with default headers values
func NewIsContentDeveloperUsingGETOK() *IsContentDeveloperUsingGETOK {
	return &IsContentDeveloperUsingGETOK{}
}

/*IsContentDeveloperUsingGETOK handles this case with default header values.

OK
*/
type IsContentDeveloperUsingGETOK struct {
	Payload *models.AuthenticatedUserDTO
}

func (o *IsContentDeveloperUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/contentdev][%d] isContentDeveloperUsingGETOK  %+v", 200, o.Payload)
}

func (o *IsContentDeveloperUsingGETOK) GetPayload() *models.AuthenticatedUserDTO {
	return o.Payload
}

func (o *IsContentDeveloperUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticatedUserDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIsContentDeveloperUsingGETUnauthorized creates a IsContentDeveloperUsingGETUnauthorized with default headers values
func NewIsContentDeveloperUsingGETUnauthorized() *IsContentDeveloperUsingGETUnauthorized {
	return &IsContentDeveloperUsingGETUnauthorized{}
}

/*IsContentDeveloperUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type IsContentDeveloperUsingGETUnauthorized struct {
}

func (o *IsContentDeveloperUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/contentdev][%d] isContentDeveloperUsingGETUnauthorized ", 401)
}

func (o *IsContentDeveloperUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsContentDeveloperUsingGETForbidden creates a IsContentDeveloperUsingGETForbidden with default headers values
func NewIsContentDeveloperUsingGETForbidden() *IsContentDeveloperUsingGETForbidden {
	return &IsContentDeveloperUsingGETForbidden{}
}

/*IsContentDeveloperUsingGETForbidden handles this case with default header values.

Forbidden
*/
type IsContentDeveloperUsingGETForbidden struct {
}

func (o *IsContentDeveloperUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/contentdev][%d] isContentDeveloperUsingGETForbidden ", 403)
}

func (o *IsContentDeveloperUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIsContentDeveloperUsingGETNotFound creates a IsContentDeveloperUsingGETNotFound with default headers values
func NewIsContentDeveloperUsingGETNotFound() *IsContentDeveloperUsingGETNotFound {
	return &IsContentDeveloperUsingGETNotFound{}
}

/*IsContentDeveloperUsingGETNotFound handles this case with default header values.

Not Found
*/
type IsContentDeveloperUsingGETNotFound struct {
}

func (o *IsContentDeveloperUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/restricted/contentdev][%d] isContentDeveloperUsingGETNotFound ", 404)
}

func (o *IsContentDeveloperUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}