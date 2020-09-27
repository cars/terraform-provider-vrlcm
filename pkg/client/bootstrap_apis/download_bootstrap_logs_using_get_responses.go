// Code generated by go-swagger; DO NOT EDIT.

package bootstrap_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// DownloadBootstrapLogsUsingGETReader is a Reader for the DownloadBootstrapLogsUsingGET structure.
type DownloadBootstrapLogsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadBootstrapLogsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDownloadBootstrapLogsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDownloadBootstrapLogsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDownloadBootstrapLogsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDownloadBootstrapLogsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDownloadBootstrapLogsUsingGETOK creates a DownloadBootstrapLogsUsingGETOK with default headers values
func NewDownloadBootstrapLogsUsingGETOK() *DownloadBootstrapLogsUsingGETOK {
	return &DownloadBootstrapLogsUsingGETOK{}
}

/*DownloadBootstrapLogsUsingGETOK handles this case with default header values.

OK
*/
type DownloadBootstrapLogsUsingGETOK struct {
	Payload *models.Resource
}

func (o *DownloadBootstrapLogsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/log/download][%d] downloadBootstrapLogsUsingGETOK  %+v", 200, o.Payload)
}

func (o *DownloadBootstrapLogsUsingGETOK) GetPayload() *models.Resource {
	return o.Payload
}

func (o *DownloadBootstrapLogsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Resource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDownloadBootstrapLogsUsingGETUnauthorized creates a DownloadBootstrapLogsUsingGETUnauthorized with default headers values
func NewDownloadBootstrapLogsUsingGETUnauthorized() *DownloadBootstrapLogsUsingGETUnauthorized {
	return &DownloadBootstrapLogsUsingGETUnauthorized{}
}

/*DownloadBootstrapLogsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type DownloadBootstrapLogsUsingGETUnauthorized struct {
}

func (o *DownloadBootstrapLogsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/log/download][%d] downloadBootstrapLogsUsingGETUnauthorized ", 401)
}

func (o *DownloadBootstrapLogsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadBootstrapLogsUsingGETForbidden creates a DownloadBootstrapLogsUsingGETForbidden with default headers values
func NewDownloadBootstrapLogsUsingGETForbidden() *DownloadBootstrapLogsUsingGETForbidden {
	return &DownloadBootstrapLogsUsingGETForbidden{}
}

/*DownloadBootstrapLogsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type DownloadBootstrapLogsUsingGETForbidden struct {
}

func (o *DownloadBootstrapLogsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/log/download][%d] downloadBootstrapLogsUsingGETForbidden ", 403)
}

func (o *DownloadBootstrapLogsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDownloadBootstrapLogsUsingGETNotFound creates a DownloadBootstrapLogsUsingGETNotFound with default headers values
func NewDownloadBootstrapLogsUsingGETNotFound() *DownloadBootstrapLogsUsingGETNotFound {
	return &DownloadBootstrapLogsUsingGETNotFound{}
}

/*DownloadBootstrapLogsUsingGETNotFound handles this case with default header values.

Not Found
*/
type DownloadBootstrapLogsUsingGETNotFound struct {
}

func (o *DownloadBootstrapLogsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/bootstrap/api/log/download][%d] downloadBootstrapLogsUsingGETNotFound ", 404)
}

func (o *DownloadBootstrapLogsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}