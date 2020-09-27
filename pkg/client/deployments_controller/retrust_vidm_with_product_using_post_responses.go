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

// RetrustVidmWithProductUsingPOSTReader is a Reader for the RetrustVidmWithProductUsingPOST structure.
type RetrustVidmWithProductUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrustVidmWithProductUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrustVidmWithProductUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRetrustVidmWithProductUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRetrustVidmWithProductUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrustVidmWithProductUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRetrustVidmWithProductUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrustVidmWithProductUsingPOSTOK creates a RetrustVidmWithProductUsingPOSTOK with default headers values
func NewRetrustVidmWithProductUsingPOSTOK() *RetrustVidmWithProductUsingPOSTOK {
	return &RetrustVidmWithProductUsingPOSTOK{}
}

/*RetrustVidmWithProductUsingPOSTOK handles this case with default header values.

OK
*/
type RetrustVidmWithProductUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *RetrustVidmWithProductUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/productVidmRetrust][%d] retrustVidmWithProductUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *RetrustVidmWithProductUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *RetrustVidmWithProductUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrustVidmWithProductUsingPOSTCreated creates a RetrustVidmWithProductUsingPOSTCreated with default headers values
func NewRetrustVidmWithProductUsingPOSTCreated() *RetrustVidmWithProductUsingPOSTCreated {
	return &RetrustVidmWithProductUsingPOSTCreated{}
}

/*RetrustVidmWithProductUsingPOSTCreated handles this case with default header values.

Created
*/
type RetrustVidmWithProductUsingPOSTCreated struct {
}

func (o *RetrustVidmWithProductUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/productVidmRetrust][%d] retrustVidmWithProductUsingPOSTCreated ", 201)
}

func (o *RetrustVidmWithProductUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrustVidmWithProductUsingPOSTUnauthorized creates a RetrustVidmWithProductUsingPOSTUnauthorized with default headers values
func NewRetrustVidmWithProductUsingPOSTUnauthorized() *RetrustVidmWithProductUsingPOSTUnauthorized {
	return &RetrustVidmWithProductUsingPOSTUnauthorized{}
}

/*RetrustVidmWithProductUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type RetrustVidmWithProductUsingPOSTUnauthorized struct {
}

func (o *RetrustVidmWithProductUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/productVidmRetrust][%d] retrustVidmWithProductUsingPOSTUnauthorized ", 401)
}

func (o *RetrustVidmWithProductUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrustVidmWithProductUsingPOSTForbidden creates a RetrustVidmWithProductUsingPOSTForbidden with default headers values
func NewRetrustVidmWithProductUsingPOSTForbidden() *RetrustVidmWithProductUsingPOSTForbidden {
	return &RetrustVidmWithProductUsingPOSTForbidden{}
}

/*RetrustVidmWithProductUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type RetrustVidmWithProductUsingPOSTForbidden struct {
}

func (o *RetrustVidmWithProductUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/productVidmRetrust][%d] retrustVidmWithProductUsingPOSTForbidden ", 403)
}

func (o *RetrustVidmWithProductUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrustVidmWithProductUsingPOSTNotFound creates a RetrustVidmWithProductUsingPOSTNotFound with default headers values
func NewRetrustVidmWithProductUsingPOSTNotFound() *RetrustVidmWithProductUsingPOSTNotFound {
	return &RetrustVidmWithProductUsingPOSTNotFound{}
}

/*RetrustVidmWithProductUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type RetrustVidmWithProductUsingPOSTNotFound struct {
}

func (o *RetrustVidmWithProductUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/productVidmRetrust][%d] retrustVidmWithProductUsingPOSTNotFound ", 404)
}

func (o *RetrustVidmWithProductUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
