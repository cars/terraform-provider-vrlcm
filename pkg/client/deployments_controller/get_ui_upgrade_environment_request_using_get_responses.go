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

// GetUIUpgradeEnvironmentRequestUsingGETReader is a Reader for the GetUIUpgradeEnvironmentRequestUsingGET structure.
type GetUIUpgradeEnvironmentRequestUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUIUpgradeEnvironmentRequestUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUIUpgradeEnvironmentRequestUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUIUpgradeEnvironmentRequestUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUIUpgradeEnvironmentRequestUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUIUpgradeEnvironmentRequestUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUIUpgradeEnvironmentRequestUsingGETOK creates a GetUIUpgradeEnvironmentRequestUsingGETOK with default headers values
func NewGetUIUpgradeEnvironmentRequestUsingGETOK() *GetUIUpgradeEnvironmentRequestUsingGETOK {
	return &GetUIUpgradeEnvironmentRequestUsingGETOK{}
}

/*GetUIUpgradeEnvironmentRequestUsingGETOK handles this case with default header values.

OK
*/
type GetUIUpgradeEnvironmentRequestUsingGETOK struct {
	Payload *models.UpgradeEnvironmentUIRequest
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/upgrade/{requestId}/ui][%d] getUiUpgradeEnvironmentRequestUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETOK) GetPayload() *models.UpgradeEnvironmentUIRequest {
	return o.Payload
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpgradeEnvironmentUIRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUIUpgradeEnvironmentRequestUsingGETUnauthorized creates a GetUIUpgradeEnvironmentRequestUsingGETUnauthorized with default headers values
func NewGetUIUpgradeEnvironmentRequestUsingGETUnauthorized() *GetUIUpgradeEnvironmentRequestUsingGETUnauthorized {
	return &GetUIUpgradeEnvironmentRequestUsingGETUnauthorized{}
}

/*GetUIUpgradeEnvironmentRequestUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUIUpgradeEnvironmentRequestUsingGETUnauthorized struct {
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/upgrade/{requestId}/ui][%d] getUiUpgradeEnvironmentRequestUsingGETUnauthorized ", 401)
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUIUpgradeEnvironmentRequestUsingGETForbidden creates a GetUIUpgradeEnvironmentRequestUsingGETForbidden with default headers values
func NewGetUIUpgradeEnvironmentRequestUsingGETForbidden() *GetUIUpgradeEnvironmentRequestUsingGETForbidden {
	return &GetUIUpgradeEnvironmentRequestUsingGETForbidden{}
}

/*GetUIUpgradeEnvironmentRequestUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetUIUpgradeEnvironmentRequestUsingGETForbidden struct {
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/upgrade/{requestId}/ui][%d] getUiUpgradeEnvironmentRequestUsingGETForbidden ", 403)
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUIUpgradeEnvironmentRequestUsingGETNotFound creates a GetUIUpgradeEnvironmentRequestUsingGETNotFound with default headers values
func NewGetUIUpgradeEnvironmentRequestUsingGETNotFound() *GetUIUpgradeEnvironmentRequestUsingGETNotFound {
	return &GetUIUpgradeEnvironmentRequestUsingGETNotFound{}
}

/*GetUIUpgradeEnvironmentRequestUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetUIUpgradeEnvironmentRequestUsingGETNotFound struct {
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/upgrade/{requestId}/ui][%d] getUiUpgradeEnvironmentRequestUsingGETNotFound ", 404)
}

func (o *GetUIUpgradeEnvironmentRequestUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
