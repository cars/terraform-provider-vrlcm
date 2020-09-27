// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetEnvironmenstByStatusUsingGETReader is a Reader for the GetEnvironmenstByStatusUsingGET structure.
type GetEnvironmenstByStatusUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEnvironmenstByStatusUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEnvironmenstByStatusUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEnvironmenstByStatusUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetEnvironmenstByStatusUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEnvironmenstByStatusUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEnvironmenstByStatusUsingGETOK creates a GetEnvironmenstByStatusUsingGETOK with default headers values
func NewGetEnvironmenstByStatusUsingGETOK() *GetEnvironmenstByStatusUsingGETOK {
	return &GetEnvironmenstByStatusUsingGETOK{}
}

/*GetEnvironmenstByStatusUsingGETOK handles this case with default header values.

OK
*/
type GetEnvironmenstByStatusUsingGETOK struct {
	Payload []*models.EnvironmentRequestDTO
}

func (o *GetEnvironmenstByStatusUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/status/{status}][%d] getEnvironmenstByStatusUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEnvironmenstByStatusUsingGETOK) GetPayload() []*models.EnvironmentRequestDTO {
	return o.Payload
}

func (o *GetEnvironmenstByStatusUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEnvironmenstByStatusUsingGETUnauthorized creates a GetEnvironmenstByStatusUsingGETUnauthorized with default headers values
func NewGetEnvironmenstByStatusUsingGETUnauthorized() *GetEnvironmenstByStatusUsingGETUnauthorized {
	return &GetEnvironmenstByStatusUsingGETUnauthorized{}
}

/*GetEnvironmenstByStatusUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEnvironmenstByStatusUsingGETUnauthorized struct {
}

func (o *GetEnvironmenstByStatusUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/status/{status}][%d] getEnvironmenstByStatusUsingGETUnauthorized ", 401)
}

func (o *GetEnvironmenstByStatusUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmenstByStatusUsingGETForbidden creates a GetEnvironmenstByStatusUsingGETForbidden with default headers values
func NewGetEnvironmenstByStatusUsingGETForbidden() *GetEnvironmenstByStatusUsingGETForbidden {
	return &GetEnvironmenstByStatusUsingGETForbidden{}
}

/*GetEnvironmenstByStatusUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetEnvironmenstByStatusUsingGETForbidden struct {
}

func (o *GetEnvironmenstByStatusUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/status/{status}][%d] getEnvironmenstByStatusUsingGETForbidden ", 403)
}

func (o *GetEnvironmenstByStatusUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEnvironmenstByStatusUsingGETNotFound creates a GetEnvironmenstByStatusUsingGETNotFound with default headers values
func NewGetEnvironmenstByStatusUsingGETNotFound() *GetEnvironmenstByStatusUsingGETNotFound {
	return &GetEnvironmenstByStatusUsingGETNotFound{}
}

/*GetEnvironmenstByStatusUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetEnvironmenstByStatusUsingGETNotFound struct {
}

func (o *GetEnvironmenstByStatusUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/status/{status}][%d] getEnvironmenstByStatusUsingGETNotFound ", 404)
}

func (o *GetEnvironmenstByStatusUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
