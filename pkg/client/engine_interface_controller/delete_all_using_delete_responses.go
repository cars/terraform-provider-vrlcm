// Code generated by go-swagger; DO NOT EDIT.

package engine_interface_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAllUsingDELETEReader is a Reader for the DeleteAllUsingDELETE structure.
type DeleteAllUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAllUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteAllUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAllUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAllUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAllUsingDELETEOK creates a DeleteAllUsingDELETEOK with default headers values
func NewDeleteAllUsingDELETEOK() *DeleteAllUsingDELETEOK {
	return &DeleteAllUsingDELETEOK{}
}

/*DeleteAllUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteAllUsingDELETEOK struct {
}

func (o *DeleteAllUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /lcm/automata/api/engine/wip/statemachine/reset][%d] deleteAllUsingDELETEOK ", 200)
}

func (o *DeleteAllUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAllUsingDELETENoContent creates a DeleteAllUsingDELETENoContent with default headers values
func NewDeleteAllUsingDELETENoContent() *DeleteAllUsingDELETENoContent {
	return &DeleteAllUsingDELETENoContent{}
}

/*DeleteAllUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteAllUsingDELETENoContent struct {
}

func (o *DeleteAllUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /lcm/automata/api/engine/wip/statemachine/reset][%d] deleteAllUsingDELETENoContent ", 204)
}

func (o *DeleteAllUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAllUsingDELETEUnauthorized creates a DeleteAllUsingDELETEUnauthorized with default headers values
func NewDeleteAllUsingDELETEUnauthorized() *DeleteAllUsingDELETEUnauthorized {
	return &DeleteAllUsingDELETEUnauthorized{}
}

/*DeleteAllUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAllUsingDELETEUnauthorized struct {
}

func (o *DeleteAllUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lcm/automata/api/engine/wip/statemachine/reset][%d] deleteAllUsingDELETEUnauthorized ", 401)
}

func (o *DeleteAllUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAllUsingDELETEForbidden creates a DeleteAllUsingDELETEForbidden with default headers values
func NewDeleteAllUsingDELETEForbidden() *DeleteAllUsingDELETEForbidden {
	return &DeleteAllUsingDELETEForbidden{}
}

/*DeleteAllUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteAllUsingDELETEForbidden struct {
}

func (o *DeleteAllUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lcm/automata/api/engine/wip/statemachine/reset][%d] deleteAllUsingDELETEForbidden ", 403)
}

func (o *DeleteAllUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
