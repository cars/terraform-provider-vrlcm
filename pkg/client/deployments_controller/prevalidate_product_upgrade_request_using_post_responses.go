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

// PrevalidateProductUpgradeRequestUsingPOSTReader is a Reader for the PrevalidateProductUpgradeRequestUsingPOST structure.
type PrevalidateProductUpgradeRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PrevalidateProductUpgradeRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPrevalidateProductUpgradeRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewPrevalidateProductUpgradeRequestUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPrevalidateProductUpgradeRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPrevalidateProductUpgradeRequestUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPrevalidateProductUpgradeRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPrevalidateProductUpgradeRequestUsingPOSTOK creates a PrevalidateProductUpgradeRequestUsingPOSTOK with default headers values
func NewPrevalidateProductUpgradeRequestUsingPOSTOK() *PrevalidateProductUpgradeRequestUsingPOSTOK {
	return &PrevalidateProductUpgradeRequestUsingPOSTOK{}
}

/*PrevalidateProductUpgradeRequestUsingPOSTOK handles this case with default header values.

OK
*/
type PrevalidateProductUpgradeRequestUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/upgrade/prevalidate][%d] prevalidateProductUpgradeRequestUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPrevalidateProductUpgradeRequestUsingPOSTCreated creates a PrevalidateProductUpgradeRequestUsingPOSTCreated with default headers values
func NewPrevalidateProductUpgradeRequestUsingPOSTCreated() *PrevalidateProductUpgradeRequestUsingPOSTCreated {
	return &PrevalidateProductUpgradeRequestUsingPOSTCreated{}
}

/*PrevalidateProductUpgradeRequestUsingPOSTCreated handles this case with default header values.

Created
*/
type PrevalidateProductUpgradeRequestUsingPOSTCreated struct {
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/upgrade/prevalidate][%d] prevalidateProductUpgradeRequestUsingPOSTCreated ", 201)
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPrevalidateProductUpgradeRequestUsingPOSTUnauthorized creates a PrevalidateProductUpgradeRequestUsingPOSTUnauthorized with default headers values
func NewPrevalidateProductUpgradeRequestUsingPOSTUnauthorized() *PrevalidateProductUpgradeRequestUsingPOSTUnauthorized {
	return &PrevalidateProductUpgradeRequestUsingPOSTUnauthorized{}
}

/*PrevalidateProductUpgradeRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type PrevalidateProductUpgradeRequestUsingPOSTUnauthorized struct {
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/upgrade/prevalidate][%d] prevalidateProductUpgradeRequestUsingPOSTUnauthorized ", 401)
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPrevalidateProductUpgradeRequestUsingPOSTForbidden creates a PrevalidateProductUpgradeRequestUsingPOSTForbidden with default headers values
func NewPrevalidateProductUpgradeRequestUsingPOSTForbidden() *PrevalidateProductUpgradeRequestUsingPOSTForbidden {
	return &PrevalidateProductUpgradeRequestUsingPOSTForbidden{}
}

/*PrevalidateProductUpgradeRequestUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type PrevalidateProductUpgradeRequestUsingPOSTForbidden struct {
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/upgrade/prevalidate][%d] prevalidateProductUpgradeRequestUsingPOSTForbidden ", 403)
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPrevalidateProductUpgradeRequestUsingPOSTNotFound creates a PrevalidateProductUpgradeRequestUsingPOSTNotFound with default headers values
func NewPrevalidateProductUpgradeRequestUsingPOSTNotFound() *PrevalidateProductUpgradeRequestUsingPOSTNotFound {
	return &PrevalidateProductUpgradeRequestUsingPOSTNotFound{}
}

/*PrevalidateProductUpgradeRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type PrevalidateProductUpgradeRequestUsingPOSTNotFound struct {
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/upgrade/prevalidate][%d] prevalidateProductUpgradeRequestUsingPOSTNotFound ", 404)
}

func (o *PrevalidateProductUpgradeRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}