// Code generated by go-swagger; DO NOT EDIT.

package content_download_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RenameContentFileUsingPOSTReader is a Reader for the RenameContentFileUsingPOST structure.
type RenameContentFileUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenameContentFileUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenameContentFileUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRenameContentFileUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRenameContentFileUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRenameContentFileUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRenameContentFileUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRenameContentFileUsingPOSTOK creates a RenameContentFileUsingPOSTOK with default headers values
func NewRenameContentFileUsingPOSTOK() *RenameContentFileUsingPOSTOK {
	return &RenameContentFileUsingPOSTOK{}
}

/*RenameContentFileUsingPOSTOK handles this case with default header values.

OK
*/
type RenameContentFileUsingPOSTOK struct {
	Payload interface{}
}

func (o *RenameContentFileUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/rename][%d] renameContentFileUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *RenameContentFileUsingPOSTOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RenameContentFileUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenameContentFileUsingPOSTCreated creates a RenameContentFileUsingPOSTCreated with default headers values
func NewRenameContentFileUsingPOSTCreated() *RenameContentFileUsingPOSTCreated {
	return &RenameContentFileUsingPOSTCreated{}
}

/*RenameContentFileUsingPOSTCreated handles this case with default header values.

Created
*/
type RenameContentFileUsingPOSTCreated struct {
}

func (o *RenameContentFileUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/rename][%d] renameContentFileUsingPOSTCreated ", 201)
}

func (o *RenameContentFileUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameContentFileUsingPOSTUnauthorized creates a RenameContentFileUsingPOSTUnauthorized with default headers values
func NewRenameContentFileUsingPOSTUnauthorized() *RenameContentFileUsingPOSTUnauthorized {
	return &RenameContentFileUsingPOSTUnauthorized{}
}

/*RenameContentFileUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type RenameContentFileUsingPOSTUnauthorized struct {
}

func (o *RenameContentFileUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/rename][%d] renameContentFileUsingPOSTUnauthorized ", 401)
}

func (o *RenameContentFileUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameContentFileUsingPOSTForbidden creates a RenameContentFileUsingPOSTForbidden with default headers values
func NewRenameContentFileUsingPOSTForbidden() *RenameContentFileUsingPOSTForbidden {
	return &RenameContentFileUsingPOSTForbidden{}
}

/*RenameContentFileUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type RenameContentFileUsingPOSTForbidden struct {
}

func (o *RenameContentFileUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/rename][%d] renameContentFileUsingPOSTForbidden ", 403)
}

func (o *RenameContentFileUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameContentFileUsingPOSTNotFound creates a RenameContentFileUsingPOSTNotFound with default headers values
func NewRenameContentFileUsingPOSTNotFound() *RenameContentFileUsingPOSTNotFound {
	return &RenameContentFileUsingPOSTNotFound{}
}

/*RenameContentFileUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type RenameContentFileUsingPOSTNotFound struct {
}

func (o *RenameContentFileUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/rename][%d] renameContentFileUsingPOSTNotFound ", 404)
}

func (o *RenameContentFileUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
