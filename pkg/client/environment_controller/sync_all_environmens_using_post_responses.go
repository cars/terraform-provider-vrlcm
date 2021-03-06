// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SyncAllEnvironmensUsingPOSTReader is a Reader for the SyncAllEnvironmensUsingPOST structure.
type SyncAllEnvironmensUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SyncAllEnvironmensUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSyncAllEnvironmensUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSyncAllEnvironmensUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSyncAllEnvironmensUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSyncAllEnvironmensUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSyncAllEnvironmensUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSyncAllEnvironmensUsingPOSTOK creates a SyncAllEnvironmensUsingPOSTOK with default headers values
func NewSyncAllEnvironmensUsingPOSTOK() *SyncAllEnvironmensUsingPOSTOK {
	return &SyncAllEnvironmensUsingPOSTOK{}
}

/*SyncAllEnvironmensUsingPOSTOK handles this case with default header values.

OK
*/
type SyncAllEnvironmensUsingPOSTOK struct {
	Payload interface{}
}

func (o *SyncAllEnvironmensUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/inventorysync][%d] syncAllEnvironmensUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *SyncAllEnvironmensUsingPOSTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *SyncAllEnvironmensUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSyncAllEnvironmensUsingPOSTCreated creates a SyncAllEnvironmensUsingPOSTCreated with default headers values
func NewSyncAllEnvironmensUsingPOSTCreated() *SyncAllEnvironmensUsingPOSTCreated {
	return &SyncAllEnvironmensUsingPOSTCreated{}
}

/*SyncAllEnvironmensUsingPOSTCreated handles this case with default header values.

Created
*/
type SyncAllEnvironmensUsingPOSTCreated struct {
}

func (o *SyncAllEnvironmensUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/inventorysync][%d] syncAllEnvironmensUsingPOSTCreated ", 201)
}

func (o *SyncAllEnvironmensUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncAllEnvironmensUsingPOSTUnauthorized creates a SyncAllEnvironmensUsingPOSTUnauthorized with default headers values
func NewSyncAllEnvironmensUsingPOSTUnauthorized() *SyncAllEnvironmensUsingPOSTUnauthorized {
	return &SyncAllEnvironmensUsingPOSTUnauthorized{}
}

/*SyncAllEnvironmensUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type SyncAllEnvironmensUsingPOSTUnauthorized struct {
}

func (o *SyncAllEnvironmensUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/inventorysync][%d] syncAllEnvironmensUsingPOSTUnauthorized ", 401)
}

func (o *SyncAllEnvironmensUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncAllEnvironmensUsingPOSTForbidden creates a SyncAllEnvironmensUsingPOSTForbidden with default headers values
func NewSyncAllEnvironmensUsingPOSTForbidden() *SyncAllEnvironmensUsingPOSTForbidden {
	return &SyncAllEnvironmensUsingPOSTForbidden{}
}

/*SyncAllEnvironmensUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type SyncAllEnvironmensUsingPOSTForbidden struct {
}

func (o *SyncAllEnvironmensUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/inventorysync][%d] syncAllEnvironmensUsingPOSTForbidden ", 403)
}

func (o *SyncAllEnvironmensUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSyncAllEnvironmensUsingPOSTNotFound creates a SyncAllEnvironmensUsingPOSTNotFound with default headers values
func NewSyncAllEnvironmensUsingPOSTNotFound() *SyncAllEnvironmensUsingPOSTNotFound {
	return &SyncAllEnvironmensUsingPOSTNotFound{}
}

/*SyncAllEnvironmensUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type SyncAllEnvironmensUsingPOSTNotFound struct {
}

func (o *SyncAllEnvironmensUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/inventorysync][%d] syncAllEnvironmensUsingPOSTNotFound ", 404)
}

func (o *SyncAllEnvironmensUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
