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
)

// NewGetDataCenterUsingGETParams creates a new GetDataCenterUsingGETParams object
// with the default values initialized.
func NewGetDataCenterUsingGETParams() *GetDataCenterUsingGETParams {
	var ()
	return &GetDataCenterUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDataCenterUsingGETParamsWithTimeout creates a new GetDataCenterUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDataCenterUsingGETParamsWithTimeout(timeout time.Duration) *GetDataCenterUsingGETParams {
	var ()
	return &GetDataCenterUsingGETParams{

		timeout: timeout,
	}
}

// NewGetDataCenterUsingGETParamsWithContext creates a new GetDataCenterUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDataCenterUsingGETParamsWithContext(ctx context.Context) *GetDataCenterUsingGETParams {
	var ()
	return &GetDataCenterUsingGETParams{

		Context: ctx,
	}
}

// NewGetDataCenterUsingGETParamsWithHTTPClient creates a new GetDataCenterUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDataCenterUsingGETParamsWithHTTPClient(client *http.Client) *GetDataCenterUsingGETParams {
	var ()
	return &GetDataCenterUsingGETParams{
		HTTPClient: client,
	}
}

/*GetDataCenterUsingGETParams contains all the parameters to send to the API endpoint
for the get data center using g e t operation typically these are written to a http.Request
*/
type GetDataCenterUsingGETParams struct {

	/*DataCenterName
	  dataCenterName

	*/
	DataCenterName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) WithTimeout(timeout time.Duration) *GetDataCenterUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) WithContext(ctx context.Context) *GetDataCenterUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) WithHTTPClient(client *http.Client) *GetDataCenterUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataCenterName adds the dataCenterName to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) WithDataCenterName(dataCenterName string) *GetDataCenterUsingGETParams {
	o.SetDataCenterName(dataCenterName)
	return o
}

// SetDataCenterName adds the dataCenterName to the get data center using g e t params
func (o *GetDataCenterUsingGETParams) SetDataCenterName(dataCenterName string) {
	o.DataCenterName = dataCenterName
}

// WriteToRequest writes these params to a swagger request
func (o *GetDataCenterUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dataCenterName
	if err := r.SetPathParam("dataCenterName", o.DataCenterName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
