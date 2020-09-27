// Code generated by go-swagger; DO NOT EDIT.

package engine_message_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportMessagesUsingPOSTReader is a Reader for the ImportMessagesUsingPOST structure.
type ImportMessagesUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportMessagesUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportMessagesUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewImportMessagesUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImportMessagesUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportMessagesUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportMessagesUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportMessagesUsingPOSTOK creates a ImportMessagesUsingPOSTOK with default headers values
func NewImportMessagesUsingPOSTOK() *ImportMessagesUsingPOSTOK {
	return &ImportMessagesUsingPOSTOK{}
}

/*ImportMessagesUsingPOSTOK handles this case with default header values.

OK
*/
type ImportMessagesUsingPOSTOK struct {
}

func (o *ImportMessagesUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/automata/api/engine/message/import][%d] importMessagesUsingPOSTOK ", 200)
}

func (o *ImportMessagesUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportMessagesUsingPOSTCreated creates a ImportMessagesUsingPOSTCreated with default headers values
func NewImportMessagesUsingPOSTCreated() *ImportMessagesUsingPOSTCreated {
	return &ImportMessagesUsingPOSTCreated{}
}

/*ImportMessagesUsingPOSTCreated handles this case with default header values.

Created
*/
type ImportMessagesUsingPOSTCreated struct {
}

func (o *ImportMessagesUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/automata/api/engine/message/import][%d] importMessagesUsingPOSTCreated ", 201)
}

func (o *ImportMessagesUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportMessagesUsingPOSTUnauthorized creates a ImportMessagesUsingPOSTUnauthorized with default headers values
func NewImportMessagesUsingPOSTUnauthorized() *ImportMessagesUsingPOSTUnauthorized {
	return &ImportMessagesUsingPOSTUnauthorized{}
}

/*ImportMessagesUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ImportMessagesUsingPOSTUnauthorized struct {
}

func (o *ImportMessagesUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/automata/api/engine/message/import][%d] importMessagesUsingPOSTUnauthorized ", 401)
}

func (o *ImportMessagesUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportMessagesUsingPOSTForbidden creates a ImportMessagesUsingPOSTForbidden with default headers values
func NewImportMessagesUsingPOSTForbidden() *ImportMessagesUsingPOSTForbidden {
	return &ImportMessagesUsingPOSTForbidden{}
}

/*ImportMessagesUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ImportMessagesUsingPOSTForbidden struct {
}

func (o *ImportMessagesUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/automata/api/engine/message/import][%d] importMessagesUsingPOSTForbidden ", 403)
}

func (o *ImportMessagesUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportMessagesUsingPOSTNotFound creates a ImportMessagesUsingPOSTNotFound with default headers values
func NewImportMessagesUsingPOSTNotFound() *ImportMessagesUsingPOSTNotFound {
	return &ImportMessagesUsingPOSTNotFound{}
}

/*ImportMessagesUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ImportMessagesUsingPOSTNotFound struct {
}

func (o *ImportMessagesUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/automata/api/engine/message/import][%d] importMessagesUsingPOSTNotFound ", 404)
}

func (o *ImportMessagesUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}