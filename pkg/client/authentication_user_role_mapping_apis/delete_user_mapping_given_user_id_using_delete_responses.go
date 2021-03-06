// Code generated by go-swagger; DO NOT EDIT.

package authentication_user_role_mapping_a_p_is

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteUserMappingGivenUserIDUsingDELETEReader is a Reader for the DeleteUserMappingGivenUserIDUsingDELETE structure.
type DeleteUserMappingGivenUserIDUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserMappingGivenUserIDUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserMappingGivenUserIDUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteUserMappingGivenUserIDUsingDELETENoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUserMappingGivenUserIDUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserMappingGivenUserIDUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserMappingGivenUserIDUsingDELETEOK creates a DeleteUserMappingGivenUserIDUsingDELETEOK with default headers values
func NewDeleteUserMappingGivenUserIDUsingDELETEOK() *DeleteUserMappingGivenUserIDUsingDELETEOK {
	return &DeleteUserMappingGivenUserIDUsingDELETEOK{}
}

/*DeleteUserMappingGivenUserIDUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteUserMappingGivenUserIDUsingDELETEOK struct {
	Payload interface{}
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/userrolemapping/userId/{uservmid}][%d] deleteUserMappingGivenUserIdUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserMappingGivenUserIDUsingDELETENoContent creates a DeleteUserMappingGivenUserIDUsingDELETENoContent with default headers values
func NewDeleteUserMappingGivenUserIDUsingDELETENoContent() *DeleteUserMappingGivenUserIDUsingDELETENoContent {
	return &DeleteUserMappingGivenUserIDUsingDELETENoContent{}
}

/*DeleteUserMappingGivenUserIDUsingDELETENoContent handles this case with default header values.

No Content
*/
type DeleteUserMappingGivenUserIDUsingDELETENoContent struct {
}

func (o *DeleteUserMappingGivenUserIDUsingDELETENoContent) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/userrolemapping/userId/{uservmid}][%d] deleteUserMappingGivenUserIdUsingDELETENoContent ", 204)
}

func (o *DeleteUserMappingGivenUserIDUsingDELETENoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserMappingGivenUserIDUsingDELETEUnauthorized creates a DeleteUserMappingGivenUserIDUsingDELETEUnauthorized with default headers values
func NewDeleteUserMappingGivenUserIDUsingDELETEUnauthorized() *DeleteUserMappingGivenUserIDUsingDELETEUnauthorized {
	return &DeleteUserMappingGivenUserIDUsingDELETEUnauthorized{}
}

/*DeleteUserMappingGivenUserIDUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteUserMappingGivenUserIDUsingDELETEUnauthorized struct {
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/userrolemapping/userId/{uservmid}][%d] deleteUserMappingGivenUserIdUsingDELETEUnauthorized ", 401)
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserMappingGivenUserIDUsingDELETEForbidden creates a DeleteUserMappingGivenUserIDUsingDELETEForbidden with default headers values
func NewDeleteUserMappingGivenUserIDUsingDELETEForbidden() *DeleteUserMappingGivenUserIDUsingDELETEForbidden {
	return &DeleteUserMappingGivenUserIDUsingDELETEForbidden{}
}

/*DeleteUserMappingGivenUserIDUsingDELETEForbidden handles this case with default header values.

Forbidden
*/
type DeleteUserMappingGivenUserIDUsingDELETEForbidden struct {
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /lcm/authzn/api/userrolemapping/userId/{uservmid}][%d] deleteUserMappingGivenUserIdUsingDELETEForbidden ", 403)
}

func (o *DeleteUserMappingGivenUserIDUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
