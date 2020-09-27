// Code generated by go-swagger; DO NOT EDIT.

package engine_interface_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetJobSpecificationUsingGETReader is a Reader for the GetJobSpecificationUsingGET structure.
type GetJobSpecificationUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobSpecificationUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobSpecificationUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetJobSpecificationUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJobSpecificationUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobSpecificationUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobSpecificationUsingGETOK creates a GetJobSpecificationUsingGETOK with default headers values
func NewGetJobSpecificationUsingGETOK() *GetJobSpecificationUsingGETOK {
	return &GetJobSpecificationUsingGETOK{}
}

/*GetJobSpecificationUsingGETOK handles this case with default header values.

OK
*/
type GetJobSpecificationUsingGETOK struct {
	Payload *models.StateMachineInvocationRequest
}

func (o *GetJobSpecificationUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/detailedspec/{vmid}][%d] getJobSpecificationUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetJobSpecificationUsingGETOK) GetPayload() *models.StateMachineInvocationRequest {
	return o.Payload
}

func (o *GetJobSpecificationUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StateMachineInvocationRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobSpecificationUsingGETUnauthorized creates a GetJobSpecificationUsingGETUnauthorized with default headers values
func NewGetJobSpecificationUsingGETUnauthorized() *GetJobSpecificationUsingGETUnauthorized {
	return &GetJobSpecificationUsingGETUnauthorized{}
}

/*GetJobSpecificationUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetJobSpecificationUsingGETUnauthorized struct {
}

func (o *GetJobSpecificationUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/detailedspec/{vmid}][%d] getJobSpecificationUsingGETUnauthorized ", 401)
}

func (o *GetJobSpecificationUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJobSpecificationUsingGETForbidden creates a GetJobSpecificationUsingGETForbidden with default headers values
func NewGetJobSpecificationUsingGETForbidden() *GetJobSpecificationUsingGETForbidden {
	return &GetJobSpecificationUsingGETForbidden{}
}

/*GetJobSpecificationUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetJobSpecificationUsingGETForbidden struct {
}

func (o *GetJobSpecificationUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/detailedspec/{vmid}][%d] getJobSpecificationUsingGETForbidden ", 403)
}

func (o *GetJobSpecificationUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJobSpecificationUsingGETNotFound creates a GetJobSpecificationUsingGETNotFound with default headers values
func NewGetJobSpecificationUsingGETNotFound() *GetJobSpecificationUsingGETNotFound {
	return &GetJobSpecificationUsingGETNotFound{}
}

/*GetJobSpecificationUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetJobSpecificationUsingGETNotFound struct {
}

func (o *GetJobSpecificationUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/detailedspec/{vmid}][%d] getJobSpecificationUsingGETNotFound ", 404)
}

func (o *GetJobSpecificationUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}