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

// CrossProductMonitervRAWithvROPSUsingPOSTReader is a Reader for the CrossProductMonitervRAWithvROPSUsingPOST structure.
type CrossProductMonitervRAWithvROPSUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CrossProductMonitervRAWithvROPSUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCrossProductMonitervRAWithvROPSUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCrossProductMonitervRAWithvROPSUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCrossProductMonitervRAWithvROPSUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCrossProductMonitervRAWithvROPSUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCrossProductMonitervRAWithvROPSUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCrossProductMonitervRAWithvROPSUsingPOSTOK creates a CrossProductMonitervRAWithvROPSUsingPOSTOK with default headers values
func NewCrossProductMonitervRAWithvROPSUsingPOSTOK() *CrossProductMonitervRAWithvROPSUsingPOSTOK {
	return &CrossProductMonitervRAWithvROPSUsingPOSTOK{}
}

/*CrossProductMonitervRAWithvROPSUsingPOSTOK handles this case with default header values.

OK
*/
type CrossProductMonitervRAWithvROPSUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/monitorvRAwithvROPS][%d] crossProductMonitervRAWithvROPSUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCrossProductMonitervRAWithvROPSUsingPOSTCreated creates a CrossProductMonitervRAWithvROPSUsingPOSTCreated with default headers values
func NewCrossProductMonitervRAWithvROPSUsingPOSTCreated() *CrossProductMonitervRAWithvROPSUsingPOSTCreated {
	return &CrossProductMonitervRAWithvROPSUsingPOSTCreated{}
}

/*CrossProductMonitervRAWithvROPSUsingPOSTCreated handles this case with default header values.

Created
*/
type CrossProductMonitervRAWithvROPSUsingPOSTCreated struct {
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/monitorvRAwithvROPS][%d] crossProductMonitervRAWithvROPSUsingPOSTCreated ", 201)
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCrossProductMonitervRAWithvROPSUsingPOSTUnauthorized creates a CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized with default headers values
func NewCrossProductMonitervRAWithvROPSUsingPOSTUnauthorized() *CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized {
	return &CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized{}
}

/*CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized struct {
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/monitorvRAwithvROPS][%d] crossProductMonitervRAWithvROPSUsingPOSTUnauthorized ", 401)
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCrossProductMonitervRAWithvROPSUsingPOSTForbidden creates a CrossProductMonitervRAWithvROPSUsingPOSTForbidden with default headers values
func NewCrossProductMonitervRAWithvROPSUsingPOSTForbidden() *CrossProductMonitervRAWithvROPSUsingPOSTForbidden {
	return &CrossProductMonitervRAWithvROPSUsingPOSTForbidden{}
}

/*CrossProductMonitervRAWithvROPSUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CrossProductMonitervRAWithvROPSUsingPOSTForbidden struct {
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/monitorvRAwithvROPS][%d] crossProductMonitervRAWithvROPSUsingPOSTForbidden ", 403)
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCrossProductMonitervRAWithvROPSUsingPOSTNotFound creates a CrossProductMonitervRAWithvROPSUsingPOSTNotFound with default headers values
func NewCrossProductMonitervRAWithvROPSUsingPOSTNotFound() *CrossProductMonitervRAWithvROPSUsingPOSTNotFound {
	return &CrossProductMonitervRAWithvROPSUsingPOSTNotFound{}
}

/*CrossProductMonitervRAWithvROPSUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CrossProductMonitervRAWithvROPSUsingPOSTNotFound struct {
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/monitorvRAwithvROPS][%d] crossProductMonitervRAWithvROPSUsingPOSTNotFound ", 404)
}

func (o *CrossProductMonitervRAWithvROPSUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
