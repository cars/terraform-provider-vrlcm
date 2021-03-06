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

// AddroductDataSourcesUsingPOSTReader is a Reader for the AddroductDataSourcesUsingPOST structure.
type AddroductDataSourcesUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddroductDataSourcesUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddroductDataSourcesUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAddroductDataSourcesUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewAddroductDataSourcesUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddroductDataSourcesUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddroductDataSourcesUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddroductDataSourcesUsingPOSTOK creates a AddroductDataSourcesUsingPOSTOK with default headers values
func NewAddroductDataSourcesUsingPOSTOK() *AddroductDataSourcesUsingPOSTOK {
	return &AddroductDataSourcesUsingPOSTOK{}
}

/*AddroductDataSourcesUsingPOSTOK handles this case with default header values.

OK
*/
type AddroductDataSourcesUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *AddroductDataSourcesUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/adddatasource][%d] addroductDataSourcesUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *AddroductDataSourcesUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *AddroductDataSourcesUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddroductDataSourcesUsingPOSTCreated creates a AddroductDataSourcesUsingPOSTCreated with default headers values
func NewAddroductDataSourcesUsingPOSTCreated() *AddroductDataSourcesUsingPOSTCreated {
	return &AddroductDataSourcesUsingPOSTCreated{}
}

/*AddroductDataSourcesUsingPOSTCreated handles this case with default header values.

Created
*/
type AddroductDataSourcesUsingPOSTCreated struct {
}

func (o *AddroductDataSourcesUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/adddatasource][%d] addroductDataSourcesUsingPOSTCreated ", 201)
}

func (o *AddroductDataSourcesUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddroductDataSourcesUsingPOSTUnauthorized creates a AddroductDataSourcesUsingPOSTUnauthorized with default headers values
func NewAddroductDataSourcesUsingPOSTUnauthorized() *AddroductDataSourcesUsingPOSTUnauthorized {
	return &AddroductDataSourcesUsingPOSTUnauthorized{}
}

/*AddroductDataSourcesUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type AddroductDataSourcesUsingPOSTUnauthorized struct {
}

func (o *AddroductDataSourcesUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/adddatasource][%d] addroductDataSourcesUsingPOSTUnauthorized ", 401)
}

func (o *AddroductDataSourcesUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddroductDataSourcesUsingPOSTForbidden creates a AddroductDataSourcesUsingPOSTForbidden with default headers values
func NewAddroductDataSourcesUsingPOSTForbidden() *AddroductDataSourcesUsingPOSTForbidden {
	return &AddroductDataSourcesUsingPOSTForbidden{}
}

/*AddroductDataSourcesUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type AddroductDataSourcesUsingPOSTForbidden struct {
}

func (o *AddroductDataSourcesUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/adddatasource][%d] addroductDataSourcesUsingPOSTForbidden ", 403)
}

func (o *AddroductDataSourcesUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddroductDataSourcesUsingPOSTNotFound creates a AddroductDataSourcesUsingPOSTNotFound with default headers values
func NewAddroductDataSourcesUsingPOSTNotFound() *AddroductDataSourcesUsingPOSTNotFound {
	return &AddroductDataSourcesUsingPOSTNotFound{}
}

/*AddroductDataSourcesUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type AddroductDataSourcesUsingPOSTNotFound struct {
}

func (o *AddroductDataSourcesUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/adddatasource][%d] addroductDataSourcesUsingPOSTNotFound ", 404)
}

func (o *AddroductDataSourcesUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
