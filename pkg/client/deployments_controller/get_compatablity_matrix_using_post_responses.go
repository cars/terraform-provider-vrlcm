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

// GetCompatablityMatrixUsingPOSTReader is a Reader for the GetCompatablityMatrixUsingPOST structure.
type GetCompatablityMatrixUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCompatablityMatrixUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCompatablityMatrixUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewGetCompatablityMatrixUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCompatablityMatrixUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCompatablityMatrixUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetCompatablityMatrixUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCompatablityMatrixUsingPOSTOK creates a GetCompatablityMatrixUsingPOSTOK with default headers values
func NewGetCompatablityMatrixUsingPOSTOK() *GetCompatablityMatrixUsingPOSTOK {
	return &GetCompatablityMatrixUsingPOSTOK{}
}

/*GetCompatablityMatrixUsingPOSTOK handles this case with default header values.

OK
*/
type GetCompatablityMatrixUsingPOSTOK struct {
	Payload map[string]models.InteropResponseDTO
}

func (o *GetCompatablityMatrixUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/products/interop][%d] getCompatablityMatrixUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *GetCompatablityMatrixUsingPOSTOK) GetPayload() map[string]models.InteropResponseDTO {
	return o.Payload
}

func (o *GetCompatablityMatrixUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCompatablityMatrixUsingPOSTCreated creates a GetCompatablityMatrixUsingPOSTCreated with default headers values
func NewGetCompatablityMatrixUsingPOSTCreated() *GetCompatablityMatrixUsingPOSTCreated {
	return &GetCompatablityMatrixUsingPOSTCreated{}
}

/*GetCompatablityMatrixUsingPOSTCreated handles this case with default header values.

Created
*/
type GetCompatablityMatrixUsingPOSTCreated struct {
}

func (o *GetCompatablityMatrixUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/products/interop][%d] getCompatablityMatrixUsingPOSTCreated ", 201)
}

func (o *GetCompatablityMatrixUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCompatablityMatrixUsingPOSTUnauthorized creates a GetCompatablityMatrixUsingPOSTUnauthorized with default headers values
func NewGetCompatablityMatrixUsingPOSTUnauthorized() *GetCompatablityMatrixUsingPOSTUnauthorized {
	return &GetCompatablityMatrixUsingPOSTUnauthorized{}
}

/*GetCompatablityMatrixUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCompatablityMatrixUsingPOSTUnauthorized struct {
}

func (o *GetCompatablityMatrixUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/products/interop][%d] getCompatablityMatrixUsingPOSTUnauthorized ", 401)
}

func (o *GetCompatablityMatrixUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCompatablityMatrixUsingPOSTForbidden creates a GetCompatablityMatrixUsingPOSTForbidden with default headers values
func NewGetCompatablityMatrixUsingPOSTForbidden() *GetCompatablityMatrixUsingPOSTForbidden {
	return &GetCompatablityMatrixUsingPOSTForbidden{}
}

/*GetCompatablityMatrixUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type GetCompatablityMatrixUsingPOSTForbidden struct {
}

func (o *GetCompatablityMatrixUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/products/interop][%d] getCompatablityMatrixUsingPOSTForbidden ", 403)
}

func (o *GetCompatablityMatrixUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCompatablityMatrixUsingPOSTNotFound creates a GetCompatablityMatrixUsingPOSTNotFound with default headers values
func NewGetCompatablityMatrixUsingPOSTNotFound() *GetCompatablityMatrixUsingPOSTNotFound {
	return &GetCompatablityMatrixUsingPOSTNotFound{}
}

/*GetCompatablityMatrixUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type GetCompatablityMatrixUsingPOSTNotFound struct {
}

func (o *GetCompatablityMatrixUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /lcm/lcops/api/products/interop][%d] getCompatablityMatrixUsingPOSTNotFound ", 404)
}

func (o *GetCompatablityMatrixUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
