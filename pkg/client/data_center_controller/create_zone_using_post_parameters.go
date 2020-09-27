// Code generated by go-swagger; DO NOT EDIT.

package data_center_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cars/terraform-provider-vrlcm/models"
)

// NewCreateZoneUsingPOSTParams creates a new CreateZoneUsingPOSTParams object
// with the default values initialized.
func NewCreateZoneUsingPOSTParams() *CreateZoneUsingPOSTParams {
	var ()
	return &CreateZoneUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateZoneUsingPOSTParamsWithTimeout creates a new CreateZoneUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateZoneUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateZoneUsingPOSTParams {
	var ()
	return &CreateZoneUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateZoneUsingPOSTParamsWithContext creates a new CreateZoneUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateZoneUsingPOSTParamsWithContext(ctx context.Context) *CreateZoneUsingPOSTParams {
	var ()
	return &CreateZoneUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateZoneUsingPOSTParamsWithHTTPClient creates a new CreateZoneUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateZoneUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateZoneUsingPOSTParams {
	var ()
	return &CreateZoneUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateZoneUsingPOSTParams contains all the parameters to send to the API endpoint
for the create zone using p o s t operation typically these are written to a http.Request
*/
type CreateZoneUsingPOSTParams struct {

	/*DataCenterName
	  dataCenterName

	*/
	DataCenterName string
	/*RegionName
	  regionName

	*/
	RegionName string
	/*ZoneRequestDTO
	  zoneRequestDTO

	*/
	ZoneRequestDTO *models.ZoneRequestDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateZoneUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) WithContext(ctx context.Context) *CreateZoneUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateZoneUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataCenterName adds the dataCenterName to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) WithDataCenterName(dataCenterName string) *CreateZoneUsingPOSTParams {
	o.SetDataCenterName(dataCenterName)
	return o
}

// SetDataCenterName adds the dataCenterName to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) SetDataCenterName(dataCenterName string) {
	o.DataCenterName = dataCenterName
}

// WithRegionName adds the regionName to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) WithRegionName(regionName string) *CreateZoneUsingPOSTParams {
	o.SetRegionName(regionName)
	return o
}

// SetRegionName adds the regionName to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) SetRegionName(regionName string) {
	o.RegionName = regionName
}

// WithZoneRequestDTO adds the zoneRequestDTO to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) WithZoneRequestDTO(zoneRequestDTO *models.ZoneRequestDTO) *CreateZoneUsingPOSTParams {
	o.SetZoneRequestDTO(zoneRequestDTO)
	return o
}

// SetZoneRequestDTO adds the zoneRequestDTO to the create zone using p o s t params
func (o *CreateZoneUsingPOSTParams) SetZoneRequestDTO(zoneRequestDTO *models.ZoneRequestDTO) {
	o.ZoneRequestDTO = zoneRequestDTO
}

// WriteToRequest writes these params to a swagger request
func (o *CreateZoneUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dataCenterName
	if err := r.SetPathParam("dataCenterName", o.DataCenterName); err != nil {
		return err
	}

	// path param regionName
	if err := r.SetPathParam("regionName", o.RegionName); err != nil {
		return err
	}

	if o.ZoneRequestDTO != nil {
		if err := r.SetBodyParam(o.ZoneRequestDTO); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}