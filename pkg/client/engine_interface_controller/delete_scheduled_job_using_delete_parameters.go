// Code generated by go-swagger; DO NOT EDIT.

package engine_interface_controller

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

// NewDeleteScheduledJobUsingDELETEParams creates a new DeleteScheduledJobUsingDELETEParams object
// with the default values initialized.
func NewDeleteScheduledJobUsingDELETEParams() *DeleteScheduledJobUsingDELETEParams {
	var ()
	return &DeleteScheduledJobUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScheduledJobUsingDELETEParamsWithTimeout creates a new DeleteScheduledJobUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteScheduledJobUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteScheduledJobUsingDELETEParams {
	var ()
	return &DeleteScheduledJobUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteScheduledJobUsingDELETEParamsWithContext creates a new DeleteScheduledJobUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteScheduledJobUsingDELETEParamsWithContext(ctx context.Context) *DeleteScheduledJobUsingDELETEParams {
	var ()
	return &DeleteScheduledJobUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteScheduledJobUsingDELETEParamsWithHTTPClient creates a new DeleteScheduledJobUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteScheduledJobUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteScheduledJobUsingDELETEParams {
	var ()
	return &DeleteScheduledJobUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteScheduledJobUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete scheduled job using d e l e t e operation typically these are written to a http.Request
*/
type DeleteScheduledJobUsingDELETEParams struct {

	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteScheduledJobUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) WithContext(ctx context.Context) *DeleteScheduledJobUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteScheduledJobUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVmid adds the vmid to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) WithVmid(vmid string) *DeleteScheduledJobUsingDELETEParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the delete scheduled job using d e l e t e params
func (o *DeleteScheduledJobUsingDELETEParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScheduledJobUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
