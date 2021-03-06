// Code generated by go-swagger; DO NOT EDIT.

package deployments_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// CreateProductSnapshotRequestUsingPOSTReader is a Reader for the CreateProductSnapshotRequestUsingPOST structure.
type CreateProductSnapshotRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProductSnapshotRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProductSnapshotRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateProductSnapshotRequestUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateProductSnapshotRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProductSnapshotRequestUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProductSnapshotRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProductSnapshotRequestUsingPOSTOK creates a CreateProductSnapshotRequestUsingPOSTOK with default headers values
func NewCreateProductSnapshotRequestUsingPOSTOK() *CreateProductSnapshotRequestUsingPOSTOK {
	return &CreateProductSnapshotRequestUsingPOSTOK{}
}

/*CreateProductSnapshotRequestUsingPOSTOK handles this case with default header values.

OK
*/
type CreateProductSnapshotRequestUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *CreateProductSnapshotRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/snapshot][%d] createProductSnapshotRequestUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateProductSnapshotRequestUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *CreateProductSnapshotRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductSnapshotRequestUsingPOSTCreated creates a CreateProductSnapshotRequestUsingPOSTCreated with default headers values
func NewCreateProductSnapshotRequestUsingPOSTCreated() *CreateProductSnapshotRequestUsingPOSTCreated {
	return &CreateProductSnapshotRequestUsingPOSTCreated{}
}

/*CreateProductSnapshotRequestUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateProductSnapshotRequestUsingPOSTCreated struct {
}

func (o *CreateProductSnapshotRequestUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/snapshot][%d] createProductSnapshotRequestUsingPOSTCreated ", 201)
}

func (o *CreateProductSnapshotRequestUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductSnapshotRequestUsingPOSTUnauthorized creates a CreateProductSnapshotRequestUsingPOSTUnauthorized with default headers values
func NewCreateProductSnapshotRequestUsingPOSTUnauthorized() *CreateProductSnapshotRequestUsingPOSTUnauthorized {
	return &CreateProductSnapshotRequestUsingPOSTUnauthorized{}
}

/*CreateProductSnapshotRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateProductSnapshotRequestUsingPOSTUnauthorized struct {
}

func (o *CreateProductSnapshotRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/snapshot][%d] createProductSnapshotRequestUsingPOSTUnauthorized ", 401)
}

func (o *CreateProductSnapshotRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductSnapshotRequestUsingPOSTForbidden creates a CreateProductSnapshotRequestUsingPOSTForbidden with default headers values
func NewCreateProductSnapshotRequestUsingPOSTForbidden() *CreateProductSnapshotRequestUsingPOSTForbidden {
	return &CreateProductSnapshotRequestUsingPOSTForbidden{}
}

/*CreateProductSnapshotRequestUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateProductSnapshotRequestUsingPOSTForbidden struct {
}

func (o *CreateProductSnapshotRequestUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/snapshot][%d] createProductSnapshotRequestUsingPOSTForbidden ", 403)
}

func (o *CreateProductSnapshotRequestUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductSnapshotRequestUsingPOSTNotFound creates a CreateProductSnapshotRequestUsingPOSTNotFound with default headers values
func NewCreateProductSnapshotRequestUsingPOSTNotFound() *CreateProductSnapshotRequestUsingPOSTNotFound {
	return &CreateProductSnapshotRequestUsingPOSTNotFound{}
}

/*CreateProductSnapshotRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateProductSnapshotRequestUsingPOSTNotFound struct {
}

func (o *CreateProductSnapshotRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/snapshot][%d] createProductSnapshotRequestUsingPOSTNotFound ", 404)
}

func (o *CreateProductSnapshotRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
