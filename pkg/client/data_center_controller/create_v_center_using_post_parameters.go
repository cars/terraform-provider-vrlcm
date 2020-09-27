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

// NewCreateVCenterUsingPOSTParams creates a new CreateVCenterUsingPOSTParams object
// with the default values initialized.
func NewCreateVCenterUsingPOSTParams() *CreateVCenterUsingPOSTParams {
	var ()
	return &CreateVCenterUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVCenterUsingPOSTParamsWithTimeout creates a new CreateVCenterUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVCenterUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateVCenterUsingPOSTParams {
	var ()
	return &CreateVCenterUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateVCenterUsingPOSTParamsWithContext creates a new CreateVCenterUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVCenterUsingPOSTParamsWithContext(ctx context.Context) *CreateVCenterUsingPOSTParams {
	var ()
	return &CreateVCenterUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateVCenterUsingPOSTParamsWithHTTPClient creates a new CreateVCenterUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVCenterUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateVCenterUsingPOSTParams {
	var ()
	return &CreateVCenterUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateVCenterUsingPOSTParams contains all the parameters to send to the API endpoint
for the create v center using p o s t operation typically these are written to a http.Request
*/
type CreateVCenterUsingPOSTParams struct {

	/*DataCenterName
	  dataCenterName

	*/
	DataCenterName string
	/*RegionName
	  regionName

	*/
	RegionName string
	/*VCenterRequestDTO
	  vCenterRequestDTO

	*/
	VCenterRequestDTO *models.VCenterRequestDTO
	/*ZoneName
	  zoneName

	*/
	ZoneName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateVCenterUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithContext(ctx context.Context) *CreateVCenterUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateVCenterUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataCenterName adds the dataCenterName to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithDataCenterName(dataCenterName string) *CreateVCenterUsingPOSTParams {
	o.SetDataCenterName(dataCenterName)
	return o
}

// SetDataCenterName adds the dataCenterName to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetDataCenterName(dataCenterName string) {
	o.DataCenterName = dataCenterName
}

// WithRegionName adds the regionName to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithRegionName(regionName string) *CreateVCenterUsingPOSTParams {
	o.SetRegionName(regionName)
	return o
}

// SetRegionName adds the regionName to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetRegionName(regionName string) {
	o.RegionName = regionName
}

// WithVCenterRequestDTO adds the vCenterRequestDTO to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithVCenterRequestDTO(vCenterRequestDTO *models.VCenterRequestDTO) *CreateVCenterUsingPOSTParams {
	o.SetVCenterRequestDTO(vCenterRequestDTO)
	return o
}

// SetVCenterRequestDTO adds the vCenterRequestDTO to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetVCenterRequestDTO(vCenterRequestDTO *models.VCenterRequestDTO) {
	o.VCenterRequestDTO = vCenterRequestDTO
}

// WithZoneName adds the zoneName to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) WithZoneName(zoneName string) *CreateVCenterUsingPOSTParams {
	o.SetZoneName(zoneName)
	return o
}

// SetZoneName adds the zoneName to the create v center using p o s t params
func (o *CreateVCenterUsingPOSTParams) SetZoneName(zoneName string) {
	o.ZoneName = zoneName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVCenterUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.VCenterRequestDTO != nil {
		if err := r.SetBodyParam(o.VCenterRequestDTO); err != nil {
			return err
		}
	}

	// path param zoneName
	if err := r.SetPathParam("zoneName", o.ZoneName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
