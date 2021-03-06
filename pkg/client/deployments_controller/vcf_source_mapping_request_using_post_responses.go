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

// VcfSourceMappingRequestUsingPOSTReader is a Reader for the VcfSourceMappingRequestUsingPOST structure.
type VcfSourceMappingRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VcfSourceMappingRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVcfSourceMappingRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewVcfSourceMappingRequestUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVcfSourceMappingRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVcfSourceMappingRequestUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVcfSourceMappingRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVcfSourceMappingRequestUsingPOSTOK creates a VcfSourceMappingRequestUsingPOSTOK with default headers values
func NewVcfSourceMappingRequestUsingPOSTOK() *VcfSourceMappingRequestUsingPOSTOK {
	return &VcfSourceMappingRequestUsingPOSTOK{}
}

/*VcfSourceMappingRequestUsingPOSTOK handles this case with default header values.

OK
*/
type VcfSourceMappingRequestUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *VcfSourceMappingRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/vcfSourceMapping][%d] vcfSourceMappingRequestUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *VcfSourceMappingRequestUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *VcfSourceMappingRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVcfSourceMappingRequestUsingPOSTCreated creates a VcfSourceMappingRequestUsingPOSTCreated with default headers values
func NewVcfSourceMappingRequestUsingPOSTCreated() *VcfSourceMappingRequestUsingPOSTCreated {
	return &VcfSourceMappingRequestUsingPOSTCreated{}
}

/*VcfSourceMappingRequestUsingPOSTCreated handles this case with default header values.

Created
*/
type VcfSourceMappingRequestUsingPOSTCreated struct {
}

func (o *VcfSourceMappingRequestUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/vcfSourceMapping][%d] vcfSourceMappingRequestUsingPOSTCreated ", 201)
}

func (o *VcfSourceMappingRequestUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVcfSourceMappingRequestUsingPOSTUnauthorized creates a VcfSourceMappingRequestUsingPOSTUnauthorized with default headers values
func NewVcfSourceMappingRequestUsingPOSTUnauthorized() *VcfSourceMappingRequestUsingPOSTUnauthorized {
	return &VcfSourceMappingRequestUsingPOSTUnauthorized{}
}

/*VcfSourceMappingRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type VcfSourceMappingRequestUsingPOSTUnauthorized struct {
}

func (o *VcfSourceMappingRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/vcfSourceMapping][%d] vcfSourceMappingRequestUsingPOSTUnauthorized ", 401)
}

func (o *VcfSourceMappingRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVcfSourceMappingRequestUsingPOSTForbidden creates a VcfSourceMappingRequestUsingPOSTForbidden with default headers values
func NewVcfSourceMappingRequestUsingPOSTForbidden() *VcfSourceMappingRequestUsingPOSTForbidden {
	return &VcfSourceMappingRequestUsingPOSTForbidden{}
}

/*VcfSourceMappingRequestUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type VcfSourceMappingRequestUsingPOSTForbidden struct {
}

func (o *VcfSourceMappingRequestUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/vcfSourceMapping][%d] vcfSourceMappingRequestUsingPOSTForbidden ", 403)
}

func (o *VcfSourceMappingRequestUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVcfSourceMappingRequestUsingPOSTNotFound creates a VcfSourceMappingRequestUsingPOSTNotFound with default headers values
func NewVcfSourceMappingRequestUsingPOSTNotFound() *VcfSourceMappingRequestUsingPOSTNotFound {
	return &VcfSourceMappingRequestUsingPOSTNotFound{}
}

/*VcfSourceMappingRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type VcfSourceMappingRequestUsingPOSTNotFound struct {
}

func (o *VcfSourceMappingRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/vcfSourceMapping][%d] vcfSourceMappingRequestUsingPOSTNotFound ", 404)
}

func (o *VcfSourceMappingRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
