// Code generated by go-swagger; DO NOT EDIT.

package i_market_place_app_metadata_controller_impl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// GetAllMarketPlaceAppsUsingGETReader is a Reader for the GetAllMarketPlaceAppsUsingGET structure.
type GetAllMarketPlaceAppsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllMarketPlaceAppsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllMarketPlaceAppsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllMarketPlaceAppsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllMarketPlaceAppsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllMarketPlaceAppsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllMarketPlaceAppsUsingGETOK creates a GetAllMarketPlaceAppsUsingGETOK with default headers values
func NewGetAllMarketPlaceAppsUsingGETOK() *GetAllMarketPlaceAppsUsingGETOK {
	return &GetAllMarketPlaceAppsUsingGETOK{}
}

/*GetAllMarketPlaceAppsUsingGETOK handles this case with default header values.

OK
*/
type GetAllMarketPlaceAppsUsingGETOK struct {
	Payload []*models.MarketPlaceAppDTO
}

func (o *GetAllMarketPlaceAppsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /lcm/marketplace/api/apps][%d] getAllMarketPlaceAppsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllMarketPlaceAppsUsingGETOK) GetPayload() []*models.MarketPlaceAppDTO {
	return o.Payload
}

func (o *GetAllMarketPlaceAppsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllMarketPlaceAppsUsingGETUnauthorized creates a GetAllMarketPlaceAppsUsingGETUnauthorized with default headers values
func NewGetAllMarketPlaceAppsUsingGETUnauthorized() *GetAllMarketPlaceAppsUsingGETUnauthorized {
	return &GetAllMarketPlaceAppsUsingGETUnauthorized{}
}

/*GetAllMarketPlaceAppsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllMarketPlaceAppsUsingGETUnauthorized struct {
}

func (o *GetAllMarketPlaceAppsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /lcm/marketplace/api/apps][%d] getAllMarketPlaceAppsUsingGETUnauthorized ", 401)
}

func (o *GetAllMarketPlaceAppsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllMarketPlaceAppsUsingGETForbidden creates a GetAllMarketPlaceAppsUsingGETForbidden with default headers values
func NewGetAllMarketPlaceAppsUsingGETForbidden() *GetAllMarketPlaceAppsUsingGETForbidden {
	return &GetAllMarketPlaceAppsUsingGETForbidden{}
}

/*GetAllMarketPlaceAppsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllMarketPlaceAppsUsingGETForbidden struct {
}

func (o *GetAllMarketPlaceAppsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /lcm/marketplace/api/apps][%d] getAllMarketPlaceAppsUsingGETForbidden ", 403)
}

func (o *GetAllMarketPlaceAppsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllMarketPlaceAppsUsingGETNotFound creates a GetAllMarketPlaceAppsUsingGETNotFound with default headers values
func NewGetAllMarketPlaceAppsUsingGETNotFound() *GetAllMarketPlaceAppsUsingGETNotFound {
	return &GetAllMarketPlaceAppsUsingGETNotFound{}
}

/*GetAllMarketPlaceAppsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllMarketPlaceAppsUsingGETNotFound struct {
}

func (o *GetAllMarketPlaceAppsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /lcm/marketplace/api/apps][%d] getAllMarketPlaceAppsUsingGETNotFound ", 404)
}

func (o *GetAllMarketPlaceAppsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
