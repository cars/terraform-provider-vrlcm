// Code generated by go-swagger; DO NOT EDIT.

package configuration_parameter_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAllCofigsUsingGETReader is a Reader for the GetAllCofigsUsingGET structure.
type GetAllCofigsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllCofigsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllCofigsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllCofigsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllCofigsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllCofigsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllCofigsUsingGETOK creates a GetAllCofigsUsingGETOK with default headers values
func NewGetAllCofigsUsingGETOK() *GetAllCofigsUsingGETOK {
	return &GetAllCofigsUsingGETOK{}
}

/*GetAllCofigsUsingGETOK handles this case with default header values.

OK
*/
type GetAllCofigsUsingGETOK struct {
	Payload interface{}
}

func (o *GetAllCofigsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/configparams][%d] getAllCofigsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllCofigsUsingGETOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAllCofigsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllCofigsUsingGETUnauthorized creates a GetAllCofigsUsingGETUnauthorized with default headers values
func NewGetAllCofigsUsingGETUnauthorized() *GetAllCofigsUsingGETUnauthorized {
	return &GetAllCofigsUsingGETUnauthorized{}
}

/*GetAllCofigsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllCofigsUsingGETUnauthorized struct {
}

func (o *GetAllCofigsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/configparams][%d] getAllCofigsUsingGETUnauthorized ", 401)
}

func (o *GetAllCofigsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCofigsUsingGETForbidden creates a GetAllCofigsUsingGETForbidden with default headers values
func NewGetAllCofigsUsingGETForbidden() *GetAllCofigsUsingGETForbidden {
	return &GetAllCofigsUsingGETForbidden{}
}

/*GetAllCofigsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllCofigsUsingGETForbidden struct {
}

func (o *GetAllCofigsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/configparams][%d] getAllCofigsUsingGETForbidden ", 403)
}

func (o *GetAllCofigsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllCofigsUsingGETNotFound creates a GetAllCofigsUsingGETNotFound with default headers values
func NewGetAllCofigsUsingGETNotFound() *GetAllCofigsUsingGETNotFound {
	return &GetAllCofigsUsingGETNotFound{}
}

/*GetAllCofigsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllCofigsUsingGETNotFound struct {
}

func (o *GetAllCofigsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/authzn/api/configparams][%d] getAllCofigsUsingGETNotFound ", 404)
}

func (o *GetAllCofigsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}