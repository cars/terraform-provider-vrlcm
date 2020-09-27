// Code generated by go-swagger; DO NOT EDIT.

package authentication_group_a_p_is

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

// NewGetGroupByIDUsingGETParams creates a new GetGroupByIDUsingGETParams object
// with the default values initialized.
func NewGetGroupByIDUsingGETParams() *GetGroupByIDUsingGETParams {
	var ()
	return &GetGroupByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupByIDUsingGETParamsWithTimeout creates a new GetGroupByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetGroupByIDUsingGETParams {
	var ()
	return &GetGroupByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetGroupByIDUsingGETParamsWithContext creates a new GetGroupByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupByIDUsingGETParamsWithContext(ctx context.Context) *GetGroupByIDUsingGETParams {
	var ()
	return &GetGroupByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetGroupByIDUsingGETParamsWithHTTPClient creates a new GetGroupByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetGroupByIDUsingGETParams {
	var ()
	return &GetGroupByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetGroupByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get group by Id using g e t operation typically these are written to a http.Request
*/
type GetGroupByIDUsingGETParams struct {

	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetGroupByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) WithContext(ctx context.Context) *GetGroupByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetGroupByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVmid adds the vmid to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) WithVmid(vmid string) *GetGroupByIDUsingGETParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the get group by Id using g e t params
func (o *GetGroupByIDUsingGETParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vmid
	if err := r.SetPathParam("vmid", o.Vmid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}