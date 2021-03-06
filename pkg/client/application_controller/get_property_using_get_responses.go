// Code generated by go-swagger; DO NOT EDIT.

package application_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetPropertyUsingGETReader is a Reader for the GetPropertyUsingGET structure.
type GetPropertyUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPropertyUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPropertyUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetPropertyUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPropertyUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPropertyUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPropertyUsingGETOK creates a GetPropertyUsingGETOK with default headers values
func NewGetPropertyUsingGETOK() *GetPropertyUsingGETOK {
	return &GetPropertyUsingGETOK{}
}

/*GetPropertyUsingGETOK handles this case with default header values.

OK
*/
type GetPropertyUsingGETOK struct {
	Payload *models.ApplicationDTO
}

func (o *GetPropertyUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/shell/api/apps/{vmid}][%d] getPropertyUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetPropertyUsingGETOK) GetPayload() *models.ApplicationDTO {
	return o.Payload
}

func (o *GetPropertyUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPropertyUsingGETUnauthorized creates a GetPropertyUsingGETUnauthorized with default headers values
func NewGetPropertyUsingGETUnauthorized() *GetPropertyUsingGETUnauthorized {
	return &GetPropertyUsingGETUnauthorized{}
}

/*GetPropertyUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPropertyUsingGETUnauthorized struct {
}

func (o *GetPropertyUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/shell/api/apps/{vmid}][%d] getPropertyUsingGETUnauthorized ", 401)
}

func (o *GetPropertyUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyUsingGETForbidden creates a GetPropertyUsingGETForbidden with default headers values
func NewGetPropertyUsingGETForbidden() *GetPropertyUsingGETForbidden {
	return &GetPropertyUsingGETForbidden{}
}

/*GetPropertyUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetPropertyUsingGETForbidden struct {
}

func (o *GetPropertyUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/shell/api/apps/{vmid}][%d] getPropertyUsingGETForbidden ", 403)
}

func (o *GetPropertyUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPropertyUsingGETNotFound creates a GetPropertyUsingGETNotFound with default headers values
func NewGetPropertyUsingGETNotFound() *GetPropertyUsingGETNotFound {
	return &GetPropertyUsingGETNotFound{}
}

/*GetPropertyUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetPropertyUsingGETNotFound struct {
}

func (o *GetPropertyUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/shell/api/apps/{vmid}][%d] getPropertyUsingGETNotFound ", 404)
}

func (o *GetPropertyUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
