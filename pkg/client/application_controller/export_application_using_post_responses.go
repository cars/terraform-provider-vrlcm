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

// ExportApplicationUsingPOSTReader is a Reader for the ExportApplicationUsingPOST structure.
type ExportApplicationUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportApplicationUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExportApplicationUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewExportApplicationUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExportApplicationUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExportApplicationUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExportApplicationUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExportApplicationUsingPOSTOK creates a ExportApplicationUsingPOSTOK with default headers values
func NewExportApplicationUsingPOSTOK() *ExportApplicationUsingPOSTOK {
	return &ExportApplicationUsingPOSTOK{}
}

/*ExportApplicationUsingPOSTOK handles this case with default header values.

OK
*/
type ExportApplicationUsingPOSTOK struct {
	Payload []*models.ApplicationDTO
}

func (o *ExportApplicationUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/shell/api/apps/export][%d] exportApplicationUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *ExportApplicationUsingPOSTOK) GetPayload() []*models.ApplicationDTO {
	return o.Payload
}

func (o *ExportApplicationUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExportApplicationUsingPOSTCreated creates a ExportApplicationUsingPOSTCreated with default headers values
func NewExportApplicationUsingPOSTCreated() *ExportApplicationUsingPOSTCreated {
	return &ExportApplicationUsingPOSTCreated{}
}

/*ExportApplicationUsingPOSTCreated handles this case with default header values.

Created
*/
type ExportApplicationUsingPOSTCreated struct {
}

func (o *ExportApplicationUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/shell/api/apps/export][%d] exportApplicationUsingPOSTCreated ", 201)
}

func (o *ExportApplicationUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportApplicationUsingPOSTUnauthorized creates a ExportApplicationUsingPOSTUnauthorized with default headers values
func NewExportApplicationUsingPOSTUnauthorized() *ExportApplicationUsingPOSTUnauthorized {
	return &ExportApplicationUsingPOSTUnauthorized{}
}

/*ExportApplicationUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ExportApplicationUsingPOSTUnauthorized struct {
}

func (o *ExportApplicationUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/shell/api/apps/export][%d] exportApplicationUsingPOSTUnauthorized ", 401)
}

func (o *ExportApplicationUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportApplicationUsingPOSTForbidden creates a ExportApplicationUsingPOSTForbidden with default headers values
func NewExportApplicationUsingPOSTForbidden() *ExportApplicationUsingPOSTForbidden {
	return &ExportApplicationUsingPOSTForbidden{}
}

/*ExportApplicationUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ExportApplicationUsingPOSTForbidden struct {
}

func (o *ExportApplicationUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/shell/api/apps/export][%d] exportApplicationUsingPOSTForbidden ", 403)
}

func (o *ExportApplicationUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportApplicationUsingPOSTNotFound creates a ExportApplicationUsingPOSTNotFound with default headers values
func NewExportApplicationUsingPOSTNotFound() *ExportApplicationUsingPOSTNotFound {
	return &ExportApplicationUsingPOSTNotFound{}
}

/*ExportApplicationUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ExportApplicationUsingPOSTNotFound struct {
}

func (o *ExportApplicationUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/shell/api/apps/export][%d] exportApplicationUsingPOSTNotFound ", 404)
}

func (o *ExportApplicationUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
