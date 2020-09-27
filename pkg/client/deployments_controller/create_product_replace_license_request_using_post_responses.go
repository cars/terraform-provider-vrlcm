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

// CreateProductReplaceLicenseRequestUsingPOSTReader is a Reader for the CreateProductReplaceLicenseRequestUsingPOST structure.
type CreateProductReplaceLicenseRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProductReplaceLicenseRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateProductReplaceLicenseRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateProductReplaceLicenseRequestUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateProductReplaceLicenseRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProductReplaceLicenseRequestUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProductReplaceLicenseRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProductReplaceLicenseRequestUsingPOSTOK creates a CreateProductReplaceLicenseRequestUsingPOSTOK with default headers values
func NewCreateProductReplaceLicenseRequestUsingPOSTOK() *CreateProductReplaceLicenseRequestUsingPOSTOK {
	return &CreateProductReplaceLicenseRequestUsingPOSTOK{}
}

/*CreateProductReplaceLicenseRequestUsingPOSTOK handles this case with default header values.

OK
*/
type CreateProductReplaceLicenseRequestUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/replacelicense][%d] createProductReplaceLicenseRequestUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProductReplaceLicenseRequestUsingPOSTCreated creates a CreateProductReplaceLicenseRequestUsingPOSTCreated with default headers values
func NewCreateProductReplaceLicenseRequestUsingPOSTCreated() *CreateProductReplaceLicenseRequestUsingPOSTCreated {
	return &CreateProductReplaceLicenseRequestUsingPOSTCreated{}
}

/*CreateProductReplaceLicenseRequestUsingPOSTCreated handles this case with default header values.

Created
*/
type CreateProductReplaceLicenseRequestUsingPOSTCreated struct {
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/replacelicense][%d] createProductReplaceLicenseRequestUsingPOSTCreated ", 201)
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductReplaceLicenseRequestUsingPOSTUnauthorized creates a CreateProductReplaceLicenseRequestUsingPOSTUnauthorized with default headers values
func NewCreateProductReplaceLicenseRequestUsingPOSTUnauthorized() *CreateProductReplaceLicenseRequestUsingPOSTUnauthorized {
	return &CreateProductReplaceLicenseRequestUsingPOSTUnauthorized{}
}

/*CreateProductReplaceLicenseRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateProductReplaceLicenseRequestUsingPOSTUnauthorized struct {
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/replacelicense][%d] createProductReplaceLicenseRequestUsingPOSTUnauthorized ", 401)
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductReplaceLicenseRequestUsingPOSTForbidden creates a CreateProductReplaceLicenseRequestUsingPOSTForbidden with default headers values
func NewCreateProductReplaceLicenseRequestUsingPOSTForbidden() *CreateProductReplaceLicenseRequestUsingPOSTForbidden {
	return &CreateProductReplaceLicenseRequestUsingPOSTForbidden{}
}

/*CreateProductReplaceLicenseRequestUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateProductReplaceLicenseRequestUsingPOSTForbidden struct {
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/replacelicense][%d] createProductReplaceLicenseRequestUsingPOSTForbidden ", 403)
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateProductReplaceLicenseRequestUsingPOSTNotFound creates a CreateProductReplaceLicenseRequestUsingPOSTNotFound with default headers values
func NewCreateProductReplaceLicenseRequestUsingPOSTNotFound() *CreateProductReplaceLicenseRequestUsingPOSTNotFound {
	return &CreateProductReplaceLicenseRequestUsingPOSTNotFound{}
}

/*CreateProductReplaceLicenseRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateProductReplaceLicenseRequestUsingPOSTNotFound struct {
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/replacelicense][%d] createProductReplaceLicenseRequestUsingPOSTNotFound ", 404)
}

func (o *CreateProductReplaceLicenseRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
