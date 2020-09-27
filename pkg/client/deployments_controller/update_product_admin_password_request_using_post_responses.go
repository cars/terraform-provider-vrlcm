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

// UpdateProductAdminPasswordRequestUsingPOSTReader is a Reader for the UpdateProductAdminPasswordRequestUsingPOST structure.
type UpdateProductAdminPasswordRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductAdminPasswordRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProductAdminPasswordRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateProductAdminPasswordRequestUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateProductAdminPasswordRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateProductAdminPasswordRequestUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProductAdminPasswordRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductAdminPasswordRequestUsingPOSTOK creates a UpdateProductAdminPasswordRequestUsingPOSTOK with default headers values
func NewUpdateProductAdminPasswordRequestUsingPOSTOK() *UpdateProductAdminPasswordRequestUsingPOSTOK {
	return &UpdateProductAdminPasswordRequestUsingPOSTOK{}
}

/*UpdateProductAdminPasswordRequestUsingPOSTOK handles this case with default header values.

OK
*/
type UpdateProductAdminPasswordRequestUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/updateadminpassword][%d] updateProductAdminPasswordRequestUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProductAdminPasswordRequestUsingPOSTCreated creates a UpdateProductAdminPasswordRequestUsingPOSTCreated with default headers values
func NewUpdateProductAdminPasswordRequestUsingPOSTCreated() *UpdateProductAdminPasswordRequestUsingPOSTCreated {
	return &UpdateProductAdminPasswordRequestUsingPOSTCreated{}
}

/*UpdateProductAdminPasswordRequestUsingPOSTCreated handles this case with default header values.

Created
*/
type UpdateProductAdminPasswordRequestUsingPOSTCreated struct {
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/updateadminpassword][%d] updateProductAdminPasswordRequestUsingPOSTCreated ", 201)
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductAdminPasswordRequestUsingPOSTUnauthorized creates a UpdateProductAdminPasswordRequestUsingPOSTUnauthorized with default headers values
func NewUpdateProductAdminPasswordRequestUsingPOSTUnauthorized() *UpdateProductAdminPasswordRequestUsingPOSTUnauthorized {
	return &UpdateProductAdminPasswordRequestUsingPOSTUnauthorized{}
}

/*UpdateProductAdminPasswordRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateProductAdminPasswordRequestUsingPOSTUnauthorized struct {
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/updateadminpassword][%d] updateProductAdminPasswordRequestUsingPOSTUnauthorized ", 401)
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductAdminPasswordRequestUsingPOSTForbidden creates a UpdateProductAdminPasswordRequestUsingPOSTForbidden with default headers values
func NewUpdateProductAdminPasswordRequestUsingPOSTForbidden() *UpdateProductAdminPasswordRequestUsingPOSTForbidden {
	return &UpdateProductAdminPasswordRequestUsingPOSTForbidden{}
}

/*UpdateProductAdminPasswordRequestUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type UpdateProductAdminPasswordRequestUsingPOSTForbidden struct {
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/updateadminpassword][%d] updateProductAdminPasswordRequestUsingPOSTForbidden ", 403)
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProductAdminPasswordRequestUsingPOSTNotFound creates a UpdateProductAdminPasswordRequestUsingPOSTNotFound with default headers values
func NewUpdateProductAdminPasswordRequestUsingPOSTNotFound() *UpdateProductAdminPasswordRequestUsingPOSTNotFound {
	return &UpdateProductAdminPasswordRequestUsingPOSTNotFound{}
}

/*UpdateProductAdminPasswordRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type UpdateProductAdminPasswordRequestUsingPOSTNotFound struct {
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/updateadminpassword][%d] updateProductAdminPasswordRequestUsingPOSTNotFound ", 404)
}

func (o *UpdateProductAdminPasswordRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}