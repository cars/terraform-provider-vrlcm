// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteGroupByIDUsingDELETEReader is a Reader for the DeleteGroupByIDUsingDELETE structure.
type DeleteGroupByIDUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupByIDUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGroupByIDUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteGroupByIDUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteGroupByIDUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGroupByIDUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGroupByIDUsingDELETEOK creates a DeleteGroupByIDUsingDELETEOK with default headers values
func NewDeleteGroupByIDUsingDELETEOK() *DeleteGroupByIDUsingDELETEOK {
	return &DeleteGroupByIDUsingDELETEOK{}
}

/*DeleteGroupByIDUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteGroupByIDUsingDELETEOK struct {
	Payload interface{}
}

func (o *DeleteGroupByIDUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/groups/{vmid}][%d] deleteGroupByIdUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGroupByIDUsingDELETEOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteGroupByIDUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupByIDUsingDELETENoContent creates a DeleteGroupByIDUsingDELETENoContent with default headers values
func NewDeleteGroupByIDUsingDELETENoContent() *DeleteGroupByIDUsingDELETENoContent {
	return &DeleteGroupByIDUsingDELETENoContent{}
}

/*DeleteGroupByIDUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteGroupByIDUsingDELETENoContent struct {
}

func (o *DeleteGroupByIDUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/groups/{vmid}][%d] deleteGroupByIdUsingDELETENoContent ", 204)
}

func (o *DeleteGroupByIDUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupByIDUsingDELETEUnauthorized creates a DeleteGroupByIDUsingDELETEUnauthorized with default headers values
func NewDeleteGroupByIDUsingDELETEUnauthorized() *DeleteGroupByIDUsingDELETEUnauthorized {
	return &DeleteGroupByIDUsingDELETEUnauthorized{}
}

/*DeleteGroupByIDUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteGroupByIDUsingDELETEUnauthorized struct {
}

func (o *DeleteGroupByIDUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/groups/{vmid}][%d] deleteGroupByIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteGroupByIDUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupByIDUsingDELETEForbidden creates a DeleteGroupByIDUsingDELETEForbidden with default headers values
func NewDeleteGroupByIDUsingDELETEForbidden() *DeleteGroupByIDUsingDELETEForbidden {
	return &DeleteGroupByIDUsingDELETEForbidden{}
}

/*DeleteGroupByIDUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteGroupByIDUsingDELETEForbidden struct {
}

func (o *DeleteGroupByIDUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/groups/{vmid}][%d] deleteGroupByIdUsingDELETEForbidden ", 403)
}

func (o *DeleteGroupByIDUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
