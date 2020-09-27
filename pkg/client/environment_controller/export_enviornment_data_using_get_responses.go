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

// ExportEnviornmentDataUsingGETReader is a Reader for the ExportEnviornmentDataUsingGET structure.
type ExportEnviornmentDataUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportEnviornmentDataUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportEnviornmentDataUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportEnviornmentDataUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportEnviornmentDataUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportEnviornmentDataUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportEnviornmentDataUsingGETOK creates a ExportEnviornmentDataUsingGETOK with default headers values
func NewExportEnviornmentDataUsingGETOK() *ExportEnviornmentDataUsingGETOK {
	return &ExportEnviornmentDataUsingGETOK{}
}

/*ExportEnviornmentDataUsingGETOK handles this case with default header values.

OK
*/
type ExportEnviornmentDataUsingGETOK struct {
	Payload *models.EnvironmentRequestDTO
}

func (o *ExportEnviornmentDataUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/configuration][%d] exportEnviornmentDataUsingGETOK  %+v", 200, o.Payload)
}

func (o *ExportEnviornmentDataUsingGETOK) GetPayload() *models.EnvironmentRequestDTO {
	return o.Payload
}

func (o *ExportEnviornmentDataUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentRequestDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportEnviornmentDataUsingGETUnauthorized creates a ExportEnviornmentDataUsingGETUnauthorized with default headers values
func NewExportEnviornmentDataUsingGETUnauthorized() *ExportEnviornmentDataUsingGETUnauthorized {
	return &ExportEnviornmentDataUsingGETUnauthorized{}
}

/*ExportEnviornmentDataUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ExportEnviornmentDataUsingGETUnauthorized struct {
}

func (o *ExportEnviornmentDataUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/configuration][%d] exportEnviornmentDataUsingGETUnauthorized ", 401)
}

func (o *ExportEnviornmentDataUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportEnviornmentDataUsingGETForbidden creates a ExportEnviornmentDataUsingGETForbidden with default headers values
func NewExportEnviornmentDataUsingGETForbidden() *ExportEnviornmentDataUsingGETForbidden {
	return &ExportEnviornmentDataUsingGETForbidden{}
}

/*ExportEnviornmentDataUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ExportEnviornmentDataUsingGETForbidden struct {
}

func (o *ExportEnviornmentDataUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/configuration][%d] exportEnviornmentDataUsingGETForbidden ", 403)
}

func (o *ExportEnviornmentDataUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportEnviornmentDataUsingGETNotFound creates a ExportEnviornmentDataUsingGETNotFound with default headers values
func NewExportEnviornmentDataUsingGETNotFound() *ExportEnviornmentDataUsingGETNotFound {
	return &ExportEnviornmentDataUsingGETNotFound{}
}

/*ExportEnviornmentDataUsingGETNotFound handles this case with default header values.

Not Found
*/
type ExportEnviornmentDataUsingGETNotFound struct {
}

func (o *ExportEnviornmentDataUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/{environmentId}/configuration][%d] exportEnviornmentDataUsingGETNotFound ", 404)
}

func (o *ExportEnviornmentDataUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
