// Code generated by go-swagger; DO NOT EDIT.

package file_browse_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// ReplaceNodeUsingPUTReader is a Reader for the ReplaceNodeUsingPUT structure.
type ReplaceNodeUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceNodeUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceNodeUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewReplaceNodeUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewReplaceNodeUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReplaceNodeUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReplaceNodeUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceNodeUsingPUTOK creates a ReplaceNodeUsingPUTOK with default headers values
func NewReplaceNodeUsingPUTOK() *ReplaceNodeUsingPUTOK {
	return &ReplaceNodeUsingPUTOK{}
}

/*ReplaceNodeUsingPUTOK handles this case with default header values.

OK
*/
type ReplaceNodeUsingPUTOK struct {
	Payload *models.NodeDTO
}

func (o *ReplaceNodeUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /lcm/crepo/api/content/nodes][%d] replaceNodeUsingPUTOK  %+v", 200, o.Payload)
}

func (o *ReplaceNodeUsingPUTOK) GetPayload() *models.NodeDTO {
	return o.Payload
}

func (o *ReplaceNodeUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceNodeUsingPUTCreated creates a ReplaceNodeUsingPUTCreated with default headers values
func NewReplaceNodeUsingPUTCreated() *ReplaceNodeUsingPUTCreated {
	return &ReplaceNodeUsingPUTCreated{}
}

/*ReplaceNodeUsingPUTCreated handles this case with default header values.

Created
*/
type ReplaceNodeUsingPUTCreated struct {
}

func (o *ReplaceNodeUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /lcm/crepo/api/content/nodes][%d] replaceNodeUsingPUTCreated ", 201)
}

func (o *ReplaceNodeUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceNodeUsingPUTUnauthorized creates a ReplaceNodeUsingPUTUnauthorized with default headers values
func NewReplaceNodeUsingPUTUnauthorized() *ReplaceNodeUsingPUTUnauthorized {
	return &ReplaceNodeUsingPUTUnauthorized{}
}

/*ReplaceNodeUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type ReplaceNodeUsingPUTUnauthorized struct {
}

func (o *ReplaceNodeUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lcm/crepo/api/content/nodes][%d] replaceNodeUsingPUTUnauthorized ", 401)
}

func (o *ReplaceNodeUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceNodeUsingPUTForbidden creates a ReplaceNodeUsingPUTForbidden with default headers values
func NewReplaceNodeUsingPUTForbidden() *ReplaceNodeUsingPUTForbidden {
	return &ReplaceNodeUsingPUTForbidden{}
}

/*ReplaceNodeUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type ReplaceNodeUsingPUTForbidden struct {
}

func (o *ReplaceNodeUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /lcm/crepo/api/content/nodes][%d] replaceNodeUsingPUTForbidden ", 403)
}

func (o *ReplaceNodeUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceNodeUsingPUTNotFound creates a ReplaceNodeUsingPUTNotFound with default headers values
func NewReplaceNodeUsingPUTNotFound() *ReplaceNodeUsingPUTNotFound {
	return &ReplaceNodeUsingPUTNotFound{}
}

/*ReplaceNodeUsingPUTNotFound handles this case with default header values.

Not Found
*/
type ReplaceNodeUsingPUTNotFound struct {
}

func (o *ReplaceNodeUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /lcm/crepo/api/content/nodes][%d] replaceNodeUsingPUTNotFound ", 404)
}

func (o *ReplaceNodeUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
