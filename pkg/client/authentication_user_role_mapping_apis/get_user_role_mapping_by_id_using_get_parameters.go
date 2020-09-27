// Code generated by go-swagger; DO NOT EDIT.

package authentication_user_role_mapping_a_p_is

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

// NewGetUserRoleMappingByIDUsingGETParams creates a new GetUserRoleMappingByIDUsingGETParams object
// with the default values initialized.
func NewGetUserRoleMappingByIDUsingGETParams() *GetUserRoleMappingByIDUsingGETParams {
	var ()
	return &GetUserRoleMappingByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRoleMappingByIDUsingGETParamsWithTimeout creates a new GetUserRoleMappingByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRoleMappingByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetUserRoleMappingByIDUsingGETParams {
	var ()
	return &GetUserRoleMappingByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetUserRoleMappingByIDUsingGETParamsWithContext creates a new GetUserRoleMappingByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRoleMappingByIDUsingGETParamsWithContext(ctx context.Context) *GetUserRoleMappingByIDUsingGETParams {
	var ()
	return &GetUserRoleMappingByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetUserRoleMappingByIDUsingGETParamsWithHTTPClient creates a new GetUserRoleMappingByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRoleMappingByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetUserRoleMappingByIDUsingGETParams {
	var ()
	return &GetUserRoleMappingByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetUserRoleMappingByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get user role mapping by Id using g e t operation typically these are written to a http.Request
*/
type GetUserRoleMappingByIDUsingGETParams struct {

	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetUserRoleMappingByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) WithContext(ctx context.Context) *GetUserRoleMappingByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetUserRoleMappingByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVmid adds the vmid to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) WithVmid(vmid string) *GetUserRoleMappingByIDUsingGETParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the get user role mapping by Id using g e t params
func (o *GetUserRoleMappingByIDUsingGETParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRoleMappingByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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