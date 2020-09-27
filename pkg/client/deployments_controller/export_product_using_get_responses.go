// Code generated by go-swagger; DO NOT EDIT.

package deployments_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// ExportProductUsingGETReader is a Reader for the ExportProductUsingGET structure.
type ExportProductUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportProductUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportProductUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportProductUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportProductUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportProductUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportProductUsingGETOK creates a ExportProductUsingGETOK with default headers values
func NewExportProductUsingGETOK() *ExportProductUsingGETOK {
	return &ExportProductUsingGETOK{}
}

/*ExportProductUsingGETOK handles this case with default header values.

OK
*/
type ExportProductUsingGETOK struct {
	Payload *models.EnvironmentRequestDTO
}

func (o *ExportProductUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/products/{productId}/configuration][%d] exportProductUsingGETOK  %+v", 200, o.Payload)
}

func (o *ExportProductUsingGETOK) GetPayload() *models.EnvironmentRequestDTO {
	return o.Payload
}

func (o *ExportProductUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentRequestDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportProductUsingGETUnauthorized creates a ExportProductUsingGETUnauthorized with default headers values
func NewExportProductUsingGETUnauthorized() *ExportProductUsingGETUnauthorized {
	return &ExportProductUsingGETUnauthorized{}
}

/*ExportProductUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ExportProductUsingGETUnauthorized struct {
}

func (o *ExportProductUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/products/{productId}/configuration][%d] exportProductUsingGETUnauthorized ", 401)
}

func (o *ExportProductUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportProductUsingGETForbidden creates a ExportProductUsingGETForbidden with default headers values
func NewExportProductUsingGETForbidden() *ExportProductUsingGETForbidden {
	return &ExportProductUsingGETForbidden{}
}

/*ExportProductUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ExportProductUsingGETForbidden struct {
}

func (o *ExportProductUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/products/{productId}/configuration][%d] exportProductUsingGETForbidden ", 403)
}

func (o *ExportProductUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportProductUsingGETNotFound creates a ExportProductUsingGETNotFound with default headers values
func NewExportProductUsingGETNotFound() *ExportProductUsingGETNotFound {
	return &ExportProductUsingGETNotFound{}
}

/*ExportProductUsingGETNotFound handles this case with default header values.

Not Found
*/
type ExportProductUsingGETNotFound struct {
}

func (o *ExportProductUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/products/{productId}/configuration][%d] exportProductUsingGETNotFound ", 404)
}

func (o *ExportProductUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
