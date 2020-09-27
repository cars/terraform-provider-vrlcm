// Code generated by go-swagger; DO NOT EDIT.

package data_center_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateDataCenterUsingPUTReader is a Reader for the UpdateDataCenterUsingPUT structure.
type UpdateDataCenterUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDataCenterUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDataCenterUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUpdateDataCenterUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateDataCenterUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDataCenterUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDataCenterUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDataCenterUsingPUTOK creates a UpdateDataCenterUsingPUTOK with default headers values
func NewUpdateDataCenterUsingPUTOK() *UpdateDataCenterUsingPUTOK {
	return &UpdateDataCenterUsingPUTOK{}
}

/*UpdateDataCenterUsingPUTOK handles this case with default header values.

OK
*/
type UpdateDataCenterUsingPUTOK struct {
}

func (o *UpdateDataCenterUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/datacenters/{dataCenterName}][%d] updateDataCenterUsingPUTOK ", 200)
}

func (o *UpdateDataCenterUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDataCenterUsingPUTCreated creates a UpdateDataCenterUsingPUTCreated with default headers values
func NewUpdateDataCenterUsingPUTCreated() *UpdateDataCenterUsingPUTCreated {
	return &UpdateDataCenterUsingPUTCreated{}
}

/*UpdateDataCenterUsingPUTCreated handles this case with default header values.

Created
*/
type UpdateDataCenterUsingPUTCreated struct {
}

func (o *UpdateDataCenterUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/datacenters/{dataCenterName}][%d] updateDataCenterUsingPUTCreated ", 201)
}

func (o *UpdateDataCenterUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDataCenterUsingPUTUnauthorized creates a UpdateDataCenterUsingPUTUnauthorized with default headers values
func NewUpdateDataCenterUsingPUTUnauthorized() *UpdateDataCenterUsingPUTUnauthorized {
	return &UpdateDataCenterUsingPUTUnauthorized{}
}

/*UpdateDataCenterUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateDataCenterUsingPUTUnauthorized struct {
}

func (o *UpdateDataCenterUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/datacenters/{dataCenterName}][%d] updateDataCenterUsingPUTUnauthorized ", 401)
}

func (o *UpdateDataCenterUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDataCenterUsingPUTForbidden creates a UpdateDataCenterUsingPUTForbidden with default headers values
func NewUpdateDataCenterUsingPUTForbidden() *UpdateDataCenterUsingPUTForbidden {
	return &UpdateDataCenterUsingPUTForbidden{}
}

/*UpdateDataCenterUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type UpdateDataCenterUsingPUTForbidden struct {
}

func (o *UpdateDataCenterUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/datacenters/{dataCenterName}][%d] updateDataCenterUsingPUTForbidden ", 403)
}

func (o *UpdateDataCenterUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDataCenterUsingPUTNotFound creates a UpdateDataCenterUsingPUTNotFound with default headers values
func NewUpdateDataCenterUsingPUTNotFound() *UpdateDataCenterUsingPUTNotFound {
	return &UpdateDataCenterUsingPUTNotFound{}
}

/*UpdateDataCenterUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdateDataCenterUsingPUTNotFound struct {
}

func (o *UpdateDataCenterUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /lcm/lcops/api/datacenters/{dataCenterName}][%d] updateDataCenterUsingPUTNotFound ", 404)
}

func (o *UpdateDataCenterUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
