// Code generated by go-swagger; DO NOT EDIT.

package file_browse_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ImportNodeUsingPOSTReader is a Reader for the ImportNodeUsingPOST structure.
type ImportNodeUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImportNodeUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImportNodeUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewImportNodeUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewImportNodeUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImportNodeUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImportNodeUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImportNodeUsingPOSTOK creates a ImportNodeUsingPOSTOK with default headers values
func NewImportNodeUsingPOSTOK() *ImportNodeUsingPOSTOK {
	return &ImportNodeUsingPOSTOK{}
}

/*ImportNodeUsingPOSTOK handles this case with default header values.

OK
*/
type ImportNodeUsingPOSTOK struct {
}

func (o *ImportNodeUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/nodes/import][%d] importNodeUsingPOSTOK ", 200)
}

func (o *ImportNodeUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportNodeUsingPOSTCreated creates a ImportNodeUsingPOSTCreated with default headers values
func NewImportNodeUsingPOSTCreated() *ImportNodeUsingPOSTCreated {
	return &ImportNodeUsingPOSTCreated{}
}

/*ImportNodeUsingPOSTCreated handles this case with default header values.

Created
*/
type ImportNodeUsingPOSTCreated struct {
}

func (o *ImportNodeUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/nodes/import][%d] importNodeUsingPOSTCreated ", 201)
}

func (o *ImportNodeUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportNodeUsingPOSTUnauthorized creates a ImportNodeUsingPOSTUnauthorized with default headers values
func NewImportNodeUsingPOSTUnauthorized() *ImportNodeUsingPOSTUnauthorized {
	return &ImportNodeUsingPOSTUnauthorized{}
}

/*ImportNodeUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ImportNodeUsingPOSTUnauthorized struct {
}

func (o *ImportNodeUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/nodes/import][%d] importNodeUsingPOSTUnauthorized ", 401)
}

func (o *ImportNodeUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportNodeUsingPOSTForbidden creates a ImportNodeUsingPOSTForbidden with default headers values
func NewImportNodeUsingPOSTForbidden() *ImportNodeUsingPOSTForbidden {
	return &ImportNodeUsingPOSTForbidden{}
}

/*ImportNodeUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type ImportNodeUsingPOSTForbidden struct {
}

func (o *ImportNodeUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/nodes/import][%d] importNodeUsingPOSTForbidden ", 403)
}

func (o *ImportNodeUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewImportNodeUsingPOSTNotFound creates a ImportNodeUsingPOSTNotFound with default headers values
func NewImportNodeUsingPOSTNotFound() *ImportNodeUsingPOSTNotFound {
	return &ImportNodeUsingPOSTNotFound{}
}

/*ImportNodeUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ImportNodeUsingPOSTNotFound struct {
}

func (o *ImportNodeUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/crepo/api/content/nodes/import][%d] importNodeUsingPOSTNotFound ", 404)
}

func (o *ImportNodeUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
