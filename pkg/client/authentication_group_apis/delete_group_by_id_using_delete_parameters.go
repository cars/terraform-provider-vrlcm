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

// NewDeleteGroupByIDUsingDELETEParams creates a new DeleteGroupByIDUsingDELETEParams object
// with the default values initialized.
func NewDeleteGroupByIDUsingDELETEParams() *DeleteGroupByIDUsingDELETEParams {
	var ()
	return &DeleteGroupByIDUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupByIDUsingDELETEParamsWithTimeout creates a new DeleteGroupByIDUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteGroupByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteGroupByIDUsingDELETEParams {
	var ()
	return &DeleteGroupByIDUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteGroupByIDUsingDELETEParamsWithContext creates a new DeleteGroupByIDUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteGroupByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteGroupByIDUsingDELETEParams {
	var ()
	return &DeleteGroupByIDUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteGroupByIDUsingDELETEParamsWithHTTPClient creates a new DeleteGroupByIDUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteGroupByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteGroupByIDUsingDELETEParams {
	var ()
	return &DeleteGroupByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteGroupByIDUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete group by Id using d e l e t e operation typically these are written to a http.Request
*/
type DeleteGroupByIDUsingDELETEParams struct {

	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteGroupByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteGroupByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteGroupByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVmid adds the vmid to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) WithVmid(vmid string) *DeleteGroupByIDUsingDELETEParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the delete group by Id using d e l e t e params
func (o *DeleteGroupByIDUsingDELETEParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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