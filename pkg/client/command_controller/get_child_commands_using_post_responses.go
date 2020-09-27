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

// GetChildCommandsUsingPOSTReader is a Reader for the GetChildCommandsUsingPOST structure.
type GetChildCommandsUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChildCommandsUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetChildCommandsUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewGetChildCommandsUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetChildCommandsUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetChildCommandsUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetChildCommandsUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetChildCommandsUsingPOSTOK creates a GetChildCommandsUsingPOSTOK with default headers values
func NewGetChildCommandsUsingPOSTOK() *GetChildCommandsUsingPOSTOK {
	return &GetChildCommandsUsingPOSTOK{}
}

/*GetChildCommandsUsingPOSTOK handles this case with default header values.

OK
*/
type GetChildCommandsUsingPOSTOK struct {
	Payload []*models.CommandInfoObject
}

func (o *GetChildCommandsUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /command/get/childs][%d] getChildCommandsUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetChildCommandsUsingPOSTOK) GetPayload() []*models.CommandInfoObject {
	return o.Payload
}

func (o *GetChildCommandsUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChildCommandsUsingPOSTCreated creates a GetChildCommandsUsingPOSTCreated with default headers values
func NewGetChildCommandsUsingPOSTCreated() *GetChildCommandsUsingPOSTCreated {
	return &GetChildCommandsUsingPOSTCreated{}
}

/*GetChildCommandsUsingPOSTCreated handles this case with default header values.

Created
*/
type GetChildCommandsUsingPOSTCreated struct {
}

func (o *GetChildCommandsUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /command/get/childs][%d] getChildCommandsUsingPOSTCreated ", 201)
}

func (o *GetChildCommandsUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildCommandsUsingPOSTUnauthorized creates a GetChildCommandsUsingPOSTUnauthorized with default headers values
func NewGetChildCommandsUsingPOSTUnauthorized() *GetChildCommandsUsingPOSTUnauthorized {
	return &GetChildCommandsUsingPOSTUnauthorized{}
}

/*GetChildCommandsUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type GetChildCommandsUsingPOSTUnauthorized struct {
}

func (o *GetChildCommandsUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /command/get/childs][%d] getChildCommandsUsingPOSTUnauthorized ", 401)
}

func (o *GetChildCommandsUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildCommandsUsingPOSTForbidden creates a GetChildCommandsUsingPOSTForbidden with default headers values
func NewGetChildCommandsUsingPOSTForbidden() *GetChildCommandsUsingPOSTForbidden {
	return &GetChildCommandsUsingPOSTForbidden{}
}

/*GetChildCommandsUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type GetChildCommandsUsingPOSTForbidden struct {
}

func (o *GetChildCommandsUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /command/get/childs][%d] getChildCommandsUsingPOSTForbidden ", 403)
}

func (o *GetChildCommandsUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildCommandsUsingPOSTNotFound creates a GetChildCommandsUsingPOSTNotFound with default headers values
func NewGetChildCommandsUsingPOSTNotFound() *GetChildCommandsUsingPOSTNotFound {
	return &GetChildCommandsUsingPOSTNotFound{}
}

/*GetChildCommandsUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type GetChildCommandsUsingPOSTNotFound struct {
}

func (o *GetChildCommandsUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /command/get/childs][%d] getChildCommandsUsingPOSTNotFound ", 404)
}

func (o *GetChildCommandsUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}