// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddNodeUsingPOSTReader is a Reader for the AddNodeUsingPOST structure.
type AddNodeUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddNodeUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddNodeUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAddNodeUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddNodeUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddNodeUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddNodeUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddNodeUsingPOSTOK creates a AddNodeUsingPOSTOK with default headers values
func NewAddNodeUsingPOSTOK() *AddNodeUsingPOSTOK {
	return &AddNodeUsingPOSTOK{}
}

/*AddNodeUsingPOSTOK handles this case with default header values.

OK
*/
type AddNodeUsingPOSTOK struct {
}

func (o *AddNodeUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productName}/nodes][%d] addNodeUsingPOSTOK ", 200)
}

func (o *AddNodeUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodeUsingPOSTCreated creates a AddNodeUsingPOSTCreated with default headers values
func NewAddNodeUsingPOSTCreated() *AddNodeUsingPOSTCreated {
	return &AddNodeUsingPOSTCreated{}
}

/*AddNodeUsingPOSTCreated handles this case with default header values.

Created
*/
type AddNodeUsingPOSTCreated struct {
}

func (o *AddNodeUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productName}/nodes][%d] addNodeUsingPOSTCreated ", 201)
}

func (o *AddNodeUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodeUsingPOSTUnauthorized creates a AddNodeUsingPOSTUnauthorized with default headers values
func NewAddNodeUsingPOSTUnauthorized() *AddNodeUsingPOSTUnauthorized {
	return &AddNodeUsingPOSTUnauthorized{}
}

/*AddNodeUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type AddNodeUsingPOSTUnauthorized struct {
}

func (o *AddNodeUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productName}/nodes][%d] addNodeUsingPOSTUnauthorized ", 401)
}

func (o *AddNodeUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodeUsingPOSTForbidden creates a AddNodeUsingPOSTForbidden with default headers values
func NewAddNodeUsingPOSTForbidden() *AddNodeUsingPOSTForbidden {
	return &AddNodeUsingPOSTForbidden{}
}

/*AddNodeUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type AddNodeUsingPOSTForbidden struct {
}

func (o *AddNodeUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productName}/nodes][%d] addNodeUsingPOSTForbidden ", 403)
}

func (o *AddNodeUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddNodeUsingPOSTNotFound creates a AddNodeUsingPOSTNotFound with default headers values
func NewAddNodeUsingPOSTNotFound() *AddNodeUsingPOSTNotFound {
	return &AddNodeUsingPOSTNotFound{}
}

/*AddNodeUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type AddNodeUsingPOSTNotFound struct {
}

func (o *AddNodeUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productName}/nodes][%d] addNodeUsingPOSTNotFound ", 404)
}

func (o *AddNodeUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
