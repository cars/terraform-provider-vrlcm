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

// QueueCommandUsingPOSTReader is a Reader for the QueueCommandUsingPOST structure.
type QueueCommandUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueueCommandUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueueCommandUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewQueueCommandUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewQueueCommandUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueueCommandUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewQueueCommandUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewQueueCommandUsingPOSTOK creates a QueueCommandUsingPOSTOK with default headers values
func NewQueueCommandUsingPOSTOK() *QueueCommandUsingPOSTOK {
	return &QueueCommandUsingPOSTOK{}
}

/*QueueCommandUsingPOSTOK handles this case with default header values.

OK
*/
type QueueCommandUsingPOSTOK struct {
	Payload *models.Status
}

func (o *QueueCommandUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /command/create][%d] queueCommandUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *QueueCommandUsingPOSTOK) GetPayload() *models.Status {
	return o.Payload
}

func (o *QueueCommandUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueueCommandUsingPOSTCreated creates a QueueCommandUsingPOSTCreated with default headers values
func NewQueueCommandUsingPOSTCreated() *QueueCommandUsingPOSTCreated {
	return &QueueCommandUsingPOSTCreated{}
}

/*QueueCommandUsingPOSTCreated handles this case with default header values.

Created
*/
type QueueCommandUsingPOSTCreated struct {
}

func (o *QueueCommandUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /command/create][%d] queueCommandUsingPOSTCreated ", 201)
}

func (o *QueueCommandUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueueCommandUsingPOSTUnauthorized creates a QueueCommandUsingPOSTUnauthorized with default headers values
func NewQueueCommandUsingPOSTUnauthorized() *QueueCommandUsingPOSTUnauthorized {
	return &QueueCommandUsingPOSTUnauthorized{}
}

/*QueueCommandUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type QueueCommandUsingPOSTUnauthorized struct {
}

func (o *QueueCommandUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /command/create][%d] queueCommandUsingPOSTUnauthorized ", 401)
}

func (o *QueueCommandUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueueCommandUsingPOSTForbidden creates a QueueCommandUsingPOSTForbidden with default headers values
func NewQueueCommandUsingPOSTForbidden() *QueueCommandUsingPOSTForbidden {
	return &QueueCommandUsingPOSTForbidden{}
}

/*QueueCommandUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type QueueCommandUsingPOSTForbidden struct {
}

func (o *QueueCommandUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /command/create][%d] queueCommandUsingPOSTForbidden ", 403)
}

func (o *QueueCommandUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQueueCommandUsingPOSTNotFound creates a QueueCommandUsingPOSTNotFound with default headers values
func NewQueueCommandUsingPOSTNotFound() *QueueCommandUsingPOSTNotFound {
	return &QueueCommandUsingPOSTNotFound{}
}

/*QueueCommandUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type QueueCommandUsingPOSTNotFound struct {
}

func (o *QueueCommandUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /command/create][%d] queueCommandUsingPOSTNotFound ", 404)
}

func (o *QueueCommandUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
