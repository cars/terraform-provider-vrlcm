// Code generated by go-swagger; DO NOT EDIT.

package data_center_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetAllRegionsUsingGETReader is a Reader for the GetAllRegionsUsingGET structure.
type GetAllRegionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllRegionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllRegionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllRegionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllRegionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllRegionsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllRegionsUsingGETOK creates a GetAllRegionsUsingGETOK with default headers values
func NewGetAllRegionsUsingGETOK() *GetAllRegionsUsingGETOK {
	return &GetAllRegionsUsingGETOK{}
}

/*GetAllRegionsUsingGETOK handles this case with default header values.

OK
*/
type GetAllRegionsUsingGETOK struct {
	Payload []*models.RegionRequestDTO
}

func (o *GetAllRegionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions][%d] getAllRegionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllRegionsUsingGETOK) GetPayload() []*models.RegionRequestDTO {
	return o.Payload
}

func (o *GetAllRegionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllRegionsUsingGETUnauthorized creates a GetAllRegionsUsingGETUnauthorized with default headers values
func NewGetAllRegionsUsingGETUnauthorized() *GetAllRegionsUsingGETUnauthorized {
	return &GetAllRegionsUsingGETUnauthorized{}
}

/*GetAllRegionsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllRegionsUsingGETUnauthorized struct {
}

func (o *GetAllRegionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions][%d] getAllRegionsUsingGETUnauthorized ", 401)
}

func (o *GetAllRegionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllRegionsUsingGETForbidden creates a GetAllRegionsUsingGETForbidden with default headers values
func NewGetAllRegionsUsingGETForbidden() *GetAllRegionsUsingGETForbidden {
	return &GetAllRegionsUsingGETForbidden{}
}

/*GetAllRegionsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllRegionsUsingGETForbidden struct {
}

func (o *GetAllRegionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions][%d] getAllRegionsUsingGETForbidden ", 403)
}

func (o *GetAllRegionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllRegionsUsingGETNotFound creates a GetAllRegionsUsingGETNotFound with default headers values
func NewGetAllRegionsUsingGETNotFound() *GetAllRegionsUsingGETNotFound {
	return &GetAllRegionsUsingGETNotFound{}
}

/*GetAllRegionsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllRegionsUsingGETNotFound struct {
}

func (o *GetAllRegionsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions][%d] getAllRegionsUsingGETNotFound ", 404)
}

func (o *GetAllRegionsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}