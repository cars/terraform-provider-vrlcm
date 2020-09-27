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

// ValidateEnvironmentUsingPOST1Reader is a Reader for the ValidateEnvironmentUsingPOST1 structure.
type ValidateEnvironmentUsingPOST1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateEnvironmentUsingPOST1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidateEnvironmentUsingPOST1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewValidateEnvironmentUsingPOST1Created()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewValidateEnvironmentUsingPOST1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewValidateEnvironmentUsingPOST1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewValidateEnvironmentUsingPOST1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewValidateEnvironmentUsingPOST1OK creates a ValidateEnvironmentUsingPOST1OK with default headers values
func NewValidateEnvironmentUsingPOST1OK() *ValidateEnvironmentUsingPOST1OK {
	return &ValidateEnvironmentUsingPOST1OK{}
}

/*ValidateEnvironmentUsingPOST1OK handles this case with default header values.

OK
*/
type ValidateEnvironmentUsingPOST1OK struct {
	Payload *models.GenericRequestResponse
}

func (o *ValidateEnvironmentUsingPOST1OK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/prevalidate][%d] validateEnvironmentUsingPOST1OK  %+v", 200, o.Payload)
}

func (o *ValidateEnvironmentUsingPOST1OK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *ValidateEnvironmentUsingPOST1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidateEnvironmentUsingPOST1Created creates a ValidateEnvironmentUsingPOST1Created with default headers values
func NewValidateEnvironmentUsingPOST1Created() *ValidateEnvironmentUsingPOST1Created {
	return &ValidateEnvironmentUsingPOST1Created{}
}

/*ValidateEnvironmentUsingPOST1Created handles this case with default header values.

Created
*/
type ValidateEnvironmentUsingPOST1Created struct {
}

func (o *ValidateEnvironmentUsingPOST1Created) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/prevalidate][%d] validateEnvironmentUsingPOST1Created ", 201)
}

func (o *ValidateEnvironmentUsingPOST1Created) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateEnvironmentUsingPOST1Unauthorized creates a ValidateEnvironmentUsingPOST1Unauthorized with default headers values
func NewValidateEnvironmentUsingPOST1Unauthorized() *ValidateEnvironmentUsingPOST1Unauthorized {
	return &ValidateEnvironmentUsingPOST1Unauthorized{}
}

/*ValidateEnvironmentUsingPOST1Unauthorized handles this case with default header values.

Unauthorized
*/
type ValidateEnvironmentUsingPOST1Unauthorized struct {
}

func (o *ValidateEnvironmentUsingPOST1Unauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/prevalidate][%d] validateEnvironmentUsingPOST1Unauthorized ", 401)
}

func (o *ValidateEnvironmentUsingPOST1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateEnvironmentUsingPOST1Forbidden creates a ValidateEnvironmentUsingPOST1Forbidden with default headers values
func NewValidateEnvironmentUsingPOST1Forbidden() *ValidateEnvironmentUsingPOST1Forbidden {
	return &ValidateEnvironmentUsingPOST1Forbidden{}
}

/*ValidateEnvironmentUsingPOST1Forbidden handles this case with default header values.

Forbidden
*/
type ValidateEnvironmentUsingPOST1Forbidden struct {
}

func (o *ValidateEnvironmentUsingPOST1Forbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/prevalidate][%d] validateEnvironmentUsingPOST1Forbidden ", 403)
}

func (o *ValidateEnvironmentUsingPOST1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewValidateEnvironmentUsingPOST1NotFound creates a ValidateEnvironmentUsingPOST1NotFound with default headers values
func NewValidateEnvironmentUsingPOST1NotFound() *ValidateEnvironmentUsingPOST1NotFound {
	return &ValidateEnvironmentUsingPOST1NotFound{}
}

/*ValidateEnvironmentUsingPOST1NotFound handles this case with default header values.

Not Found
*/
type ValidateEnvironmentUsingPOST1NotFound struct {
}

func (o *ValidateEnvironmentUsingPOST1NotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/prevalidate][%d] validateEnvironmentUsingPOST1NotFound ", 404)
}

func (o *ValidateEnvironmentUsingPOST1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
