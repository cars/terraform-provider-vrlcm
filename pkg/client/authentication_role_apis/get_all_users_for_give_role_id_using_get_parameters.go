// Code generated by go-swagger; DO NOT EDIT.

package authentication_role_a_p_is

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

// NewGetAllUsersForGiveRoleIDUsingGETParams creates a new GetAllUsersForGiveRoleIDUsingGETParams object
// with the default values initialized.
func NewGetAllUsersForGiveRoleIDUsingGETParams() *GetAllUsersForGiveRoleIDUsingGETParams {
	var ()
	return &GetAllUsersForGiveRoleIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllUsersForGiveRoleIDUsingGETParamsWithTimeout creates a new GetAllUsersForGiveRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllUsersForGiveRoleIDUsingGETParamsWithTimeout(timeout time.Duration) *GetAllUsersForGiveRoleIDUsingGETParams {
	var ()
	return &GetAllUsersForGiveRoleIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetAllUsersForGiveRoleIDUsingGETParamsWithContext creates a new GetAllUsersForGiveRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllUsersForGiveRoleIDUsingGETParamsWithContext(ctx context.Context) *GetAllUsersForGiveRoleIDUsingGETParams {
	var ()
	return &GetAllUsersForGiveRoleIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetAllUsersForGiveRoleIDUsingGETParamsWithHTTPClient creates a new GetAllUsersForGiveRoleIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllUsersForGiveRoleIDUsingGETParamsWithHTTPClient(client *http.Client) *GetAllUsersForGiveRoleIDUsingGETParams {
	var ()
	return &GetAllUsersForGiveRoleIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetAllUsersForGiveRoleIDUsingGETParams contains all the parameters to send to the API endpoint
for the get all users for give role Id using g e t operation typically these are written to a http.Request
*/
type GetAllUsersForGiveRoleIDUsingGETParams struct {

	/*Vmid
	  vmid

	*/
	Vmid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) WithTimeout(timeout time.Duration) *GetAllUsersForGiveRoleIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) WithContext(ctx context.Context) *GetAllUsersForGiveRoleIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) WithHTTPClient(client *http.Client) *GetAllUsersForGiveRoleIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVmid adds the vmid to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) WithVmid(vmid string) *GetAllUsersForGiveRoleIDUsingGETParams {
	o.SetVmid(vmid)
	return o
}

// SetVmid adds the vmid to the get all users for give role Id using g e t params
func (o *GetAllUsersForGiveRoleIDUsingGETParams) SetVmid(vmid string) {
	o.Vmid = vmid
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllUsersForGiveRoleIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
