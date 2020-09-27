// Code generated by go-swagger; DO NOT EDIT.

package data_center_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRegionUsingGETReader is a Reader for the GetRegionUsingGET structure.
type GetRegionUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegionUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegionUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegionUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegionUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRegionUsingGETOK creates a GetRegionUsingGETOK with default headers values
func NewGetRegionUsingGETOK() *GetRegionUsingGETOK {
	return &GetRegionUsingGETOK{}
}

/*GetRegionUsingGETOK handles this case with default header values.

OK
*/
type GetRegionUsingGETOK struct {
}

func (o *GetRegionUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}][%d] getRegionUsingGETOK ", 200)
}

func (o *GetRegionUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegionUsingGETUnauthorized creates a GetRegionUsingGETUnauthorized with default headers values
func NewGetRegionUsingGETUnauthorized() *GetRegionUsingGETUnauthorized {
	return &GetRegionUsingGETUnauthorized{}
}

/*GetRegionUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRegionUsingGETUnauthorized struct {
}

func (o *GetRegionUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}][%d] getRegionUsingGETUnauthorized ", 401)
}

func (o *GetRegionUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegionUsingGETForbidden creates a GetRegionUsingGETForbidden with default headers values
func NewGetRegionUsingGETForbidden() *GetRegionUsingGETForbidden {
	return &GetRegionUsingGETForbidden{}
}

/*GetRegionUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetRegionUsingGETForbidden struct {
}

func (o *GetRegionUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}][%d] getRegionUsingGETForbidden ", 403)
}

func (o *GetRegionUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegionUsingGETNotFound creates a GetRegionUsingGETNotFound with default headers values
func NewGetRegionUsingGETNotFound() *GetRegionUsingGETNotFound {
	return &GetRegionUsingGETNotFound{}
}

/*GetRegionUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetRegionUsingGETNotFound struct {
}

func (o *GetRegionUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/datacenters/{dataCenterName}/regions/{regionName}][%d] getRegionUsingGETNotFound ", 404)
}

func (o *GetRegionUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}