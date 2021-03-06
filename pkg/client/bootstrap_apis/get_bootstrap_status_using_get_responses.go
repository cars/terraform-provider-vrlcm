// Code generated by go-swagger; DO NOT EDIT.

package bootstrap_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetBootstrapStatusUsingGETReader is a Reader for the GetBootstrapStatusUsingGET structure.
type GetBootstrapStatusUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBootstrapStatusUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBootstrapStatusUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBootstrapStatusUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBootstrapStatusUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBootstrapStatusUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBootstrapStatusUsingGETOK creates a GetBootstrapStatusUsingGETOK with default headers values
func NewGetBootstrapStatusUsingGETOK() *GetBootstrapStatusUsingGETOK {
	return &GetBootstrapStatusUsingGETOK{}
}

/*GetBootstrapStatusUsingGETOK handles this case with default header values.

OK
*/
type GetBootstrapStatusUsingGETOK struct {
	Payload *models.BootstrapResponseDTO
}

func (o *GetBootstrapStatusUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/status][%d] getBootstrapStatusUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetBootstrapStatusUsingGETOK) GetPayload() *models.BootstrapResponseDTO {
	return o.Payload
}

func (o *GetBootstrapStatusUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BootstrapResponseDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBootstrapStatusUsingGETUnauthorized creates a GetBootstrapStatusUsingGETUnauthorized with default headers values
func NewGetBootstrapStatusUsingGETUnauthorized() *GetBootstrapStatusUsingGETUnauthorized {
	return &GetBootstrapStatusUsingGETUnauthorized{}
}

/*GetBootstrapStatusUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBootstrapStatusUsingGETUnauthorized struct {
}

func (o *GetBootstrapStatusUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/status][%d] getBootstrapStatusUsingGETUnauthorized ", 401)
}

func (o *GetBootstrapStatusUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBootstrapStatusUsingGETForbidden creates a GetBootstrapStatusUsingGETForbidden with default headers values
func NewGetBootstrapStatusUsingGETForbidden() *GetBootstrapStatusUsingGETForbidden {
	return &GetBootstrapStatusUsingGETForbidden{}
}

/*GetBootstrapStatusUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetBootstrapStatusUsingGETForbidden struct {
}

func (o *GetBootstrapStatusUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/status][%d] getBootstrapStatusUsingGETForbidden ", 403)
}

func (o *GetBootstrapStatusUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBootstrapStatusUsingGETNotFound creates a GetBootstrapStatusUsingGETNotFound with default headers values
func NewGetBootstrapStatusUsingGETNotFound() *GetBootstrapStatusUsingGETNotFound {
	return &GetBootstrapStatusUsingGETNotFound{}
}

/*GetBootstrapStatusUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetBootstrapStatusUsingGETNotFound struct {
}

func (o *GetBootstrapStatusUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/status][%d] getBootstrapStatusUsingGETNotFound ", 404)
}

func (o *GetBootstrapStatusUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
