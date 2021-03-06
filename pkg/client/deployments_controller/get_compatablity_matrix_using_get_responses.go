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

// GetCompatablityMatrixUsingGETReader is a Reader for the GetCompatablityMatrixUsingGET structure.
type GetCompatablityMatrixUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompatablityMatrixUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCompatablityMatrixUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCompatablityMatrixUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCompatablityMatrixUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCompatablityMatrixUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCompatablityMatrixUsingGETOK creates a GetCompatablityMatrixUsingGETOK with default headers values
func NewGetCompatablityMatrixUsingGETOK() *GetCompatablityMatrixUsingGETOK {
	return &GetCompatablityMatrixUsingGETOK{}
}

/*GetCompatablityMatrixUsingGETOK handles this case with default header values.

OK
*/
type GetCompatablityMatrixUsingGETOK struct {
	Payload []*models.InteropReleaseElement
}

func (o *GetCompatablityMatrixUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/products/{product}/version/{version}/interop][%d] getCompatablityMatrixUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetCompatablityMatrixUsingGETOK) GetPayload() []*models.InteropReleaseElement {
	return o.Payload
}

func (o *GetCompatablityMatrixUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompatablityMatrixUsingGETUnauthorized creates a GetCompatablityMatrixUsingGETUnauthorized with default headers values
func NewGetCompatablityMatrixUsingGETUnauthorized() *GetCompatablityMatrixUsingGETUnauthorized {
	return &GetCompatablityMatrixUsingGETUnauthorized{}
}

/*GetCompatablityMatrixUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCompatablityMatrixUsingGETUnauthorized struct {
}

func (o *GetCompatablityMatrixUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/products/{product}/version/{version}/interop][%d] getCompatablityMatrixUsingGETUnauthorized ", 401)
}

func (o *GetCompatablityMatrixUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCompatablityMatrixUsingGETForbidden creates a GetCompatablityMatrixUsingGETForbidden with default headers values
func NewGetCompatablityMatrixUsingGETForbidden() *GetCompatablityMatrixUsingGETForbidden {
	return &GetCompatablityMatrixUsingGETForbidden{}
}

/*GetCompatablityMatrixUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetCompatablityMatrixUsingGETForbidden struct {
}

func (o *GetCompatablityMatrixUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/products/{product}/version/{version}/interop][%d] getCompatablityMatrixUsingGETForbidden ", 403)
}

func (o *GetCompatablityMatrixUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCompatablityMatrixUsingGETNotFound creates a GetCompatablityMatrixUsingGETNotFound with default headers values
func NewGetCompatablityMatrixUsingGETNotFound() *GetCompatablityMatrixUsingGETNotFound {
	return &GetCompatablityMatrixUsingGETNotFound{}
}

/*GetCompatablityMatrixUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetCompatablityMatrixUsingGETNotFound struct {
}

func (o *GetCompatablityMatrixUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/lcops/api/products/{product}/version/{version}/interop][%d] getCompatablityMatrixUsingGETNotFound ", 404)
}

func (o *GetCompatablityMatrixUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
