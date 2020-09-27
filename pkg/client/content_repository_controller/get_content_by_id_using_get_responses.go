// Code generated by go-swagger; DO NOT EDIT.

package content_repository_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetContentByIDUsingGETReader is a Reader for the GetContentByIDUsingGET structure.
type GetContentByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetContentByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentByIDUsingGETOK creates a GetContentByIDUsingGETOK with default headers values
func NewGetContentByIDUsingGETOK() *GetContentByIDUsingGETOK {
	return &GetContentByIDUsingGETOK{}
}

/*GetContentByIDUsingGETOK handles this case with default header values.

OK
*/
type GetContentByIDUsingGETOK struct {
	Payload *models.ContentDTO
}

func (o *GetContentByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/{vmid}][%d] getContentByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetContentByIDUsingGETOK) GetPayload() *models.ContentDTO {
	return o.Payload
}

func (o *GetContentByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentByIDUsingGETUnauthorized creates a GetContentByIDUsingGETUnauthorized with default headers values
func NewGetContentByIDUsingGETUnauthorized() *GetContentByIDUsingGETUnauthorized {
	return &GetContentByIDUsingGETUnauthorized{}
}

/*GetContentByIDUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetContentByIDUsingGETUnauthorized struct {
}

func (o *GetContentByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/{vmid}][%d] getContentByIdUsingGETUnauthorized ", 401)
}

func (o *GetContentByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContentByIDUsingGETForbidden creates a GetContentByIDUsingGETForbidden with default headers values
func NewGetContentByIDUsingGETForbidden() *GetContentByIDUsingGETForbidden {
	return &GetContentByIDUsingGETForbidden{}
}

/*GetContentByIDUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetContentByIDUsingGETForbidden struct {
}

func (o *GetContentByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/{vmid}][%d] getContentByIdUsingGETForbidden ", 403)
}

func (o *GetContentByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetContentByIDUsingGETNotFound creates a GetContentByIDUsingGETNotFound with default headers values
func NewGetContentByIDUsingGETNotFound() *GetContentByIDUsingGETNotFound {
	return &GetContentByIDUsingGETNotFound{}
}

/*GetContentByIDUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetContentByIDUsingGETNotFound struct {
}

func (o *GetContentByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/{vmid}][%d] getContentByIdUsingGETNotFound ", 404)
}

func (o *GetContentByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
