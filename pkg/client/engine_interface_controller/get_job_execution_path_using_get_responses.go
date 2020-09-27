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

// GetJobExecutionPathUsingGETReader is a Reader for the GetJobExecutionPathUsingGET structure.
type GetJobExecutionPathUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobExecutionPathUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJobExecutionPathUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetJobExecutionPathUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJobExecutionPathUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJobExecutionPathUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJobExecutionPathUsingGETOK creates a GetJobExecutionPathUsingGETOK with default headers values
func NewGetJobExecutionPathUsingGETOK() *GetJobExecutionPathUsingGETOK {
	return &GetJobExecutionPathUsingGETOK{}
}

/*GetJobExecutionPathUsingGETOK handles this case with default header values.

OK
*/
type GetJobExecutionPathUsingGETOK struct {
	Payload *models.StatusDetailedSuiteSpec
}

func (o *GetJobExecutionPathUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/executionpath/{vmid}][%d] getJobExecutionPathUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetJobExecutionPathUsingGETOK) GetPayload() *models.StatusDetailedSuiteSpec {
	return o.Payload
}

func (o *GetJobExecutionPathUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StatusDetailedSuiteSpec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJobExecutionPathUsingGETUnauthorized creates a GetJobExecutionPathUsingGETUnauthorized with default headers values
func NewGetJobExecutionPathUsingGETUnauthorized() *GetJobExecutionPathUsingGETUnauthorized {
	return &GetJobExecutionPathUsingGETUnauthorized{}
}

/*GetJobExecutionPathUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetJobExecutionPathUsingGETUnauthorized struct {
}

func (o *GetJobExecutionPathUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/executionpath/{vmid}][%d] getJobExecutionPathUsingGETUnauthorized ", 401)
}

func (o *GetJobExecutionPathUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJobExecutionPathUsingGETForbidden creates a GetJobExecutionPathUsingGETForbidden with default headers values
func NewGetJobExecutionPathUsingGETForbidden() *GetJobExecutionPathUsingGETForbidden {
	return &GetJobExecutionPathUsingGETForbidden{}
}

/*GetJobExecutionPathUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetJobExecutionPathUsingGETForbidden struct {
}

func (o *GetJobExecutionPathUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/executionpath/{vmid}][%d] getJobExecutionPathUsingGETForbidden ", 403)
}

func (o *GetJobExecutionPathUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJobExecutionPathUsingGETNotFound creates a GetJobExecutionPathUsingGETNotFound with default headers values
func NewGetJobExecutionPathUsingGETNotFound() *GetJobExecutionPathUsingGETNotFound {
	return &GetJobExecutionPathUsingGETNotFound{}
}

/*GetJobExecutionPathUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetJobExecutionPathUsingGETNotFound struct {
}

func (o *GetJobExecutionPathUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/automata/api/engine/statemachine/executionpath/{vmid}][%d] getJobExecutionPathUsingGETNotFound ", 404)
}

func (o *GetJobExecutionPathUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
