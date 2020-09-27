// Code generated by go-swagger; DO NOT EDIT.

package command_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetCommandsForEnvironmentUsingPOSTReader is a Reader for the GetCommandsForEnvironmentUsingPOST structure.
type GetCommandsForEnvironmentUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCommandsForEnvironmentUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCommandsForEnvironmentUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewGetCommandsForEnvironmentUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCommandsForEnvironmentUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCommandsForEnvironmentUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCommandsForEnvironmentUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCommandsForEnvironmentUsingPOSTOK creates a GetCommandsForEnvironmentUsingPOSTOK with default headers values
func NewGetCommandsForEnvironmentUsingPOSTOK() *GetCommandsForEnvironmentUsingPOSTOK {
	return &GetCommandsForEnvironmentUsingPOSTOK{}
}

/*GetCommandsForEnvironmentUsingPOSTOK handles this case with default header values.

OK
*/
type GetCommandsForEnvironmentUsingPOSTOK struct {
	Payload []*models.CommandInfoObject
}

func (o *GetCommandsForEnvironmentUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /command/get/by-environment][%d] getCommandsForEnvironmentUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetCommandsForEnvironmentUsingPOSTOK) GetPayload() []*models.CommandInfoObject {
	return o.Payload
}

func (o *GetCommandsForEnvironmentUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCommandsForEnvironmentUsingPOSTCreated creates a GetCommandsForEnvironmentUsingPOSTCreated with default headers values
func NewGetCommandsForEnvironmentUsingPOSTCreated() *GetCommandsForEnvironmentUsingPOSTCreated {
	return &GetCommandsForEnvironmentUsingPOSTCreated{}
}

/*GetCommandsForEnvironmentUsingPOSTCreated handles this case with default header values.

Created
*/
type GetCommandsForEnvironmentUsingPOSTCreated struct {
}

func (o *GetCommandsForEnvironmentUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /command/get/by-environment][%d] getCommandsForEnvironmentUsingPOSTCreated ", 201)
}

func (o *GetCommandsForEnvironmentUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCommandsForEnvironmentUsingPOSTUnauthorized creates a GetCommandsForEnvironmentUsingPOSTUnauthorized with default headers values
func NewGetCommandsForEnvironmentUsingPOSTUnauthorized() *GetCommandsForEnvironmentUsingPOSTUnauthorized {
	return &GetCommandsForEnvironmentUsingPOSTUnauthorized{}
}

/*GetCommandsForEnvironmentUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCommandsForEnvironmentUsingPOSTUnauthorized struct {
}

func (o *GetCommandsForEnvironmentUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /command/get/by-environment][%d] getCommandsForEnvironmentUsingPOSTUnauthorized ", 401)
}

func (o *GetCommandsForEnvironmentUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCommandsForEnvironmentUsingPOSTForbidden creates a GetCommandsForEnvironmentUsingPOSTForbidden with default headers values
func NewGetCommandsForEnvironmentUsingPOSTForbidden() *GetCommandsForEnvironmentUsingPOSTForbidden {
	return &GetCommandsForEnvironmentUsingPOSTForbidden{}
}

/*GetCommandsForEnvironmentUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type GetCommandsForEnvironmentUsingPOSTForbidden struct {
}

func (o *GetCommandsForEnvironmentUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /command/get/by-environment][%d] getCommandsForEnvironmentUsingPOSTForbidden ", 403)
}

func (o *GetCommandsForEnvironmentUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCommandsForEnvironmentUsingPOSTNotFound creates a GetCommandsForEnvironmentUsingPOSTNotFound with default headers values
func NewGetCommandsForEnvironmentUsingPOSTNotFound() *GetCommandsForEnvironmentUsingPOSTNotFound {
	return &GetCommandsForEnvironmentUsingPOSTNotFound{}
}

/*GetCommandsForEnvironmentUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type GetCommandsForEnvironmentUsingPOSTNotFound struct {
}

func (o *GetCommandsForEnvironmentUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /command/get/by-environment][%d] getCommandsForEnvironmentUsingPOSTNotFound ", 404)
}

func (o *GetCommandsForEnvironmentUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
