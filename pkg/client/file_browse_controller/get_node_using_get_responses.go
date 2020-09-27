// Code generated by go-swagger; DO NOT EDIT.

package file_browse_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetNodeUsingGETReader is a Reader for the GetNodeUsingGET structure.
type GetNodeUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetNodeUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNodeUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNodeUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNodeUsingGETOK creates a GetNodeUsingGETOK with default headers values
func NewGetNodeUsingGETOK() *GetNodeUsingGETOK {
	return &GetNodeUsingGETOK{}
}

/*GetNodeUsingGETOK handles this case with default header values.

OK
*/
type GetNodeUsingGETOK struct {
	Payload *models.NodeDTO
}

func (o *GetNodeUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/nodes/{vmid}][%d] getNodeUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetNodeUsingGETOK) GetPayload() *models.NodeDTO {
	return o.Payload
}

func (o *GetNodeUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDTO)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeUsingGETUnauthorized creates a GetNodeUsingGETUnauthorized with default headers values
func NewGetNodeUsingGETUnauthorized() *GetNodeUsingGETUnauthorized {
	return &GetNodeUsingGETUnauthorized{}
}

/*GetNodeUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetNodeUsingGETUnauthorized struct {
}

func (o *GetNodeUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/nodes/{vmid}][%d] getNodeUsingGETUnauthorized ", 401)
}

func (o *GetNodeUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeUsingGETForbidden creates a GetNodeUsingGETForbidden with default headers values
func NewGetNodeUsingGETForbidden() *GetNodeUsingGETForbidden {
	return &GetNodeUsingGETForbidden{}
}

/*GetNodeUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetNodeUsingGETForbidden struct {
}

func (o *GetNodeUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/nodes/{vmid}][%d] getNodeUsingGETForbidden ", 403)
}

func (o *GetNodeUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeUsingGETNotFound creates a GetNodeUsingGETNotFound with default headers values
func NewGetNodeUsingGETNotFound() *GetNodeUsingGETNotFound {
	return &GetNodeUsingGETNotFound{}
}

/*GetNodeUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetNodeUsingGETNotFound struct {
}

func (o *GetNodeUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/crepo/api/content/nodes/{vmid}][%d] getNodeUsingGETNotFound ", 404)
}

func (o *GetNodeUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
