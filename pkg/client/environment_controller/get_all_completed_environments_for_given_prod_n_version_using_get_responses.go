// Code generated by go-swagger; DO NOT EDIT.

package environment_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETReader is a Reader for the GetAllCompletedEnvironmentsForGivenProdNVersionUsingGET structure.
type GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK creates a GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK with default headers values
func NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK() *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK {
	return &GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK{}
}

/*GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK handles this case with default header values.

OK
*/
type GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK struct {
	Payload []*models.EnvironmentRequestDTO
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/prodId/{prodId}/prodVersion/{prodVersion}][%d] getAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK) GetPayload() []*models.EnvironmentRequestDTO {
	return o.Payload
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized creates a GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized with default headers values
func NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized() *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized {
	return &GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized{}
}

/*GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized struct {
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/prodId/{prodId}/prodVersion/{prodVersion}][%d] getAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized ", 401)
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden creates a GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden with default headers values
func NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden() *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden {
	return &GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden{}
}

/*GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden struct {
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/prodId/{prodId}/prodVersion/{prodVersion}][%d] getAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden ", 403)
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound creates a GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound with default headers values
func NewGetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound() *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound {
	return &GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound{}
}

/*GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound struct {
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/environments/prodId/{prodId}/prodVersion/{prodVersion}][%d] getAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound ", 404)
}

func (o *GetAllCompletedEnvironmentsForGivenProdNVersionUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
