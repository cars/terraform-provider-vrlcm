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

// GetUIEnvironmentRequestUsingGETReader is a Reader for the GetUIEnvironmentRequestUsingGET structure.
type GetUIEnvironmentRequestUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUIEnvironmentRequestUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUIEnvironmentRequestUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUIEnvironmentRequestUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUIEnvironmentRequestUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUIEnvironmentRequestUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUIEnvironmentRequestUsingGETOK creates a GetUIEnvironmentRequestUsingGETOK with default headers values
func NewGetUIEnvironmentRequestUsingGETOK() *GetUIEnvironmentRequestUsingGETOK {
	return &GetUIEnvironmentRequestUsingGETOK{}
}

/*GetUIEnvironmentRequestUsingGETOK handles this case with default header values.

OK
*/
type GetUIEnvironmentRequestUsingGETOK struct {
	Payload *models.CreateEnvironmentUIRequest
}

func (o *GetUIEnvironmentRequestUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/{requestId}/ui][%d] getUiEnvironmentRequestUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetUIEnvironmentRequestUsingGETOK) GetPayload() *models.CreateEnvironmentUIRequest {
	return o.Payload
}

func (o *GetUIEnvironmentRequestUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateEnvironmentUIRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUIEnvironmentRequestUsingGETUnauthorized creates a GetUIEnvironmentRequestUsingGETUnauthorized with default headers values
func NewGetUIEnvironmentRequestUsingGETUnauthorized() *GetUIEnvironmentRequestUsingGETUnauthorized {
	return &GetUIEnvironmentRequestUsingGETUnauthorized{}
}

/*GetUIEnvironmentRequestUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUIEnvironmentRequestUsingGETUnauthorized struct {
}

func (o *GetUIEnvironmentRequestUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/{requestId}/ui][%d] getUiEnvironmentRequestUsingGETUnauthorized ", 401)
}

func (o *GetUIEnvironmentRequestUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUIEnvironmentRequestUsingGETForbidden creates a GetUIEnvironmentRequestUsingGETForbidden with default headers values
func NewGetUIEnvironmentRequestUsingGETForbidden() *GetUIEnvironmentRequestUsingGETForbidden {
	return &GetUIEnvironmentRequestUsingGETForbidden{}
}

/*GetUIEnvironmentRequestUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetUIEnvironmentRequestUsingGETForbidden struct {
}

func (o *GetUIEnvironmentRequestUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/{requestId}/ui][%d] getUiEnvironmentRequestUsingGETForbidden ", 403)
}

func (o *GetUIEnvironmentRequestUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUIEnvironmentRequestUsingGETNotFound creates a GetUIEnvironmentRequestUsingGETNotFound with default headers values
func NewGetUIEnvironmentRequestUsingGETNotFound() *GetUIEnvironmentRequestUsingGETNotFound {
	return &GetUIEnvironmentRequestUsingGETNotFound{}
}

/*GetUIEnvironmentRequestUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetUIEnvironmentRequestUsingGETNotFound struct {
}

func (o *GetUIEnvironmentRequestUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/{requestId}/ui][%d] getUiEnvironmentRequestUsingGETNotFound ", 404)
}

func (o *GetUIEnvironmentRequestUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}