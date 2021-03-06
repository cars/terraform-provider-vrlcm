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

// RetrustLBCertificateUsingPOSTReader is a Reader for the RetrustLBCertificateUsingPOST structure.
type RetrustLBCertificateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrustLBCertificateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRetrustLBCertificateUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRetrustLBCertificateUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRetrustLBCertificateUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRetrustLBCertificateUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRetrustLBCertificateUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRetrustLBCertificateUsingPOSTOK creates a RetrustLBCertificateUsingPOSTOK with default headers values
func NewRetrustLBCertificateUsingPOSTOK() *RetrustLBCertificateUsingPOSTOK {
	return &RetrustLBCertificateUsingPOSTOK{}
}

/*RetrustLBCertificateUsingPOSTOK handles this case with default header values.

OK
*/
type RetrustLBCertificateUsingPOSTOK struct {
	Payload *models.GenericRequestResponse
}

func (o *RetrustLBCertificateUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/retrustLoadBalancerCertificate][%d] retrustLBCertificateUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *RetrustLBCertificateUsingPOSTOK) GetPayload() *models.GenericRequestResponse {
	return o.Payload
}

func (o *RetrustLBCertificateUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrustLBCertificateUsingPOSTCreated creates a RetrustLBCertificateUsingPOSTCreated with default headers values
func NewRetrustLBCertificateUsingPOSTCreated() *RetrustLBCertificateUsingPOSTCreated {
	return &RetrustLBCertificateUsingPOSTCreated{}
}

/*RetrustLBCertificateUsingPOSTCreated handles this case with default header values.

Created
*/
type RetrustLBCertificateUsingPOSTCreated struct {
}

func (o *RetrustLBCertificateUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/retrustLoadBalancerCertificate][%d] retrustLBCertificateUsingPOSTCreated ", 201)
}

func (o *RetrustLBCertificateUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrustLBCertificateUsingPOSTUnauthorized creates a RetrustLBCertificateUsingPOSTUnauthorized with default headers values
func NewRetrustLBCertificateUsingPOSTUnauthorized() *RetrustLBCertificateUsingPOSTUnauthorized {
	return &RetrustLBCertificateUsingPOSTUnauthorized{}
}

/*RetrustLBCertificateUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type RetrustLBCertificateUsingPOSTUnauthorized struct {
}

func (o *RetrustLBCertificateUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/retrustLoadBalancerCertificate][%d] retrustLBCertificateUsingPOSTUnauthorized ", 401)
}

func (o *RetrustLBCertificateUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrustLBCertificateUsingPOSTForbidden creates a RetrustLBCertificateUsingPOSTForbidden with default headers values
func NewRetrustLBCertificateUsingPOSTForbidden() *RetrustLBCertificateUsingPOSTForbidden {
	return &RetrustLBCertificateUsingPOSTForbidden{}
}

/*RetrustLBCertificateUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type RetrustLBCertificateUsingPOSTForbidden struct {
}

func (o *RetrustLBCertificateUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/retrustLoadBalancerCertificate][%d] retrustLBCertificateUsingPOSTForbidden ", 403)
}

func (o *RetrustLBCertificateUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrustLBCertificateUsingPOSTNotFound creates a RetrustLBCertificateUsingPOSTNotFound with default headers values
func NewRetrustLBCertificateUsingPOSTNotFound() *RetrustLBCertificateUsingPOSTNotFound {
	return &RetrustLBCertificateUsingPOSTNotFound{}
}

/*RetrustLBCertificateUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type RetrustLBCertificateUsingPOSTNotFound struct {
}

func (o *RetrustLBCertificateUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/environments/{environmentId}/products/{productId}/retrustLoadBalancerCertificate][%d] retrustLBCertificateUsingPOSTNotFound ", 404)
}

func (o *RetrustLBCertificateUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
