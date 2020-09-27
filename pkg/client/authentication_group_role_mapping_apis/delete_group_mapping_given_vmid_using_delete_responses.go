// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_role_mapping_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteGroupMappingGivenVmidUsingDELETEReader is a Reader for the DeleteGroupMappingGivenVmidUsingDELETE structure.
type DeleteGroupMappingGivenVmidUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupMappingGivenVmidUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteGroupMappingGivenVmidUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteGroupMappingGivenVmidUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteGroupMappingGivenVmidUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteGroupMappingGivenVmidUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteGroupMappingGivenVmidUsingDELETEOK creates a DeleteGroupMappingGivenVmidUsingDELETEOK with default headers values
func NewDeleteGroupMappingGivenVmidUsingDELETEOK() *DeleteGroupMappingGivenVmidUsingDELETEOK {
	return &DeleteGroupMappingGivenVmidUsingDELETEOK{}
}

/*DeleteGroupMappingGivenVmidUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteGroupMappingGivenVmidUsingDELETEOK struct {
	Payload interface{}
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/grouprolemapping/{vmid}][%d] deleteGroupMappingGivenVmidUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteGroupMappingGivenVmidUsingDELETENoContent creates a DeleteGroupMappingGivenVmidUsingDELETENoContent with default headers values
func NewDeleteGroupMappingGivenVmidUsingDELETENoContent() *DeleteGroupMappingGivenVmidUsingDELETENoContent {
	return &DeleteGroupMappingGivenVmidUsingDELETENoContent{}
}

/*DeleteGroupMappingGivenVmidUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteGroupMappingGivenVmidUsingDELETENoContent struct {
}

func (o *DeleteGroupMappingGivenVmidUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/grouprolemapping/{vmid}][%d] deleteGroupMappingGivenVmidUsingDELETENoContent ", 204)
}

func (o *DeleteGroupMappingGivenVmidUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupMappingGivenVmidUsingDELETEUnauthorized creates a DeleteGroupMappingGivenVmidUsingDELETEUnauthorized with default headers values
func NewDeleteGroupMappingGivenVmidUsingDELETEUnauthorized() *DeleteGroupMappingGivenVmidUsingDELETEUnauthorized {
	return &DeleteGroupMappingGivenVmidUsingDELETEUnauthorized{}
}

/*DeleteGroupMappingGivenVmidUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteGroupMappingGivenVmidUsingDELETEUnauthorized struct {
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/grouprolemapping/{vmid}][%d] deleteGroupMappingGivenVmidUsingDELETEUnauthorized ", 401)
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupMappingGivenVmidUsingDELETEForbidden creates a DeleteGroupMappingGivenVmidUsingDELETEForbidden with default headers values
func NewDeleteGroupMappingGivenVmidUsingDELETEForbidden() *DeleteGroupMappingGivenVmidUsingDELETEForbidden {
	return &DeleteGroupMappingGivenVmidUsingDELETEForbidden{}
}

/*DeleteGroupMappingGivenVmidUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteGroupMappingGivenVmidUsingDELETEForbidden struct {
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/grouprolemapping/{vmid}][%d] deleteGroupMappingGivenVmidUsingDELETEForbidden ", 403)
}

func (o *DeleteGroupMappingGivenVmidUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}